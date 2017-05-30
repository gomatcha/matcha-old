//
//  MochiBridge.m
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiBridge.h"
#import "Text.pbobjc.h"
#import "Imageview.pbobjc.h"

@implementation UIColor (Mochi)

- (id)initWithGoValue:(MochiGoValue *)value {
    NSArray<MochiGoValue *> *array = [value call:@"RGBA" args:nil];
    return [UIColor colorWithRed:((double)array[0].toUnsignedLongLong)/0xffff green:((double)array[1].toUnsignedLongLong)/0xffff blue:((double)array[2].toUnsignedLongLong)/0xffff alpha:((double)array[3].toUnsignedLongLong)/0xffff];
}

- (id)initWithProtobuf:(MochiPBColor *)value {
    if (value == nil) {
        return nil;
    }
    return [UIColor colorWithRed:((double)value.red)/0xffff green:((double)value.green)/0xffff blue:((double)value.blue)/0xffff alpha:((double)value.alpha)/0xffff];
}

@end

@implementation MochiGoValue (Mochi)

- (id)initWithCGPoint:(CGPoint)point {
    if ((self = [self initWithType:@"layout.Point"].elem)) {
        self[@"X"] = [[MochiGoValue alloc] initWithDouble:point.x];
        self[@"Y"] = [[MochiGoValue alloc] initWithDouble:point.y];
    }
    return self;
}

- (id)initWithCGSize:(CGSize)size {
     if ((self = [self initWithType:@"layout.Point"].elem)) {
         self[@"X"] = [[MochiGoValue alloc] initWithDouble:size.width];
         self[@"Y"] = [[MochiGoValue alloc] initWithDouble:size.height];
     }
     return self;
}

- (id)initWithCGRect:(CGRect)rect {
    if ((self = [self initWithType:@"layout.Rect"].elem)) {
        self[@"Min"] = [[MochiGoValue alloc] initWithCGPoint:rect.origin];
        self[@"Max"] = [[MochiGoValue alloc] initWithCGPoint:CGPointMake(rect.origin.x + rect.size.width, rect.origin.y + rect.size.height)];
    }
    return self;
}

- (CGPoint)toCGPoint {
    CGPoint point;
    point.x = self[@"X"].toDouble;
    point.y = self[@"Y"].toDouble;
    return point;
}

- (CGSize)toCGSize {
    CGSize size;
    size.width = self[@"X"].toDouble;
    size.height = self[@"Y"].toDouble;
    return size;
}

