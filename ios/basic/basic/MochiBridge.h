#import <UIKit/UIKit.h>
@import Mochi;

@interface MochiGoValue (Mochi)
- (id)initWithCGPoint:(CGPoint)point;
- (id)initWithCGSize:(CGSize)size;
- (id)initWithCGRect:(CGRect)rect;
- (CGPoint)toCGPoint;
- (CGSize)toCGSize;
- (CGRect)toCGRect;
- (UIEdgeInsets)toUIEdgeInsets;
@end

@interface NSMapTable (Mochi)
- (id)objectForKeyedSubscript:(id)key;
- (void)setObject:(id)obj forKeyedSubscript:(id)key;
@end
