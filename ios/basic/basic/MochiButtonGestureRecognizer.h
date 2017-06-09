#import <UIKit/UIKit.h>
#import "MochiProtobuf.h"
#import "MochiViewController.h"

@interface MochiButtonGestureRecognizer : UILongPressGestureRecognizer
- (id)initWithMochiVC:(MochiViewController *)viewRoot viewId:(int64_t)viewId protobuf:(GPBAny *)pb;
- (void)disable;
- (void)updateWithProtobuf:(GPBAny *)pb;
@end
