#import <UIKit/UIKit.h>
#import <MatchaBridge/MatchaBridge.h>
@class MatchaNode;

@interface MatchaViewController : UIViewController // view.Root
+ (NSPointerArray *)viewControllers;
+ (MatchaViewController *)viewControllerWithIdentifier:(NSInteger)identifier;

- (id)initWithGoValue:(MatchaGoValue *)value;
- (void)update:(MatchaNode *)node;
- (NSArray<MatchaGoValue *> *)call:(NSString *)funcId viewId:(int64_t)viewId args:(NSArray<MatchaGoValue *> *)args;
@property (nonatomic, readonly) NSInteger identifier;
@end


void MatchaConfigureChildViewController(UIViewController *vc);
