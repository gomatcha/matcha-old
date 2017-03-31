//
//  MochiViewController.m
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiViewController.h"
#import "MochiView.h"
#import "Bridge+Extensions.h"

@interface MochiViewController ()
@property (nonatomic, strong) MochiView *mochiView;
@end

@implementation MochiViewController

- (void)loadView {
    self.mochiView = [[MochiView alloc] initWithFrame:CGRectZero];
    self.mochiView.node = BridgeRun();
    self.view = self.mochiView;
}

@end
