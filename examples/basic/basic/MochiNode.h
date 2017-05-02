//
//  MochiNode.h
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

@import UIKit;
@import Mochi;
@class MochiPaintOptions;
@class MochiLayoutGuide;

@interface MochiNode : NSObject
- (id)initWithGoValue:(MochiGoValue *)value;
@property (nonatomic, readonly) NSMapTable<MochiGoValue *, MochiNode *> *nodeChildren;
@property (nonatomic, readonly) MochiLayoutGuide *guide;
@property (nonatomic, readonly) MochiPaintOptions *paintOptions;
@property (nonatomic, readonly) NSString *bridgeName;
@property (nonatomic, readonly) MochiGoValue *bridgeState;
@property (nonatomic, readonly) NSInteger buildId;
@end

@interface MochiPaintOptions : NSObject
- (id)initWithGoValue:(MochiGoValue *)value;
@property (nonatomic, readonly) UIColor *backgroundColor;
@end

@interface MochiLayoutGuide : NSObject
- (id)initWithGoValue:(MochiGoValue *)value;
@property (nonatomic, readonly) CGRect frame;
@property (nonatomic, readonly) UIEdgeInsets insets;
@property (nonatomic, readonly) NSInteger zIndex;
@end
