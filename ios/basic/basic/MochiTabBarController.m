#import "MochiTabBarController.h"
#import "MochiView.h"
#import "MochiProtobuf.h"
#import "MochiViewController.h"

@implementation MochiTabBarController

- (id)initWithViewNode:(MochiViewNode *)viewNode {
    if ((self = [super init])) {
        self.viewNode = viewNode;
        self.delegate = self;
    }
    return self;
}

- (void)setMochiChildViewControllers:(NSDictionary<NSNumber *, UIViewController *> *)childVCs {
    GPBAny *state = self.node.nativeViewState;
    NSError *error = nil;
    MochiPBTabNavTabNav *pbTabNavigator = (id)[state unpackMessageClass:[MochiPBTabNavTabNav class] error:&error];
    NSMutableArray *viewControllers = [NSMutableArray array];
    for (MochiPBTabNavScreen *i in pbTabNavigator.screensArray) {
        UIViewController *vc = childVCs[@(i.id_p)];
        vc.tabBarItem.title = i.title;
        vc.tabBarItem.badgeValue = i.badge.length == 0 ? nil : i.badge;
        [viewControllers addObject:vc];
    }
    
    self.viewControllers = viewControllers;
    self.selectedIndex = pbTabNavigator.selectedIndex;
    self.funcId = pbTabNavigator.eventFunc;
}

- (void)tabBarController:(UITabBarController *)tabBarController didSelectViewController:(UIViewController *)viewController {
    MochiPBTabNavEvent *event = [[MochiPBTabNavEvent alloc] init];
    event.selectedIndex = tabBarController.selectedIndex;
    
    NSData *data = [event data];
    MochiGoValue *value = [[MochiGoValue alloc] initWithData:data];
    
    [self.viewNode.rootVC call:self.funcId viewId:self.node.identifier.longLongValue args:@[value]];
}

@end
