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
@property (nonatomic, assign) CGRect lastFrame;
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
}

- (void)viewDidLayoutSubviews {
    if (!CGRectEqualToRect(self.lastFrame, self.view.frame)) {
        self.lastFrame = self.view.frame;
        [self reload];
    }
}

- (void)reload {
    NSLog(@"KD:%s", __FUNCTION__);
    [self.buildContext call:@"Build" args:nil];
    MochiGoValue *renderNode = [self.buildContext call:@"RenderNode" args:nil][0];
    [renderNode call:@"LayoutRoot" args:@[[[MochiGoValue alloc] initWithCGPoint:CGPointZero], [[MochiGoValue alloc] initWithCGPoint:CGPointMake(self.view.frame.size.width, self.view.frame.size.height)]]];
    [renderNode call:@"Paint" args:nil];
    
    self.mochiView.node = [[MochiNode alloc] initWithGoValue:renderNode];
}

@end
