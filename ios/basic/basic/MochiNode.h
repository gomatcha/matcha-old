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
@class MochiPBRecognizer;
@class GPBAny;

@interface MochiNodeRoot : NSObject // view.root
- (id)initWithProtobuf:(MochiPBRoot *)data;
@property (nonatomic, readonly) MochiNode *node;
@end

@interface MochiNode : NSObject // view.node
- (id)initWithProtobuf:(MochiPBNode *)node;
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
@property (nonatomic, readonly) NSDictionary<NSNumber *, GPBAny *> *touchRecognizers;
@end

@interface MochiPaintOptions : NSObject
- (id)initWithProtobuf:(MochiPBPaintStyle *)style;
@property (nonatomic, readonly) CGFloat transparency;
@property (nonatomic, readonly) UIColor *backgroundColor;
@property (nonatomic, readonly) UIColor *borderColor;
@property (nonatomic, readonly) CGFloat borderWidth;
@property (nonatomic, readonly) CGFloat cornerRadius;
@property (nonatomic, readonly) CGFloat shadowRadius;
@property (nonatomic, readonly) CGSize shadowOffset;
@property (nonatomic, readonly) UIColor *shadowColor;
@end

@interface MochiLayoutGuide : NSObject
- (id)initWithProtobuf:(MochiPBGuide *)guide;
@property (nonatomic, readonly) CGRect frame;
@property (nonatomic, readonly) UIEdgeInsets insets;
@property (nonatomic, readonly) NSInteger zIndex;
@end
