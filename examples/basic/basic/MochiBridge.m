//
//  MochiBridge.m
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiBridge.h"

// typedef NS_ENUM(NSInteger, BridgeKind) {
//     BridgeKindInvalid, // 0
//     BridgeKindBool,
//     BridgeKindInt,
//     BridgeKindInt8,
//     BridgeKindInt16,
//     BridgeKindInt32, // 5
//     BridgeKindInt64,
//     BridgeKindUint,
//     BridgeKindUint8,
//     BridgeKindUint16,
//     BridgeKindUint32, // 10
//     BridgeKindUint64,
//     BridgeKindUintptr,
//     BridgeKindFloat32,
//     BridgeKindFloat64,
//     BridgeKindComplex64, // 15
//     BridgeKindComplex128,
//     BridgeKindArray,
//     BridgeKindChan,
//     BridgeKindFunc,
//     BridgeKindInterface, // 20
//     BridgeKindMap,
//     BridgeKindPtr,
//     BridgeKindSlice,
//     BridgeKindString,
//     BridgeKindStruct, // 25
//     BridgeKindUnsafePointer,
// };

// @implementation BridgeValue (Extensions)

// - (BOOL)isEqual:(id)object {
//     if ([object isKindOfClass:[BridgeValue class]]) {
//         return [self ptrEqual:object];
//     }
//     return NO;
// }

// - (NSArray *)call:(NSString *)method args:(NSArray *)args {
//     BridgeValueSlice *valueSlice = BridgeNewValueSlice();
//     for (BridgeValue *i in args) {
//         [valueSlice append:i];
//     }
    
//     BridgeValue *m = [self methodByName:method];
//     BridgeValueSlice *result = [m call:valueSlice];
//     return result.toArray;
// }

// - (BridgeValue *)get:(NSString *)field {
//     BridgeValue *value = self.toUnderlying;
//     return [value fieldByName:field];
// }

// - (BridgeValue *)toUnderlying {
//     BridgeValue *value = self;
//     while (true) {
//         BridgeKind kind = value.kind;
//         if (kind != BridgeKindInterface && kind != BridgeKindPtr) {
//             break;
//         }
//         value = value.elem;
//     }
//     return value;
// }

// - (double)toDouble {
//     return self.toUnderlying.float_;
// }

// - (unsigned long)toUnsignedLong {
//     return self.toUnderlying.uint_;
// }

// - (long)toLong {
//     return self.toUnderlying.int_;
// }

// - (NSData *)toData {
//     return self.toUnderlying.bytes_;
// }

// - (NSString *)toString {
//     return self.toUnderlying.string_;
// }

// - (NSDictionary *)toDictionary {
//     BridgeValue *value = self.toUnderlying;
//     NSMutableDictionary *dict = [NSMutableDictionary dictionary];
//     for (BridgeValue *i in value.mapKeys.toArray) {
//         id v = [value mapIndex:i];
//         dict[i] = v;
//     }
//     return dict;
// }

// - (NSMapTable *)toMapTable {
    // BridgeValue *value = self.toUnderlying;
    // NSMapTable *mapTable = [NSMapTable strongToStrongObjectsMapTable];
    // for (BridgeValue *i in value.mapKeys.toArray) {
    //     id v = [value mapIndex:i];
    //     mapTable[i] = v;
    // }
    // return mapTable;
// }

// - (NSArray *)toArray {
//     BridgeValue *value = self.toUnderlying;
//     NSMutableArray *array = [NSMutableArray array];
//     for (NSInteger i = 0; i < value.len; i++) {
//         [array addObject:[value index:i]];
//     }
//     return array;
// }

// - (NSNumber *)toNumber {
//     BridgeValue *value = self.toUnderlying;
    
//     BridgeKind kind = value.kind;
//     if (kind == BridgeKindBool) {
//         return [NSNumber numberWithBool:value.bool_];
//     } else if (kind == BridgeKindInt ||
//                kind == BridgeKindInt8 ||
//                kind == BridgeKindInt16 ||
//                kind == BridgeKindInt32 ||
//                kind == BridgeKindInt64) {
//         return [NSNumber numberWithLong:value.int_];
//     } else if (kind == BridgeKindUint ||
//                kind == BridgeKindUint8 ||
//                kind == BridgeKindUint16 ||
//                kind == BridgeKindUint32 ||
//                kind == BridgeKindUint64 ||
//                kind == BridgeKindUintptr) {
//         return [NSNumber numberWithUnsignedLong:value.uint_];
//     } else if (kind == BridgeKindFloat32 ||
//                kind == BridgeKindFloat64) {
//         return [NSNumber numberWithDouble:value.float_];
//     }
//     return nil;
// }

// - (BridgeValue *)copyWithZone:(NSZone *)zone {
//     return self.copy;
// }

// - (BridgeValue *)objectForKeyedSubscript:(NSString *)key {
//     return [self get:key];
// }

// @end

// @implementation BridgeValueSlice (Extensions)

// - (NSArray *)toArray {
//     NSMutableArray *array = [NSMutableArray array];
//     for (NSInteger i = 0; i < self.len; i++) {
//         [array addObject:[self index:i]];
//     }
//     return array;
// }

// @end

@implementation UIColor (Mochi)
- (id)initWithGoValue:(MochiGoValue *)value {
    NSArray<MochiGoValue *> *array = [value call:@"RGBA" args:nil];
    return [UIColor colorWithRed:array[0].toUnsignedLongLong/0xffff green:array[1].toUnsignedLongLong/0xffff blue:array[2].toUnsignedLongLong/0xffff alpha:array[3].toUnsignedLongLong/0xffff];
}
@end

