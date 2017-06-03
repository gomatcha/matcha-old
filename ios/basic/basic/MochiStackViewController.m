#import "MochiStackViewController.h"
#import "MochiView.h"
#import "MochiProtobuf.h"

@implementation MochiStackViewController

- (id)initWithViewNode:(MochiViewNode *)viewNode {
    if ((self = [super init])) {
        self.viewNode = viewNode;
    }
    return self;
}

- (void)setMochiChildViewControllers:(NSDictionary<NSNumber *, UIViewController *> *)childVCs {
    GPBAny *state = self.node.nativeViewState;
    NSError *error = nil;
    MochiPBStackNavStackNav *pb = (id)[state unpackMessageClass:[MochiPBStackNavStackNav class] error:&error];
    NSMutableArray *viewControllers = [NSMutableArray array];
    for (MochiPBStackNavScreen *i in pb.screensArray) {
        UIViewController *vc = childVCs[@(i.id_p)];
        vc.navigationItem.title = i.title;
        vc.navigationItem.hidesBackButton = i.backButtonHidden;
        if (i.customBackButtonTitle) {
            vc.navigationItem.backBarButtonItem = [[UIBarButtonItem alloc] initWithTitle:i.backButtonTitle style:UIBarButtonItemStylePlain target:nil action:nil];
        }
        [viewControllers addObject:vc];
    }
    
    self.viewControllers = viewControllers;
}
@end
