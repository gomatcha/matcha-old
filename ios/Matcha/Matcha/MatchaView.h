#import <UIKit/UIKit.h>
#import "MatchaBridge.h"
#import "MatchaNode.h"
@class MatchaViewConfig;
@class MatchaViewController;
@class MatchaViewNode;

@protocol MatchaChildView <NSObject>
- (id)initWithViewNode:(MatchaViewNode *)viewNode;
- (void)setNode:(MatchaNode *)node;
@end

@protocol MatchaChildViewController <NSObject>
- (id)initWithViewNode:(MatchaViewNode *)viewNode;
- (void)setNode:(MatchaNode *)node;
- (void)setMatchaChildViewControllers:(NSDictionary<NSNumber *, UIViewController *> *)childVCs;
@end

UIGestureRecognizer *MatchaGestureRecognizerWithPB(int64_t viewId, GPBAny *any, MatchaViewNode *viewNode);
UIView<MatchaChildView> *MatchaViewWithNode(MatchaNode *node, MatchaViewNode *viewNode);
UIViewController<MatchaChildViewController> *MatchaViewControllerWithNode(MatchaNode *node, MatchaViewNode *viewNode);

@interface MatchaViewNode : NSObject
- (id)initWithParent:(MatchaViewNode *)node rootVC:(MatchaViewController *)rootVC;
@property (nonatomic, strong) UIView<MatchaChildView> *view;
@property (nonatomic, strong) NSDictionary<NSNumber *, UIGestureRecognizer *> *touchRecognizers;

@property (nonatomic, strong) UIViewController<MatchaChildViewController> *viewController;
@property (nonatomic, strong) NSDictionary<NSNumber *, MatchaViewNode *> *children;
@property (nonatomic, strong) MatchaNode *node;
@property (nonatomic, weak) MatchaViewNode *parent;
@property (nonatomic, weak) MatchaViewController *rootVC;

@property (nonatomic, strong) UIViewController *wrappedViewController;
- (UIViewController *)materializedViewController;
- (UIViewController *)wrappedViewController;
- (UIView *)materializedView;
@end
