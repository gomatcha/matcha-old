//
//  MochiViewController.h
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import <UIKit/UIKit.h>
@import Mochi;

@interface MochiViewController : UIViewController
+ (NSPointerArray *)viewControllers;
+ (void)render;
- (id)initWithName:(NSString *)name;
@end
