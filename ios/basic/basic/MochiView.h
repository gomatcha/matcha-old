//
//  MochiView.h
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

@import UIKit;
#import "MochiBridge.h"
#import "MochiNode.h"
@class MochiViewConfig;
@class MochiViewController;

@interface MochiView : UIView
- (id)initWithViewRoot:(MochiViewController *)viewRoot;
@property (nonatomic, weak) MochiViewController *viewRoot;
@property (nonatomic, strong) MochiNode *node;
@end

@interface MochiTextView : UILabel
- (id)initWithViewRoot:(MochiViewController *)viewRoot;
@property (nonatomic, weak) MochiViewController *viewRoot;
@end

@interface MochiImageView : UIImageView
- (id)initWithViewRoot:(MochiViewController *)viewRoot;
@property (nonatomic, weak) MochiViewController *viewRoot;
@end

@interface MochiButton : UIView
- (id)initWithViewRoot:(MochiViewController *)viewRoot;
@property (nonatomic, weak) MochiViewController *viewRoot;
@end

@interface MochiScrollView : UIScrollView
- (id)initWithViewRoot:(MochiViewController *)viewRoot;
@property (nonatomic, weak) MochiViewController *viewRoot;
@end
