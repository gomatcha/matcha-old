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
@property (nonatomic, strong) MochiGoValue *viewController;
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

+ (void)render {
    for (MochiViewController *i in [MochiViewController viewControllers]) {
        [i render];
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
    self.viewController = [root call:@"NewViewController" args:nil][0];
    self.mochiView = [[MochiView alloc] initWithFrame:CGRectZero];
    self.view = self.mochiView;
}

- (void)viewDidLayoutSubviews {
    if (!CGRectEqualToRect(self.lastFrame, self.view.frame)) {
        self.lastFrame = self.view.frame;
        
        [self.viewController call:@"SetSize" args:@[[[MochiGoValue alloc] initWithCGPoint:CGPointMake(self.view.frame.size.width, self.view.frame.size.height)]]];
        [self render];
    }
}

- (void)render {
    NSLog(@"RENDER");
    
    MochiGoValue *renderNode = [self.viewController call:@"Render" args:nil][0];
    self.mochiView.node = [[MochiNode alloc] initWithGoValue:renderNode];
}

@end