@implementation MochiGoValue (Mochi)
- (CGRect)toCGRect {
    CGRect rect;
    rect.origin.x = self[@"Min"][@"X"].toDouble;
    rect.origin.y = self[@"Min"][@"Y"].toDouble;
    rect.size.width = self[@"Max"][@"X"].toDouble - rect.origin.x;
    rect.size.height = self[@"Max"][@"Y"].toDouble - rect.origin.y;
    return rect;
}

- (UIEdgeInsets)toUIEdgeInsets {
    UIEdgeInsets insets;
    insets.top = self[@"Top"].toDouble;
    insets.bottom = self[@"Bottom"].toDouble;
    insets.right = self[@"Right"].toDouble;
    insets.left = self[@"Left"].toDouble;
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

@implementation NSAttributedString (Mochi)
- (id)initWithGoValue:(MochiGoValue *)value {
    NSString *string = [value call:@"String" args:nil][0].toString;
    MochiGoValue *format = [value call:@"Format" args:nil][0];
    NSMapTable *attrTable = [format call:@"Attributes" args:nil][0].toMapTable;

    NSMutableParagraphStyle *paragraphStyle = [[NSMutableParagraphStyle alloc] init];
    NSMutableDictionary *dictionary = [[NSMutableDictionary alloc] init];
    dictionary[NSParagraphStyleAttributeName] = paragraphStyle;

    for (MochiGoValue *i in attrTable.keyEnumerator) {
        MochiGoValue *value = ((MochiGoValue *)attrTable[i]).elem;
        NSInteger key = i.toLongLong;
        switch (key) {
        case 0: { // AttributeKeyAlignment
            NSTextAlignment alignment;
            switch (value.toLongLong) {
            case 0:
                alignment = NSTextAlignmentLeft;
                break;
            case 1: 
                alignment = NSTextAlignmentRight;
                break;
            case 2:
                alignment = NSTextAlignmentCenter;
                break;
            case 3:
                alignment = NSTextAlignmentJustified;
                break;
            default:
                alignment = NSTextAlignmentLeft;
            }
            paragraphStyle.alignment = alignment;
            break;
        }
        case 1: { //AttributeKeyStrikethroughStyle
            NSUnderlineStyle style;
            switch (value.toLongLong) {
            case 0:
                style = NSUnderlineStyleNone;
                break;
            case 1: 
                style = NSUnderlineStyleSingle;
                break;
            case 2:
                style = NSUnderlineStyleDouble;
                break;
            case 3:
                style = NSUnderlineStyleThick;
                break;
            case 4:
                style = NSUnderlinePatternDot;
                break;
            case 5:
                style = NSUnderlinePatternDash;
                break;
            default:
                style = NSUnderlineStyleNone;
            }
            dictionary[NSStrikethroughStyleAttributeName] = @(style);
            break;
        }
        case 2: { //AttributeKeyStrikethroughColor
            dictionary[NSStrikethroughColorAttributeName] = [[UIColor alloc] initWithGoValue:value];
            break;
        }
        case 3: { //AttributeKeyUnderlineStyle
            NSUnderlineStyle style;
            switch (value.toLongLong) {
            case 0:
                style = NSUnderlineStyleNone;
                break;
            case 1: 
                style = NSUnderlineStyleSingle;
                break;
            case 2:
                style = NSUnderlineStyleDouble;
                break;
            case 3:
                style = NSUnderlineStyleThick;
                break;
            case 4:
                style = NSUnderlinePatternDot;
                break;
            case 5:
                style = NSUnderlinePatternDash;
                break;
            default:
                style = NSUnderlineStyleNone;
            }
            dictionary[NSUnderlineStyleAttributeName] = @(style);
            break;
        }
        case 4: { //AttributeKeyUnderlineColor
            dictionary[NSUnderlineColorAttributeName] = [[UIColor alloc] initWithGoValue:value];
            break;
        }
        case 5: { //AttributeKeyFont
            dictionary[NSFontAttributeName] = [[UIFont alloc] initWithGoValue:value];
            break;
        }
        case 6: { //AttributeKeyHyphenation
            // dictionary[NSHyphenationFactorDocumentAttribute] = value.toNumber; // TODO(KD):
            break;
        }
        case 7: { //AttributeKeyLineHeightMultiple
            paragraphStyle.lineHeightMultiple = value.toDouble;
            break;
        }
        case 8: { //AttributeKeyMaxLines
            // TODO(KD):
            break;
        }
        case 9: { //AttributeKeyTextColor
            dictionary[NSForegroundColorAttributeName] = [[UIColor alloc] initWithGoValue:value];
            break;
        }
        case 10: { //AttributeKeyTextWrap
            // TODO(KD):
            break;
        }
        case 11: { //AttributeKeyTruncation
            // TODO(KD):
            break;
        }
        case 12: { //AttributeKeyTruncationString 
            // TODO(KD):
            break;
        }
        }
    }
    return [[NSAttributedString alloc] initWithString:string attributes:dictionary];
}
@end

@implementation UIFont (Mochi)
- (id)initWithGoValue:(MochiGoValue *)value {
    NSMutableDictionary *attr = [NSMutableDictionary dictionary];
    attr[UIFontDescriptorFamilyAttribute] = value[@"Family"].toString;
    attr[UIFontDescriptorFaceAttribute] = value[@"Face"].toString;
    attr[UIFontDescriptorSizeAttribute] = @(value[@"Size"].toDouble);

    UIFontDescriptor *desc = [[UIFontDescriptor alloc] initWithFontAttributes:attr];
    UIFont *font = [UIFont fontWithDescriptor:desc size:0];
    return font;
}
@end
