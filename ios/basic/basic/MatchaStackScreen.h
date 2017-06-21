#import <UIKit/UIKit.h>
#import "MatchaView.h"
@class MatchaViewNode;

@interface MatchaStackScreen : UINavigationController <MatchaChildViewController, UINavigationControllerDelegate>
- (id)initWithViewNode:(MatchaViewNode *)viewNode;
@property (nonatomic, weak) MatchaViewNode *viewNode;
@property (nonatomic, strong) MatchaNode *node;

//Internal
@property (nonatomic, strong) NSArray *prev;
@end

@interface MatchaStackBar : UIViewController <MatchaChildViewController>
- (id)initWithViewNode:(MatchaViewNode *)viewNode;
@property (nonatomic, weak) MatchaViewNode *viewNode;
@property (nonatomic, strong) MatchaNode *node;

//Internal
@property (nonatomic, strong) NSString *titleString;
@property (nonatomic, assign) BOOL backButtonHidden;
@property (nonatomic, assign) BOOL customBackButtonTitle;
@property (nonatomic, assign) NSString *backButtonTitle;
@property (nonatomic, strong) UIView *titleView;
@property (nonatomic, strong) NSArray *rightViews;
@property (nonatomic, strong) NSArray *leftViews;
@end
