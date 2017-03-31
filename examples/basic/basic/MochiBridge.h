//
//  MochiBridge.h
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import <UIKit/UIKit.h>
#import <Bridge/Bridge.h>

@interface BridgeValue (Extensions) <NSCopying>
- (NSArray<BridgeValue *> *)call:(NSString *)method args:(NSArray<BridgeValue *> *)args;
- (BridgeValue *)get:(NSString *)field;
- (BridgeValue *)toUnderlying;
- (NSData *)toData;
- (NSString *)toString;
- (NSDictionary *)toDictionary;
- (NSArray<BridgeValue *> *)toArray;
- (NSNumber *)toNumber;
- (double)toDouble;
- (unsigned long)toUnsignedLong;
- (long)toLong;
@end

@interface BridgeValueSlice (Extensions)
- (NSArray *)toArray;
@end

@interface UIColor (Mochi)
- (id)initWithBridgeValue:(BridgeValue *)value;
@end

@interface BridgeValue (Mochi)
- (CGRect)toCGRect;
@end