- (CGRect)toCGRect {
    MochiGoValue *min = self[@"Min"];
    MochiGoValue *max = self[@"Max"];
    CGRect rect;
    rect.origin.x = min[@"X"].toDouble;
    rect.origin.y = min[@"Y"].toDouble;
    rect.size.width = max[@"X"].toDouble - rect.origin.x;
    rect.size.height = max[@"Y"].toDouble - rect.origin.y;
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
    MochiGoValue *format = [value call:@"Style" args:nil][0];
    NSMapTable *attrTable = [format call:@"Map" args:nil][0].toMapTable;

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
            dictionary[NSHyphenationFactorDocumentAttribute] = @(value.toLongLong);
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

- (id)initWithProtobuf:(MochiPBText *)value {
    NSString *string = value.text;
    MochiPBTextStyle *style = value.style;
    NSMutableParagraphStyle *paragraphStyle = [[NSMutableParagraphStyle alloc] init];
    NSMutableDictionary *dictionary = [[NSMutableDictionary alloc] init];
    dictionary[NSParagraphStyleAttributeName] = paragraphStyle;

    
    NSTextAlignment alignment;
    switch (style.textAlignment) {
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
    
    NSUnderlineStyle strikethroughStyle;
    switch (style.strikethroughStyle) {
    case 0:
        strikethroughStyle = NSUnderlineStyleNone;
        break;
    case 1: 
        strikethroughStyle = NSUnderlineStyleSingle;
        break;
    case 2:
        strikethroughStyle = NSUnderlineStyleDouble;
        break;
    case 3:
        strikethroughStyle = NSUnderlineStyleThick;
        break;
    case 4:
        strikethroughStyle = NSUnderlinePatternDot;
        break;
    case 5:
        strikethroughStyle = NSUnderlinePatternDash;
        break;
    default:
        strikethroughStyle = NSUnderlineStyleNone;
    }
    dictionary[NSStrikethroughStyleAttributeName] = @(strikethroughStyle);
    
    dictionary[NSStrikethroughColorAttributeName] = [[UIColor alloc] initWithProtobuf:style.strikethroughColor];
    
    NSUnderlineStyle underlineStyle;
    switch (style.underlineStyle) {
    case 0:
        underlineStyle = NSUnderlineStyleNone;
        break;
    case 1: 
        underlineStyle = NSUnderlineStyleSingle;
        break;
    case 2:
        underlineStyle = NSUnderlineStyleDouble;
        break;
    case 3:
        underlineStyle = NSUnderlineStyleThick;
        break;
    case 4:
        underlineStyle = NSUnderlinePatternDot;
        break;
    case 5:
        underlineStyle = NSUnderlinePatternDash;
        break;
    default:
        underlineStyle = NSUnderlineStyleNone;
    }
    dictionary[NSUnderlineStyleAttributeName] = @(underlineStyle);
    
    dictionary[NSUnderlineColorAttributeName] = [[UIColor alloc] initWithProtobuf:style.underlineColor];
    dictionary[NSFontAttributeName] = [[UIFont alloc] initWithProtobuf:style.font];
    dictionary[NSHyphenationFactorDocumentAttribute] = @(style.hyphenation);
    paragraphStyle.lineHeightMultiple = style.lineHeightMultiple;
    // TODO(KD): AttributeKeyMaxLines
    dictionary[NSForegroundColorAttributeName] = [[UIColor alloc] initWithProtobuf:style.textColor];
    // TODO(KD): AttributeKeyTextWrap
    // TODO(KD): AttributeKeyTruncation
    // TODO(KD): AttributeKeyTruncationString

    return [[NSAttributedString alloc] initWithString:string attributes:dictionary];
}

@end

@implementation UIImage (Mochi)
- (id)initWithProtobuf:(MochiPBImage *)value {
    return [self initWithData:value.data_p];
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

- (id)initWithProtobuf:(MochiPBFont *)value {
    NSMutableDictionary *attr = [NSMutableDictionary dictionary];
    attr[UIFontDescriptorFamilyAttribute] = value.family;
    attr[UIFontDescriptorFaceAttribute] = value.face;
    attr[UIFontDescriptorSizeAttribute] = @(value.size);

    UIFontDescriptor *desc = [[UIFontDescriptor alloc] initWithFontAttributes:attr];
    UIFont *font = [UIFont fontWithDescriptor:desc size:0];
    return font;
}

@end

@implementation MochiPBRect (Mochi)

- (id)initWithCGRect:(CGRect)rect {
    if ((self = [super init])) {
        self.min = [[MochiPBPoint alloc] initWithCGPoint:CGPointMake(rect.origin.x, rect.origin.y)];
        self.max = [[MochiPBPoint alloc] initWithCGPoint:CGPointMake(rect.origin.x + rect.size.width, rect.origin.y + rect.size.height)];
    }
    return self;
}

- (CGRect)toCGRect {
    CGPoint min = self.min.toCGPoint;
    CGPoint max = self.max.toCGPoint;
    return CGRectMake(min.x, min.y, max.x - min.x, max.y - min.y);
}

@end
@implementation MochiPBPoint (Mochi)

- (id)initWithCGPoint:(CGPoint)point {
    if ((self = [super init])) {
        self.x = point.x;
        self.y = point.y;
    }
    return self;
}

- (CGPoint)toCGPoint {
    return CGPointMake(self.x, self.y);
}

- (id)initWithCGSize:(CGSize)size {
    if ((self = [super init])) {
        self.x = size.width;
        self.y = size.height;
    }
    return self;
}

- (CGSize)toCGSize {
    return CGSizeMake(self.x, self.y);
}

@end
@implementation MochiPBInsets (Mochi)
- (UIEdgeInsets)toUIEdgeInsets {
    return UIEdgeInsetsMake(self.top, self.left, self.bottom, self.right);
}
@end

@implementation GPBTimestamp (Mochi)
- (id)initWithDate:(NSDate *)date {
    if ((self = [super init])) {
        double integral;
        double fractional = modf([date timeIntervalSince1970], &integral);

        self.seconds = integral;
        self.nanos = fractional * NSEC_PER_SEC;
    }
    return self;
}

- (NSDate *)toDate {
    NSTimeInterval interval = (NSTimeInterval)self.seconds + (NSTimeInterval)self.nanos / NSEC_PER_SEC;
    return [NSDate dateWithTimeIntervalSince1970:interval];
}

@end
