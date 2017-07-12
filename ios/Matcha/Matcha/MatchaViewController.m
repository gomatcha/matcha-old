#import "MatchaViewController.h"
#import "MatchaView.h"
#import "MatchaBridge.h"
#import "MatchaNode.h"

@interface MatchaViewController ()
@property (nonatomic, assign) NSInteger identifier;
@property (nonatomic, strong) MatchaViewNode *viewNode;
@property (nonatomic, strong) MatchaGoValue *goValue;
@property (nonatomic, assign) CGRect lastFrame;
@property (nonatomic, assign) BOOL loaded;
@end

@implementation MatchaViewController

+ (NSPointerArray *)viewControllers {
    static NSPointerArray *sPointerArray;
    static dispatch_once_t sOnce;
    dispatch_once(&sOnce, ^{
        sPointerArray = [NSPointerArray weakObjectsPointerArray];
    });
    return sPointerArray;
}

+ (MatchaViewController *)viewControllerWithIdentifier:(NSInteger)identifier {
    for (MatchaViewController *i in [self viewControllers]) {
        if (i.identifier == identifier) {
            return i;
        }
    }
    return nil;
}
- (id)initWithGoValue:(MatchaGoValue *)value {
    if ((self = [super initWithNibName:nil bundle:nil])) {
        self.goValue = value;
        self.identifier = (int)[value call:@"Id" args:nil][0].toLongLong;
        [[MatchaViewController viewControllers] addPointer:(__bridge void *)self];
        self.viewNode = [[MatchaViewNode alloc] initWithParent:nil rootVC:self];
        self.edgesForExtendedLayout = UIRectEdgeNone;
        self.extendedLayoutIncludesOpaqueBars=NO;
        self.automaticallyAdjustsScrollViewInsets=NO;
    }
    return self;
}

- (void)viewDidLayoutSubviews {
    if (!CGRectEqualToRect(self.lastFrame, self.view.frame)) {
        self.lastFrame = self.view.frame;
        
        [self.goValue call:@"SetSize" args:@[[[MatchaGoValue alloc] initWithCGPoint:CGPointMake(self.view.frame.size.width, self.view.frame.size.height)]]];
    }
}

- (NSArray<MatchaGoValue *> *)call:(NSString *)funcId viewId:(int64_t)viewId args:(NSArray<MatchaGoValue *> *)args {
    MatchaGoValue *goValue = [[MatchaGoValue alloc] initWithString:funcId];
    MatchaGoValue *goViewId = [[MatchaGoValue alloc] initWithLongLong:viewId];
    MatchaGoValue *goArgs = [[MatchaGoValue alloc] initWithArray:args];
    return [self.goValue call:@"Call" args:@[goValue, goViewId, goArgs]];
}


- (void)update:(MatchaNode *)node {
    self.viewNode.node = node;
    if (!self.loaded) {
        self.loaded = TRUE;
        UIView *view = self.viewNode.view ?: self.viewNode.viewController.view;
        [self.view addSubview:view];
        self.view.autoresizingMask = UIViewAutoresizingFlexibleWidth|UIViewAutoresizingFlexibleHeight;
        view.frame = self.view.bounds;
    }
}

@end

void MatchaConfigureChildViewController(UIViewController *vc) {
    vc.edgesForExtendedLayout=UIRectEdgeNone;
    vc.extendedLayoutIncludesOpaqueBars=NO;
    vc.automaticallyAdjustsScrollViewInsets=NO;
}
