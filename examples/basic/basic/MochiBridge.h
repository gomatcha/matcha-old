//
//  MochiBridge.h
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import <UIKit/UIKit.h>
@import Mochi;
@class MochiPBText;
@class MochiPBColor;

@interface UIColor (Mochi)
- (id)initWithGoValue:(MochiGoValue *)value;
- (id)initWithProtobuf(MochiPBColor *)value;
@end

@interface MochiGoValue (Mochi)
- (id)initWithCGPoint:(CGPoint)point;
- (id)initWithCGSize:(CGSize)size;
- (id)initWithCGRect:(CGRect)rect;
- (CGPoint)toCGPoint;
- (CGSize)toCGSize;
- (CGRect)toCGRect;
- (UIEdgeInsets)toUIEdgeInsets;
@end

@interface NSMapTable (Mochi)
- (id)objectForKeyedSubscript:(id)key;
- (void)setObject:(id)obj forKeyedSubscript:(id)key;
@end

@interface NSAttributedString (Mochi)
- (id)initWithGoValue:(MochiGoValue *)value;
- (id)initWithProtobuf:(MochiPBText *)value;
@end

@interface UIFont (Mochi)
- (id)initWithGoValue:(MochiGoValue *)value;
- (id)initWithProtobuf(MochiPBFont *)value;
@end
