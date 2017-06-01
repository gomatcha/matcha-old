//
//  MochiTabViewController.m
//  basic
//
//  Created by Kevin Dang on 5/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiTabBarController.h"

@implementation MochiTabBarController

- (id)initWithViewRoot:(MochiViewController *)viewRoot {
    if ((self = [super init])) {
        self.viewRoot = viewRoot;
    }
    return self;
}

@end
