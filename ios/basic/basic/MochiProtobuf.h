#import <UIKit/UIKit.h>
#import "View.pbobjc.h"
#import "Layout.pbobjc.h"
#import "Text.pbobjc.h"
#import "Scrollview.pbobjc.h"
#import "Imageview.pbobjc.h"
#import "Button.pbobjc.h"
#import "Touch.pbobjc.h"
#import "Paint.pbobjc.h"
#import "Tabnavigator.pbobjc.h"
#import "Stacknavigator.pbobjc.h"

@import Mochi;

@interface UIColor (Mochi)
- (id)initWithProtobuf:(MochiPBColor *)value;
@end

@interface NSAttributedString (Mochi)
- (id)initWithProtobuf:(MochiPBText *)value;
@end

@interface UIFont (Mochi)
- (id)initWithProtobuf:(MochiPBFont *)value;
@end

@interface UIImage (Mochi)
- (id)initWithProtobuf:(MochiPBImage *)value;
@end

@interface MochiPBRect (Mochi)
- (id)initWithCGRect:(CGRect)rect;
@property (nonatomic, readonly) CGRect toCGRect;
@end

@interface MochiPBPoint (Mochi)
- (id)initWithCGPoint:(CGPoint)point;
- (id)initWithCGSize:(CGSize)size;
@property (nonatomic, readonly) CGPoint toCGPoint;
@property (nonatomic, readonly) CGSize toCGSize;
@end

@interface MochiPBInsets (Mochi)
@property (nonatomic, readonly) UIEdgeInsets toUIEdgeInsets;
@end

@interface GPBTimestamp (Mochi)
- (id)initWithDate:(NSDate *)date;
@property (nonatomic, readonly) NSDate *toDate;
@end
