#import <UIKit/UIKit.h>
#import "MochiView.h"
@class MochiViewNode;

@interface MochiTabBarController : UITabBarController <MochiChildViewController, UITabBarControllerDelegate>
- (id)initWithViewNode:(MochiViewNode *)viewNode;
@property (nonatomic, weak) MochiViewNode *viewNode;
@property (nonatomic, strong) MochiNode *node;

// Private
@property (nonatomic, assign) int64_t funcId;
@end
