//
//  MochiViewController.h
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import <UIKit/UIKit.h>
@import Mochi;
@class MochiNode;
@class MochiViewRoot;

@interface MochiViewController : UIViewController
+ (NSPointerArray *)viewControllers;
+ (MochiViewController *)viewControllerWithIdentifier:(NSInteger)identifier;

- (id)initWithMochiViewRoot:(MochiViewRoot *)root;
- (void)update:(MochiNode *)node;
@property (nonatomic, readonly) MochiViewRoot *viewRoot;
@property (nonatomic, readonly) NSInteger identifier;
@end
