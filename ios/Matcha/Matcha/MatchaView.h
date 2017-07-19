#import <UIKit/UIKit.h>
#import <Matcha/MatchaNode.h>
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
- (void)setMatchaChildLayout:(GPBInt64ObjectDictionary *)layoutPaintNodes;
@end

typedef UIView<MatchaChildView> *(^MatchaViewRegistrationBlock)(MatchaViewNode *);
typedef UIViewController<MatchaChildViewController> *(^MatchaViewControllerRegistrationBlock)(MatchaViewNode *);

UIGestureRecognizer *MatchaGestureRecognizerWithPB(int64_t viewId, GPBAny *any, MatchaViewNode *viewNode);
UIView<MatchaChildView> *MatchaViewWithNode(MatchaNode *node, MatchaViewNode *viewNode);
UIViewController<MatchaChildViewController> *MatchaViewControllerWithNode(MatchaNode *node, MatchaViewNode *viewNode);
void MatchaRegisterView(NSString *string, MatchaViewRegistrationBlock block);
void MatchaRegisterViewController(NSString *string, MatchaViewControllerRegistrationBlock block);

@interface MatchaViewNode : NSObject
- (id)initWithParent:(MatchaViewNode *)node rootVC:(MatchaViewController *)rootVC;
@property (nonatomic, strong) UIView<MatchaChildView> *view;
@property (nonatomic, strong) NSDictionary<NSNumber *, UIGestureRecognizer *> *touchRecognizers;

@property (nonatomic, strong) UIViewController<MatchaChildViewController> *viewController;
@property (nonatomic, strong) NSDictionary<NSNumber *, MatchaViewNode *> *children;
- (void)setNode:(MatchaNode *)node root:(MatchaNodeRoot *)root;
@property (nonatomic, strong) MatchaNode *node;
@property (nonatomic, strong) MatchaLayoutPaintNode *layoutPaintNode;
@property (nonatomic, weak) MatchaViewNode *parent;
@property (nonatomic, weak) MatchaViewController *rootVC;

@property (nonatomic, strong) UIViewController *wrappedViewController;
- (UIViewController *)materializedViewController;
- (UIViewController *)wrappedViewController;
- (UIView *)materializedView;
@end
