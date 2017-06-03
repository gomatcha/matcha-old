#import <UIKit/UIKit.h>
#import "MochiView.h"
@class MochiViewNode;

@interface MochiTabBarController : UITabBarController <MochiChildViewController>
- (id)initWithViewNode:(MochiViewNode *)viewNode;
@property (nonatomic, weak) MochiViewNode *viewNode;
@property (nonatomic, strong) MochiNode *node;
@end
