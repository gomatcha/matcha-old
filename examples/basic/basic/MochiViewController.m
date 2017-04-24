//
//  MochiViewController.m
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiViewController.h"
#import "MochiView.h"
#import "MochiBridge.h"
#import "MochiNode.h"

@interface MochiViewController ()
@property (nonatomic, strong) NSString *name;
@property (nonatomic, strong) MochiView *mochiView;
@property (nonatomic, strong) MochiGoValue *buildContext;
@end

@implementation MochiViewController

+ (NSPointerArray *)viewControllers {
    static NSPointerArray *sPointerArray;
    static dispatch_once_t sOnce;
    dispatch_once(&sOnce, ^{
        sPointerArray = [NSPointerArray weakObjectsPointerArray];
    });
    return sPointerArray;
}

+ (void)reload {
    for (MochiViewController *i in [MochiViewController viewControllers]) {
        [i reload];
    }
}

- (id)initWithName:(NSString *)name {
    if ((self = [super initWithNibName:nil bundle:nil])) {
        self.name = name;
        [[MochiViewController viewControllers] addPointer:(__bridge void *)self];
    }
    return self;
}

- (void)loadView {
    MochiGoValue *root = [[MochiGoBridge sharedBridge] root];
    MochiGoValue *buildContext = [root call:@"NewBuildContext" args:nil][0];
    self.buildContext = buildContext;
    self.mochiView = [[MochiView alloc] initWithFrame:CGRectZero];
    self.view = self.mochiView;
    
    [self reload];
    dispatch_after(dispatch_time(DISPATCH_TIME_NOW, (int64_t)(3 * NSEC_PER_SEC)), dispatch_get_main_queue(), ^{
        [self reload];
    });
}

- (void)reload {
    [self.buildContext call:@"Build" args:nil];
    MochiGoValue *renderNode = [self.buildContext call:@"RenderNode" args:nil][0];
    [renderNode call:@"LayoutRoot" args:@[[[MochiGoValue alloc] initWithCGPoint:CGPointZero], [[MochiGoValue alloc] initWithCGPoint:CGPointMake(1000, 1000)]]];
    [renderNode call:@"Paint" args:nil];
    
    self.mochiView.node = [[MochiNode alloc] initWithGoValue:renderNode];
}

@end
