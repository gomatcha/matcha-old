//
//  MochiBridge.m
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiBridge.h"

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

- (BOOL)isEqual:(id)object {
    if ([object isKindOfClass:[BridgeValue class]]) {
        return [self ptrEqual:object];
    }
    return NO;
}

- (NSArray *)call:(NSString *)method args:(NSArray *)args {
    BridgeValueSlice *valueSlice = BridgeNewValueSlice();
    for (BridgeValue *i in args) {
        [valueSlice append:i];
    }
    
    BridgeValueSlice *result = [[self methodByName:method] call:valueSlice];
    return result.toArray;
}

- (BridgeValue *)get:(NSString *)field {
    BridgeValue *value = self.toUnderlying;
    return [value fieldByName:field];
}

- (BridgeValue *)toUnderlying {
    BridgeValue *value = self;
    while (true) {
        BridgeKind kind = value.kind;
        if (kind != BridgeKindInterface && kind != BridgeKindPtr) {
            break;
        }
        value = self.elem;
    }
    return value;
}

- (double)toDouble {
    return self.toUnderlying.float_;
}

- (unsigned long)toUnsignedLong {
    return self.toUnderlying.uint_;
}

- (long)toLong {
    return self.toUnderlying.int_;
}

- (NSData *)toData {
    return self.toUnderlying.bytes_;
}

- (NSString *)toString {
    return self.toUnderlying.string_;
}

- (NSDictionary *)toDictionary {
    BridgeValue *value = self.toUnderlying;
    NSMutableDictionary *dict = [NSMutableDictionary dictionary];
    for (BridgeValue *i in value.mapKeys.toArray) {
        id v = [value mapIndex:i];
        dict[i] = v;
    }
    return dict;
}

- (NSMapTable *)toMapTable {
    BridgeValue *value = self.toUnderlying;
    NSMapTable *mapTable = [NSMapTable strongToStrongObjectsMapTable];
    for (BridgeValue *i in value.mapKeys.toArray) {
        id v = [value mapIndex:i];
        mapTable[i] = v;
    }
    return mapTable;
}

- (NSArray *)toArray {
    BridgeValue *value = self.toUnderlying;
    NSMutableArray *array = [NSMutableArray array];
    for (NSInteger i = 0; i < value.len; i++) {
        [array addObject:[value index:i]];
    }
    return array;
}

- (NSNumber *)toNumber {
    BridgeValue *value = self.toUnderlying;
    
    BridgeKind kind = value.kind;
    if (kind == BridgeKindBool) {
        return [NSNumber numberWithBool:value.bool_];
    } else if (kind == BridgeKindInt ||
               kind == BridgeKindInt8 ||
               kind == BridgeKindInt16 ||
               kind == BridgeKindInt32 ||
               kind == BridgeKindInt64) {
        return [NSNumber numberWithLong:value.int_];
    } else if (kind == BridgeKindUint ||
               kind == BridgeKindUint8 ||
               kind == BridgeKindUint16 ||
               kind == BridgeKindUint32 ||
               kind == BridgeKindUint64 ||
               kind == BridgeKindUintptr) {
        return [NSNumber numberWithUnsignedLong:value.uint_];
    } else if (kind == BridgeKindFloat32 ||
               kind == BridgeKindFloat64) {
        return [NSNumber numberWithDouble:value.float_];
    }
    return nil;
}

- (BridgeValue *)copyWithZone:(NSZone *)zone {
    return self.copy;
}

- (BridgeValue *)objectForKeyedSubscript:(NSString *)key {
    return [self get:key];
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

@implementation UIColor (Mochi)
- (id)initWithBridgeValue:(BridgeValue *)value {
    NSArray<BridgeValue *> *array = [value call:@"RGBA" args:nil];
    return [UIColor colorWithRed:array[0].toUnsignedLong/0xffff green:array[1].toUnsignedLong/0xffff blue:array[2].toUnsignedLong/0xffff alpha:array[3].toUnsignedLong/0xffff];
}
@end

@implementation BridgeValue (Mochi)
- (CGRect)toCGRect {
    BridgeValue *v = self.toUnderlying;
    CGRect rect;
    rect.origin.x = [[v get:@"Min"] get:@"X"].toDouble;
    rect.origin.y = [[v get:@"Min"] get:@"Y"].toDouble;
    rect.size.width = [[v get:@"Max"] get:@"X"].toDouble - rect.origin.x;
    rect.size.height = [[v get:@"Max"] get:@"Y"].toDouble - rect.origin.y;
    return rect;
}

- (UIEdgeInsets)toUIEdgeInsets {
    BridgeValue *v = self.toUnderlying;
    UIEdgeInsets insets;
    insets.top = v[@"Top"].toDouble;
    insets.bottom = v[@"Bottom"].toDouble;
    insets.right = v[@"Right"].toDouble;
    insets.left = v[@"Left"].toDouble;
    return insets;
}
@end

@implementation NSMapTable (Mochi) 
- (id)objectForKeyedSubscript:(id)key {
    return [self objectForKey:key];
}

- (void)setObject:(id)obj forKeyedSubscript:(id)key {
    if (obj != nil) {
        [self setObject:obj forKey:key];
    } else {
        [self removeObjectForKey:key];
    }
}
@end