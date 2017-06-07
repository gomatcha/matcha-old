@import UIKit;
@import Mochi;
@class MochiNode;

@interface MochiViewController : UIViewController // view.Root
+ (NSPointerArray *)viewControllers;
+ (MochiViewController *)viewControllerWithIdentifier:(NSInteger)identifier;

- (id)initWithGoValue:(MochiGoValue *)value;
- (void)update:(MochiNode *)node;
- (NSArray<MochiGoValue *> *)call:(int64_t)funcId viewId:(int64_t)viewId args:(NSArray<MochiGoValue *> *)args;
@property (nonatomic, readonly) NSInteger identifier;
@end


void MochiConfigureChildViewController(UIViewController *vc);
