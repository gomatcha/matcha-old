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
@property (nonatomic, strong) MochiView *mochiView;
@end

@implementation MochiViewController

- (void)loadView {
    MochiGoValue *root = [[MochiGoBridge sharedBridge] root];
    MochiGoValue *value = [root call:@"Display" args:nil][0];
    MochiNode *node = [[MochiNode alloc] initWithGoValue:value];
    
    self.mochiView = [[MochiView alloc] initWithFrame:CGRectZero];
    self.mochiView.node = node;
    self.view = self.mochiView;
}

@end
