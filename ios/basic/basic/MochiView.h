@import UIKit;
#import "MochiBridge.h"
#import "MochiNode.h"
@class MochiViewConfig;
@class MochiViewController;
@class MochiViewNode;

@protocol MochiChildView <NSObject>
- (void)setNode:(MochiNode *)node;
@end

@interface MochiBasicView : UIView <MochiChildView>
- (id)initWithViewRoot:(MochiViewController *)viewRoot parentVC:(UIViewController *)parentVC;
@property (nonatomic, weak) MochiViewController *viewRoot;
@property (nonatomic, weak) UIViewController *parentVC;
@property (nonatomic, strong) MochiNode *node;
@end

@interface MochiTextView : UILabel <MochiChildView>
- (id)initWithViewRoot:(MochiViewController *)viewRoot parentVC:(UIViewController *)parentVC;
@property (nonatomic, weak) MochiViewController *viewRoot;
@property (nonatomic, weak) UIViewController *parentVC;
@end

@interface MochiImageView : UIImageView <MochiChildView>
@end

@interface MochiButton : UIView <MochiChildView>
@end

@interface MochiScrollView : UIScrollView <MochiChildView>
- (id)initWithViewRoot:(MochiViewController *)viewRoot parentVC:(UIViewController *)parentVC;
@property (nonatomic, weak) MochiViewController *viewRoot;
@property (nonatomic, weak) UIViewController *parentVC;
@end


bool MochiConfigureViewWithNode(UIView *view, MochiNode *node, MochiViewConfig *config, MochiViewController *viewRoot);
UIGestureRecognizer *MochiGestureRecognizerWithPB(int64_t viewId, GPBAny *any, MochiViewController *viewRoot);
UIView<MochiChildView> *MochiViewWithNode(MochiNode *node, MochiViewController *root, UIViewController *parentVC, MochiViewNode *viewNode);
UIViewController *MochiViewControllerWithNode(MochiNode *node, MochiViewController *root);

@interface MochiViewNode : NSObject
- (id)initWithParent:(MochiViewNode *)node rootVC:(MochiViewController *)rootVC;
@property (nonatomic, strong) UIView<MochiChildView> *view;
// @property (nonatomic, strong) UIViewController *viewController;
@property (nonatomic, strong) NSDictionary<NSNumber *, MochiViewNode *> *children;
@property (nonatomic, strong) MochiNode *node;
@property (nonatomic, weak) MochiViewNode *parent;
@property (nonatomic, weak) MochiViewController *rootVC;
@end
