#import "MatchaStackScreen.h"
#import "MatchaView.h"
#import "MatchaProtobuf.h"
#import "MatchaViewController.h"

@implementation MatchaStackScreen

- (id)initWithViewNode:(MatchaViewNode *)viewNode {
    if ((self = [super init])) {
        self.viewNode = viewNode;
        self.delegate = self;
        MatchaConfigureChildViewController(self);
        self.view.backgroundColor = [UIColor whiteColor];
    }
    return self;
}

- (void)setMatchaChildViewControllers:(NSDictionary<NSNumber *, UIViewController *> *)childVCs {
    GPBAny *state = self.node.nativeViewState;
    
    MatchaStackScreenPBView *view = (id)[state unpackMessageClass:[MatchaStackScreenPBView class] error:nil];
    NSMutableArray *viewControllers = [NSMutableArray array];
    for (MatchaStackScreenPBChildView *i in view.childrenArray) {
        MatchaStackBar *bar = (id)childVCs[@(i.barId)];
        UIViewController *vc = childVCs[@(i.viewId)];
        vc.navigationItem.title = bar.titleString;
        vc.navigationItem.hidesBackButton = bar.backButtonHidden;
        vc.navigationItem.titleView = bar.titleView;
        vc.navigationItem.rightBarButtonItems = bar.rightViews;
        vc.navigationItem.leftBarButtonItems = bar.leftViews;
        vc.navigationItem.leftItemsSupplementBackButton = true;
        if (bar.customBackButtonTitle) {
            vc.navigationItem.backBarButtonItem = [[UIBarButtonItem alloc] initWithTitle:bar.backButtonTitle style:UIBarButtonItemStylePlain target:nil action:nil];
        }
        [viewControllers addObject:vc];
    }
    
    if (self.viewControllers.count == viewControllers.count) {
        [self setViewControllers:viewControllers animated:NO];
    } else {
        [self setViewControllers:viewControllers animated:YES];
    }
    self.prev = viewControllers;
}

//- (void)navigationController:(UINavigationController *)navigationController willShowViewController:(UIViewController *)viewController animated:(BOOL)animated {
//    NSLog(@"willShow");
//}

- (void)navigationController:(UINavigationController *)navigationController didShowViewController:(UIViewController *)viewController animated:(BOOL)animated {
    [self update];
}

- (void)update {
    // TODO(KD): More accurate comparison.
    if (self.viewControllers.count == self.prev.count) {
        return;
    }
    GPBInt64Array *array = [[GPBInt64Array alloc] init];
    for (NSInteger i = 0; i < self.viewControllers.count; i++) {
        [array addValue:0];
    }
    
    MatchaStackScreenPBStackEvent *event = [[MatchaStackScreenPBStackEvent alloc] init];
    event.idArray = array;
    
    NSData *data = [event data];
    MatchaGoValue *value = [[MatchaGoValue alloc] initWithData:data];
    
    [self.viewNode.rootVC call:@"OnChange" viewId:self.node.identifier.longLongValue args:@[value]];
}

@end

@implementation MatchaStackBar

- (id)initWithViewNode:(MatchaViewNode *)viewNode {
    if ((self = [super init])) {
        self.viewNode = viewNode;
    }
    return self;
}

- (void)setMatchaChildViewControllers:(NSDictionary<NSNumber *,UIViewController *> *)childVCs {
    GPBAny *state = self.node.nativeViewState;
    MatchaStackScreenPBBar *bar = (id)[state unpackMessageClass:[MatchaStackScreenPBBar class] error:nil];
    self.titleString = bar.title;
    self.backButtonHidden = bar.backButtonHidden;
    self.backButtonTitle = bar.backButtonTitle;
    self.customBackButtonTitle = bar.customBackButtonTitle;
    self.titleView = childVCs[@(bar.titleViewId)].view;
    if (self.titleView) {
        MatchaNode *n = self.node.nodeChildren[@(bar.titleViewId)];
        self.titleView.frame = n.guide.frame;
    }
    NSMutableArray *rightViews = [NSMutableArray array];
    for (NSInteger i = 0; i < bar.rightViewIdsArray.count; i++) {
        int64_t childId = [bar.rightViewIdsArray valueAtIndex:i];
        UIView *rightView = childVCs[@(childId)].view;
        MatchaNode *n = self.node.nodeChildren[@(childId)];
        rightView.frame = n.guide.frame;
        UIBarButtonItem *item = [[UIBarButtonItem alloc] initWithCustomView:rightView];
        [rightViews addObject:item];
    }
    self.rightViews = rightViews;
    
    NSMutableArray *leftViews = [NSMutableArray array];
    for (NSInteger i = 0; i < bar.leftViewIdsArray.count; i++) {
        int64_t childId = [bar.leftViewIdsArray valueAtIndex:i];
        UIView *leftView = childVCs[@(childId)].view;
        MatchaNode *n = self.node.nodeChildren[@(childId)];
        leftView.frame = n.guide.frame;
        UIBarButtonItem *item = [[UIBarButtonItem alloc] initWithCustomView:leftView];
        [leftViews addObject:item];
    }
    self.leftViews = leftViews;
}

@end
