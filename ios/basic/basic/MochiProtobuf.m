#import <Foundation/Foundation.h>
#import "MochiProtobuf.h"

@implementation UIColor (Mochi)

- (id)initWithProtobuf:(MochiPBColor *)value {
    if (value == nil) {
        return nil;
    }
    return [UIColor colorWithRed:((double)value.red)/0xffff green:((double)value.green)/0xffff blue:((double)value.blue)/0xffff alpha:((double)value.alpha)/0xffff];
}

@end

@implementation NSAttributedString (Mochi)

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
