#import <UIKit/UIKit.h>
#import "MochiView.h"
@class MochiViewNode;

@interface MochiStackViewController : UINavigationController <MochiChildViewController, UINavigationControllerDelegate>
- (id)initWithViewNode:(MochiViewNode *)viewNode;
@property (nonatomic, weak) MochiViewNode *viewNode;
@property (nonatomic, strong) MochiNode *node;

//Internal
@property (nonatomic, assign) int64_t funcId;
@end
