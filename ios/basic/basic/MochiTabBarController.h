//
//  MochiTabViewController.h
//  basic
//
//  Created by Kevin Dang on 5/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import <UIKit/UIKit.h>
@class MochiViewController;

@interface MochiTabBarController : UITabBarController
- (id)initWithViewRoot:(MochiViewController *)viewRoot;
@property (nonatomic, weak) MochiViewController *viewRoot;
@end
