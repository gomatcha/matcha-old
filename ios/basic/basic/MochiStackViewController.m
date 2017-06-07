#import "MochiStackViewController.h"
#import "MochiView.h"
#import "MochiProtobuf.h"
#import "MochiViewController.h"

@implementation MochiStackViewController

- (id)initWithViewNode:(MochiViewNode *)viewNode {
    if ((self = [super init])) {
        self.viewNode = viewNode;
        self.delegate = self;
        MochiConfigureChildViewController(self);
        self.view.backgroundColor = [UIColor whiteColor];
    }
    return self;
}

- (void)setMochiChildViewControllers:(NSDictionary<NSNumber *, UIViewController *> *)childVCs {
    GPBAny *state = self.node.nativeViewState;
    NSError *error = nil;
    
    MochiPBStackNavStackNav *pb = (id)[state unpackMessageClass:[MochiPBStackNavStackNav class] error:&error];
    NSMutableArray *viewControllers = [NSMutableArray array];
    NSMutableDictionary *vcDict = [NSMutableDictionary dictionary];
    for (MochiPBStackNavScreen *i in pb.screensArray) {
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
    
    MochiPBStackNavStackEvent *event = [[MochiPBStackNavStackEvent alloc] init];
    event.idArray = array;
    
    NSData *data = [event data];
    MochiGoValue *value = [[MochiGoValue alloc] initWithData:data];
    
    [self.viewNode.rootVC call:self.funcId viewId:self.node.identifier.longLongValue args:@[value]];
}

@end
