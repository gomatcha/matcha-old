//
//  MochiBridge.h
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import <UIKit/UIKit.h>
@import Mochi;

// @interface BridgeValue (Extensions) <NSCopying>
// - (NSArray<BridgeValue *> *)call:(NSString *)method args:(NSArray<BridgeValue *> *)args;
// - (BridgeValue *)get:(NSString *)field;
// - (BridgeValue *)toUnderlying;
// - (NSData *)toData;
// - (NSString *)toString;
// // - (NSDictionary<BridgeValue *, BridgeValue *> *)toDictionary; TODO(KD): Issue where adding values get turned to nil in dictionary
// - (NSMapTable *)toMapTable;
// - (NSArray<BridgeValue *> *)toArray;
// - (NSNumber *)toNumber;
// - (double)toDouble;
// - (unsigned long)toUnsignedLong;
// - (long)toLong;
// - (BridgeValue *)objectForKeyedSubscript:(NSString *)key;
// @end

// @interface BridgeValueSlice (Extensions)
// - (NSArray *)toArray;
// @end

@interface UIColor (Mochi)
- (id)initWithGoValue:(MochiGoValue *)value;
@end

@interface MochiGoValue (Mochi)
- (CGRect)toCGRect;
- (UIEdgeInsets)toUIEdgeInsets;
@end

// @interface NSMapTable (Mochi)
// - (id)objectForKeyedSubscript:(id)key;
// - (void)setObject:(id)obj forKeyedSubscript:(id)key;
// @end

// @interface NSAttributedString (Mochi)
// - (id)initWithBridgeValue:(BridgeValue *)value;
// @end

// @interface UIFont (Mochi)
// - (id)initWithBridgeValue:(BridgeValue *)value;
// @end
