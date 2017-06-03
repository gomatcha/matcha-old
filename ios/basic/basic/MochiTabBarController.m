#import "MochiTabBarController.h"
#import "MochiView.h"
#import "MochiProtobuf.h"

@implementation MochiTabBarController

- (id)initWithViewNode:(MochiViewNode *)viewNode {
    if ((self = [super init])) {
        self.viewNode = viewNode;
    }
    return self;
}

//
//- (void)setNode:(MochiNode *)value {
//    _node = value;
//    GPBAny *state = value.nativeViewState;
//    NSError *error = nil;
//    MochiPBButtonButton *pbbutton = (id)[state unpackMessageClass:[MochiPBButtonButton class] error:&error];
//    
//    NSAttributedString *string = [[NSAttributedString alloc] initWithProtobuf:pbbutton.text];
//    [self.button setAttributedTitle:string forState:UIControlStateNormal];
//}


- (void)setMochiChildViewControllers:(NSDictionary<NSNumber *, UIViewController *> *)childVCs {
    GPBAny *state = self.node.nativeViewState;
    NSError *error = nil;
    MochiPBTabNavigatorTabNavigator *pbTabNavigator = (id)[state unpackMessageClass:[MochiPBTabNavigatorTabNavigator class] error:&error];
    NSMutableArray *viewControllers = [NSMutableArray array];
    for (MochiPBTabNavigatorTab *i in pbTabNavigator.tabsArray) {
        UIViewController *vc = childVCs[@(i.id_p)];
        vc.tabBarItem.title = i.title;
        vc.tabBarItem.badgeValue = i.badge.length == 0 ? nil : i.badge;
        [viewControllers addObject:vc];
    }
    
    self.viewControllers = viewControllers;
}

@end
