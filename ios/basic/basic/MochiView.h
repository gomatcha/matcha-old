@import UIKit;
#import "MochiBridge.h"
#import "MochiNode.h"
@class MochiViewConfig;
@class MochiViewController;

@interface MochiBasicView : UIView
- (id)initWithViewRoot:(MochiViewController *)viewRoot parentVC:(UIViewController *)parentVC;
@property (nonatomic, weak) MochiViewController *viewRoot;
@property (nonatomic, weak) UIViewController *parentVC;
@property (nonatomic, strong) MochiNode *node;
@end

@interface MochiTextView : UILabel
- (id)initWithViewRoot:(MochiViewController *)viewRoot parentVC:(UIViewController *)parentVC;
@property (nonatomic, weak) MochiViewController *viewRoot;
@property (nonatomic, weak) UIViewController *parentVC;
@end

@interface MochiImageView : UIImageView
- (id)initWithViewRoot:(MochiViewController *)viewRoot parentVC:(UIViewController *)parentVC;
@property (nonatomic, weak) MochiViewController *viewRoot;
@property (nonatomic, weak) UIViewController *parentVC;
@end

@interface MochiButton : UIView
- (id)initWithViewRoot:(MochiViewController *)viewRoot parentVC:(UIViewController *)parentVC;
@property (nonatomic, weak) MochiViewController *viewRoot;
@property (nonatomic, weak) UIViewController *parentVC;
@end

@interface MochiScrollView : UIScrollView
- (id)initWithViewRoot:(MochiViewController *)viewRoot parentVC:(UIViewController *)parentVC;
@property (nonatomic, weak) MochiViewController *viewRoot;
@property (nonatomic, weak) UIViewController *parentVC;
@end


bool MochiConfigureViewWithNode(UIView *view, MochiNode *node, MochiViewConfig *config, MochiViewController *viewRoot);
UIGestureRecognizer *MochiGestureRecognizerWithPB(int64_t viewId, GPBAny *any, MochiViewController *viewRoot);
MochiBasicView *MochiViewWithNode(MochiNode *node, MochiViewController *root, UIViewController *parentVC);
UIViewController *MochiViewControllerWithNode(MochiNode *node, MochiViewController *root);

@interface MochiViewNode : NSObject
@property (nonatomic, strong) UIView *view;
@property (nonatomic, strong) UIViewController *viewController;
@property (nonatomic, strong) MochiNode *node;
@property (nonatomic, weak) MochiViewNode *parent;
@property (nonatomic, weak) MochiViewController *rootVC;
@end
