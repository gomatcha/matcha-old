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

@interface MochiView : UIView
@property (nonatomic, strong) MochiNode *node;
@end

@interface MochiTextView : MochiView
@end

@interface MochiImageView : MochiView
@end

@interface MochiButton : MochiView
@end
