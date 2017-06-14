#import <UIKit/UIKit.h>
@class MochiViewController;
@class GPBAny;

@interface MochiButtonGestureRecognizer : UIGestureRecognizer
- (id)initWithMochiVC:(MochiViewController *)viewRoot viewId:(int64_t)viewId protobuf:(GPBAny *)pb;
- (void)disable;
- (void)updateWithProtobuf:(GPBAny *)pb;
@end
