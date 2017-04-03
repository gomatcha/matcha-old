//
//  MochiNode.h
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import <Foundation/Foundation.h>
#import "MochiBridge.h"
@class MochiPaintOptions;
@class MochiLayoutGuide;

@interface MochiNode : NSObject
- (id)initWithBridgeValue:(BridgeValue *)value;
@property (nonatomic, readonly) NSDictionary<BridgeValue *, MochiNode *> *nodeChildren;
@property (nonatomic, readonly) MochiLayoutGuide *guide;
@property (nonatomic, readonly) MochiPaintOptions *paintOptions;
@end

@interface MochiPaintOptions : NSObject
- (id)initWithBridgeValue:(BridgeValue *)value;
@property (nonatomic, readonly) UIColor *backgroundColor;
@end

@interface MochiLayoutGuide : NSObject
- (id)initWithBridgeValue:(BridgeValue *)value;
@property (nonatomic, readonly) CGRect frame;
@property (nonatomic, readonly) UIEdgeInsets insets;
@end