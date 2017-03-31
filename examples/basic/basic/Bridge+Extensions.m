//
//  Bridge+Extensions.m
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "Bridge+Extensions.h"

typedef NS_ENUM(NSInteger, BridgeKind) {
    BridgeKindInvalid,
    BridgeKindBool,
    BridgeKindInt,
    BridgeKindInt8,
    BridgeKindInt16,
    BridgeKindInt32,
    BridgeKindInt64,
    BridgeKindUint,
    BridgeKindUint8,
    BridgeKindUint16,
    BridgeKindUint32,
    BridgeKindUint64,
    BridgeKindUintptr,
    BridgeKindFloat32,
    BridgeKindFloat64,
    BridgeKindComplex64,
    BridgeKindComplex128,
    BridgeKindArray,
    BridgeKindChan,
    BridgeKindFunc,
    BridgeKindInterface,
    BridgeKindMap,
    BridgeKindPtr,
    BridgeKindSlice,
    BridgeKindString,
    BridgeKindStruct,
    BridgeKindUnsafePointer,
};

@implementation BridgeValue (Extensions)

- (NSArray *)call:(NSString *)method args:(NSArray *)args {
    BridgeValueSlice *valueSlice = BridgeNewValueSlice();
    for (BridgeValue *i in args) {
        [valueSlice append:i];
    }

    BridgeValueSlice *result = [[self methodByName:method] call:valueSlice];
    return result.toArray;
}

- (double)toDouble {
    BridgeKind kind = self.kind;
    if (kind == BridgeKindInterface || 
        kind == BridgeKindPtr) {
        return [self.elem toDouble];
    }
    return self.float_;
}

- (unsigned long)toUnsignedLong {
    BridgeKind kind = self.kind;
    if (kind == BridgeKindInterface || 
        kind == BridgeKindPtr) {
        return [self.elem toUnsignedLong];
    }
    return self.uint_;
}

- (long)toLong {
    BridgeKind kind = self.kind;
    if (kind == BridgeKindInterface || 
        kind == BridgeKindPtr) {
        return [self.elem toLong];
    }
    return self.int_;
}

- (NSData *)toData {
    BridgeKind kind = self.kind;
    if (kind == BridgeKindInterface || 
        kind == BridgeKindPtr) {
        return [self.elem toData];
    }
    return self.bytes_;
}

- (NSString *)toString {
    BridgeKind kind = self.kind;
    if (kind == BridgeKindInterface || 
        kind == BridgeKindPtr) {
        return [self.elem toString];
    }
    return self.string_;
}

- (NSDictionary *)toDictionary {
    BridgeKind kind = self.kind;
    if (kind == BridgeKindInterface || 
        kind == BridgeKindPtr) {
        return [self.elem toDictionary];
    }

    NSMutableDictionary *dict = [NSMutableDictionary dictionary];
    for (BridgeValue *i in self.mapKeys.toArray) {
        dict[i] = [self mapIndex:i];
    }
    return dict;
}

- (NSArray *)toArray {
    BridgeKind kind = self.kind;
    if (kind == BridgeKindInterface || 
        kind == BridgeKindPtr) {
        return [self.elem toArray];
    }

    NSMutableArray *array = [NSMutableArray array];
    for (NSInteger i = 0; i < self.len; i++) {
        [array addObject:[self index:i]];
    }
    return array;
}

- (NSNumber *)toNumber {
    BridgeKind kind = self.kind;
    if (kind == BridgeKindInterface || 
        kind == BridgeKindPtr) {
        return [self.elem toNumber];
    } else if (kind == BridgeKindBool) {
        return [NSNumber numberWithBool:self.bool_];
    } else if (kind == BridgeKindInt || 
        kind == BridgeKindInt8 || 
        kind == BridgeKindInt16 || 
        kind == BridgeKindInt32 || 
        kind == BridgeKindInt64) {
        return [NSNumber numberWithLong:self.int_];
    } else if (kind == BridgeKindUint || 
        kind == BridgeKindUint8 || 
        kind == BridgeKindUint16 || 
        kind == BridgeKindUint32 || 
        kind == BridgeKindUint64 || 
        kind == BridgeKindUintptr) {
        return [NSNumber numberWithUnsignedLong:self.uint_];
    } else if (kind == BridgeKindFloat32 || 
        kind == BridgeKindFloat64) {
        return [NSNumber numberWithDouble:self.float_];
    }
    return nil;
}

- (BridgeValue *)copyWithZone:(NSZone *)zone {
    return self.copy;
}

@end

@implementation BridgeValueSlice (Extensions)

- (NSArray *)toArray {
    NSMutableArray *array = [NSMutableArray array];
    for (NSInteger i = 0; i < self.len; i++) {
        [array addObject:[self index:i]];
    }
    return array;
}

@end
