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

- (id)initWithName:(NSString *)name {
    if ((self = [super initWithNibName:nil bundle:nil])) {
        self.name = name;
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
}

- (void)reload {
    [self.buildContext call:@"Build" args:nil];
    MochiGoValue *renderNode = [self.buildContext call:@"RenderNode" args:nil][0];
    [renderNode call:@"Layout" args:@[[[MochiGoValue alloc] initWithCGPoint:CGPointZero], [[MochiGoValue alloc] initWithCGPoint:CGPointMake(1000, 1000)]]];
    [renderNode call:@"Paint" args:nil];
    
    self.mochiView.node = [[MochiNode alloc] initWithGoValue:renderNode];
}

@end
