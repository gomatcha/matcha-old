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

@interface MochiView : UIView
- (id)initWithViewRoot:(MochiViewRoot *)viewRoot;
@property (nonatomic, weak) MochiViewRoot *viewRoot;
@property (nonatomic, strong) MochiNode *node;
@end

@interface MochiTextView : UILabel
- (id)initWithViewRoot:(MochiViewRoot *)viewRoot;
@property (nonatomic, weak) MochiViewRoot *viewRoot;
@end

@interface MochiImageView : UIImageView
- (id)initWithViewRoot:(MochiViewRoot *)viewRoot;
@property (nonatomic, weak) MochiViewRoot *viewRoot;
@end

@interface MochiButton : UIView
- (id)initWithViewRoot:(MochiViewRoot *)viewRoot;
@property (nonatomic, weak) MochiViewRoot *viewRoot;
@end

@interface MochiScrollView : UIScrollView
- (id)initWithViewRoot:(MochiViewRoot *)viewRoot;
@property (nonatomic, weak) MochiViewRoot *viewRoot;
@end
