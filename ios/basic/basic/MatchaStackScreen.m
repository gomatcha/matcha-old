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
    NSError *error = nil;
    
    MatchaPBStackNavStackNav *pb = (id)[state unpackMessageClass:[MatchaPBStackNavStackNav class] error:&error];
    NSMutableArray *viewControllers = [NSMutableArray array];
    NSMutableDictionary *vcDict = [NSMutableDictionary dictionary];
    for (MatchaPBStackNavScreen *i in pb.screensArray) {
        UIViewController *vc = childVCs[@(i.id_p)];
        vc.navigationItem.title = i.title;
        vc.navigationItem.hidesBackButton = i.backButtonHidden;
        if (i.customBackButtonTitle) {
            vc.navigationItem.backBarButtonItem = [[UIBarButtonItem alloc] initWithTitle:i.backButtonTitle style:UIBarButtonItemStylePlain target:nil action:nil];
        }
        [viewControllers addObject:vc];
    }
    
    if (self.viewControllers.count == viewControllers.count) {
        [self setViewControllers:viewControllers animated:NO];
    } else {
        [self setViewControllers:viewControllers animated:YES];
    }
    self.prev = pb;
    self.funcId = pb.eventFunc;
}

- (void)navigationController:(UINavigationController *)navigationController didShowViewController:(UIViewController *)viewController animated:(BOOL)animated {
    [self update];
}

- (void)update {
    GPBInt64Array *array = [[GPBInt64Array alloc] init];
    for (UIViewController *i in self.viewControllers) {
        [array addValue:0];
    }
    
    MatchaPBStackNavStackEvent *event = [[MatchaPBStackNavStackEvent alloc] init];
    event.idArray = array;
    
    NSData *data = [event data];
    MatchaGoValue *value = [[MatchaGoValue alloc] initWithData:data];
    
    [self.viewNode.rootVC call:self.funcId viewId:self.node.identifier.longLongValue args:@[value]];
}

@end
