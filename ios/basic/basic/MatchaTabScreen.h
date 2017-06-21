#import <UIKit/UIKit.h>
#import "MatchaView.h"
@class MatchaViewNode;

@interface MatchaTabScreen : UITabBarController <MatchaChildViewController, UITabBarControllerDelegate>
- (id)initWithViewNode:(MatchaViewNode *)viewNode;
@property (nonatomic, weak) MatchaViewNode *viewNode;
@property (nonatomic, strong) MatchaNode *node;

// Private
@end
