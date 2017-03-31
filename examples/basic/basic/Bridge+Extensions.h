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
- (NSArray *)call:(NSString *)method args:(NSArray *)args;
- (NSData *)toData;
- (NSString *)toString;
- (NSDictionary *)toDictionary;
- (NSArray *)toArray;
- (NSNumber *)toNumber;
@end

@interface BridgeValueSlice (Extensions)
- (NSArray *)toArray;
@end