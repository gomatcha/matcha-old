//
//  MochiBridge.m
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiBridge.h"

@implementation UIColor (Mochi)
- (id)initWithBridgeValue:(BridgeValue *)value {
    NSArray<BridgeValue *> *array = [value call:@"RGBA" args:nil];
    NSLog(@"%@",@(array[0].toUnsignedLong));
    return [UIColor colorWithRed:array[0].toUnsignedLong/0xffff green:array[1].toUnsignedLong/0xffff blue:array[2].toUnsignedLong/0xffff alpha:array[3].toUnsignedLong/0xffff];
}
@end
