#import "MatchaTabBarController.h"
#import "MatchaView.h"
#import "MatchaProtobuf.h"
#import "MatchaViewController.h"

@implementation MatchaTabBarController

- (id)initWithViewNode:(MatchaViewNode *)viewNode {
    if ((self = [super init])) {
        self.viewNode = viewNode;
        self.delegate = self;
        MatchaConfigureChildViewController(self);
    }
    return self;
}

- (void)setMatchaChildViewControllers:(NSDictionary<NSNumber *, UIViewController *> *)childVCs {
    GPBAny *state = self.node.nativeViewState;
    NSError *error = nil;
    MatchaPBTabNavTabNav *pbTabNavigator = (id)[state unpackMessageClass:[MatchaPBTabNavTabNav class] error:&error];
    NSMutableArray *viewControllers = [NSMutableArray array];
    for (MatchaPBTabNavScreen *i in pbTabNavigator.screensArray) {
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
    MatchaPBTabNavEvent *event = [[MatchaPBTabNavEvent alloc] init];
    event.selectedIndex = tabBarController.selectedIndex;
    
    NSData *data = [event data];
    MatchaGoValue *value = [[MatchaGoValue alloc] initWithData:data];
    
    [self.viewNode.rootVC call:self.funcId viewId:self.node.identifier.longLongValue args:@[value]];
}

@end
