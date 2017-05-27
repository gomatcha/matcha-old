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
@class MochiNode;
@class MochiPBNode;
@class MochiPBRoot;
@class MochiPBGuide;
@class MochiPBPaintStyle;
@class GPBAny;

@interface MochiNodeRoot : NSObject
- (id)initWithProtobuf:(MochiPBRoot *)data;
@property (nonatomic, readonly) MochiNode *node;
@end

@interface MochiNode : NSObject
- (id)initWithProtobuf:(MochiPBNode *)node;
- (id)initWithGoValue:(MochiGoValue *)value;
@property (nonatomic, readonly) NSDictionary<NSNumber *, MochiNode *> *nodeChildren;
@property (nonatomic, readonly) MochiLayoutGuide *guide;
@property (nonatomic, readonly) MochiPaintOptions *paintOptions;
@property (nonatomic, readonly) NSMutableDictionary<NSString*, GPBAny*> *nativeValues;
@property (nonatomic, readonly) NSString *nativeViewName;
@property (nonatomic, readonly) GPBAny *nativeViewState;
@property (nonatomic, readonly) NSNumber *identifier;
@property (nonatomic, readonly) NSNumber *buildId;
@property (nonatomic, readonly) NSNumber *layoutId;
@property (nonatomic, readonly) NSNumber *paintId;
@end

@interface MochiPaintOptions : NSObject
- (id)initWithProtobuf:(MochiPBPaintStyle *)style;
- (id)initWithGoValue:(MochiGoValue *)value;
@property (nonatomic, readonly) UIColor *backgroundColor;
@end

@interface MochiLayoutGuide : NSObject
- (id)initWithProtobuf:(MochiPBGuide *)guide;
- (id)initWithGoValue:(MochiGoValue *)value;
@property (nonatomic, readonly) CGRect frame;
@property (nonatomic, readonly) UIEdgeInsets insets;
@property (nonatomic, readonly) NSInteger zIndex;
@end
