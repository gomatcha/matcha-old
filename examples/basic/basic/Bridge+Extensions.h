//
//  Bridge+Extensions.h
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import <Foundation/Foundation.h>
#import <Bridge/Bridge.h>

@interface BridgeValue (Extensions) <NSCopying>
- (NSArray<BridgeValue *> *)call:(NSString *)method args:(NSArray<BridgeValue *> *)args;
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
