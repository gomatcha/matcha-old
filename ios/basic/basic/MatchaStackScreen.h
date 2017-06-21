#import <UIKit/UIKit.h>
#import "MatchaView.h"
@class MatchaViewNode;
@class MatchaPBStackNavStackNav;

@interface MatchaStackScreen : UINavigationController <MatchaChildViewController, UINavigationControllerDelegate>
- (id)initWithViewNode:(MatchaViewNode *)viewNode;
@property (nonatomic, weak) MatchaViewNode *viewNode;
@property (nonatomic, strong) MatchaNode *node;

//Internal
@property (nonatomic, strong) NSArray *prev;
@end
