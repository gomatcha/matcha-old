#import "MatchaStackScreen.h"
#import "MatchaView.h"
#import "MatchaProtobuf.h"
#import "MatchaViewController.h"
#import <objc/runtime.h>

#define VIEW_ID_KEY @"matchaViewId"

@interface UIViewController (MatchaStackScreen)
- (void)matcha_setViewId:(int64_t)value;
- (int64_t)matcha_viewId;
@end

@implementation UIViewController (MatchaStackScreen)

- (void)matcha_setViewId:(int64_t)value {
    @synchronized (self) {
        objc_setAssociatedObject(self, VIEW_ID_KEY, @(value), OBJC_ASSOCIATION_RETAIN);
    }
}

- (int64_t)matcha_viewId {
    @synchronized (self) {
        return ((NSNumber *)objc_getAssociatedObject(self, VIEW_ID_KEY)).longLongValue;
    }
}

@end

@implementation MatchaStackScreen

+ (void)load {
    MatchaRegisterViewController(@"gomatcha.io/matcha/view/stacknav", ^(MatchaViewNode *node){
        return [[MatchaStackScreen alloc] initWithViewNode:node];
    });
    MatchaRegisterViewController(@"gomatcha.io/matcha/view/stacknav Bar", ^(MatchaViewNode *node){
        return [[MatchaStackBar alloc] initWithViewNode:node];
    });
}

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
    MatchaStackScreenPBView *view = (id)[self.node.nativeViewState unpackMessageClass:[MatchaStackScreenPBView class] error:nil];
    
    NSMutableArray *prevIds = [NSMutableArray array];
    for (MatchaStackScreenPBChildView *i in view.childrenArray) {
        [prevIds addObject:@(i.viewId)];
    }
    if ([self.prevIds isEqual:prevIds]) {
        return;
    }
    self.prevIds = prevIds;
    
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
        [vc matcha_setViewId:i.screenId];
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
    NSMutableArray *prevIds = [NSMutableArray array];
    for (UIViewController *i in self.childViewControllers) {
        [prevIds addObject:@(i.matcha_viewId)];
    }
    if ([self.prevIds isEqual:prevIds]) {
        return;
    }
    self.prevIds = prevIds;
    
    GPBInt64Array *array = [[GPBInt64Array alloc] init];
    for (NSNumber *i in prevIds) {
        [array addValue:i.longLongValue];
    }
    MatchaStackScreenPBStackEvent *event = [[MatchaStackScreenPBStackEvent alloc] init];
    event.idArray = array;
    
    MatchaGoValue *value = [[MatchaGoValue alloc] initWithData:event.data];
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
