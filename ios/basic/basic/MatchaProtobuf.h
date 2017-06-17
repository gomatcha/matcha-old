#import <UIKit/UIKit.h>
#import "View.pbobjc.h"
#import "Layout.pbobjc.h"
#import "Text.pbobjc.h"
#import "Scrollview.pbobjc.h"
#import "Imageview.pbobjc.h"
#import "Button.pbobjc.h"
#import "Paint.pbobjc.h"
#import "Tabnav.pbobjc.h"
#import "Stacknavigator.pbobjc.h"
#import "Switchview.pbobjc.h"
#import "Touch2.pbobjc.h"
#import "Resource.pbobjc.h"
#import "Color.pbobjc.h"
#import "Textinput.pbobjc.h"

@import Matcha;

@interface UIColor (Matcha)
- (id)initWithProtobuf:(MatchaPBColor *)value;
- (MatchaPBColor *)protobuf;
@end

@interface NSAttributedString (Matcha)
- (id)initWithProtobuf:(MatchaPBStyledText *)value;
@end

@interface UIFont (Matcha)
- (id)initWithProtobuf:(MatchaPBFont *)value;
- (MatchaPBFont *)protobuf;
@end

@interface UIImage (Matcha)
- (id)initWithProtobuf:(MatchaPBImage *)value;
@end

@interface MatchaPBRect (Matcha)
- (id)initWithCGRect:(CGRect)rect;
@property (nonatomic, readonly) CGRect toCGRect;
@end

@interface MatchaPBPoint (Matcha)
- (id)initWithCGPoint:(CGPoint)point;
- (id)initWithCGSize:(CGSize)size;
@property (nonatomic, readonly) CGPoint toCGPoint;
@property (nonatomic, readonly) CGSize toCGSize;
@end

@interface MatchaPBInsets (Matcha)
@property (nonatomic, readonly) UIEdgeInsets toUIEdgeInsets;
@end

@interface GPBTimestamp (Matcha)
- (id)initWithDate:(NSDate *)date;
@property (nonatomic, readonly) NSDate *toDate;
@end
