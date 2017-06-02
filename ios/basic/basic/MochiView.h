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
@end

@interface MochiTextView : UILabel <MochiChildView>
@end

@interface MochiImageView : UIImageView <MochiChildView>
@end

@interface MochiButton : UIView <MochiChildView>
@end

@interface MochiScrollView : UIScrollView <MochiChildView>
@end

UIGestureRecognizer *MochiGestureRecognizerWithPB(int64_t viewId, GPBAny *any, MochiViewNode *viewNode);
UIView<MochiChildView> *MochiViewWithNode(MochiNode *node, MochiViewNode *viewNode);
UIViewController *MochiViewControllerWithNode(MochiNode *node, MochiViewController *root);

@interface MochiViewNode : NSObject
- (id)initWithParent:(MochiViewNode *)node rootVC:(MochiViewController *)rootVC;
@property (nonatomic, strong) UIView<MochiChildView> *view;
@property (nonatomic, strong) NSDictionary<NSNumber *, UIGestureRecognizer *> *touchRecognizers;

// @property (nonatomic, strong) UIViewController *viewController;
@property (nonatomic, strong) NSDictionary<NSNumber *, MochiViewNode *> *children;
@property (nonatomic, strong) MochiNode *node;
@property (nonatomic, weak) MochiViewNode *parent;
@property (nonatomic, weak) MochiViewController *rootVC;
@end
