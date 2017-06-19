#import <UIKit/UIKit.h>
#import "MatchaView.h"
@class MatchaViewNode;

@interface MatchaTabBarController : UITabBarController <MatchaChildViewController, UITabBarControllerDelegate>
- (id)initWithViewNode:(MatchaViewNode *)viewNode;
@property (nonatomic, weak) MatchaViewNode *viewNode;
@property (nonatomic, strong) MatchaNode *node;

// Private
@end
