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
@property (nonatomic, assign) NSInteger identifier;
@property (nonatomic, strong) MochiView *mochiView;
@property (nonatomic, strong) MochiGoValue *goVC;
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

+ (MochiViewController *)viewControllerWithIdentifier:(NSInteger)identifier {
    for (MochiViewController *i in [self viewControllers]) {
        if (i.identifier == identifier) {
            return i;
        }
    }
    return nil;
}

+ (void)render {
    for (MochiViewController *i in [MochiViewController viewControllers]) {
        [i render];
    }
}

- (id)initWithName:(NSString *)name {
    if ((self = [super initWithNibName:nil bundle:nil])) {
        self.identifier = arc4random();
        [[MochiViewController viewControllers] addPointer:(__bridge void *)self];
        
        self.goVC = [[[MochiGoBridge sharedBridge] root] call:@"NewViewController" args:@[[[MochiGoValue alloc] initWithInt:self.identifier]]][0];
    }
    return self;
}

- (void)loadView {
    self.mochiView = [[MochiView alloc] initWithFrame:CGRectZero];
    self.view = self.mochiView;
}

- (void)viewDidLayoutSubviews {
    if (!CGRectEqualToRect(self.lastFrame, self.view.frame)) {
        self.lastFrame = self.view.frame;
        
        [self.goVC call:@"SetSize" args:@[[[MochiGoValue alloc] initWithCGPoint:CGPointMake(self.view.frame.size.width, self.view.frame.size.height)]]];
        [self render];
    }
}

- (void)render {
    NSLog(@"RENDER");
    
    MochiGoValue *renderNode = [self.goVC call:@"Render" args:nil][0];
    self.mochiView.node = [[MochiNode alloc] initWithGoValue:renderNode];
}

- (void)update:(MochiNode *)node {
    self.mochiView.node = node;
}

@end
