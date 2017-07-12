#import "MatchaView.h"
#import "MatchaProtobuf.h"
#import "MatchaTapGestureRecognizer.h"
#import "MatchaPressGestureRecognizer.h"
#import "MatchaTabScreen.h"
#import "MatchaViewController.h"
#import "MatchaStackScreen.h"
#import "MatchaSwitchView.h"
#import "MatchaButtonGestureRecognizer.h"
#import "MatchaTextInput.h"
#import "MatchaScrollView.h"
#import "MatchaButton.h"
#import "MatchaSlider.h"
#import "MatchaBasicView.h"
#import "MatchaTextView.h"
#import "MatchaImageView.h"
#import "MatchaProgressView.h"
#import "MatchaSegmentView.h"

static NSLock *sLock = nil;
static NSMutableDictionary *sDict = nil;

void MatchaRegisterView(NSString *string, MatchaViewRegistrationBlock block) {
    static dispatch_once_t sOnce = 0;
    dispatch_once(&sOnce, ^{
        sLock = [[NSLock alloc] init];
        sDict = [NSMutableDictionary dictionary];
    });
    
    [sLock lock];
    sDict[string] = block;
    [sLock unlock];
}

UIGestureRecognizer *MatchaGestureRecognizerWithPB(int64_t viewId, GPBAny *any, MatchaViewNode *viewNode) {
    if ([any.typeURL isEqual:@"type.googleapis.com/matcha.touch.TapRecognizer"]) {
        return [[MatchaTapGestureRecognizer alloc] initWithMatchaVC:viewNode.rootVC viewId:viewId protobuf:any];
    } else if ([any.typeURL isEqual:@"type.googleapis.com/matcha.touch.PressRecognizer"]) {
        return [[MatchaPressGestureRecognizer alloc] initWithMatchaVC:viewNode.rootVC viewId:viewId protobuf:any];
    } else if ([any.typeURL isEqual:@"type.googleapis.com/matcha.touch.ButtonRecognizer"]) {
        return [[MatchaButtonGestureRecognizer alloc] initWithMatchaVC:viewNode.rootVC viewId:viewId protobuf:any];
    }
    return nil;
}

UIView<MatchaChildView> *MatchaViewWithNode(MatchaNode *node, MatchaViewNode *viewNode) {
    NSString *name = node.nativeViewName;
    UIView<MatchaChildView> *child = nil;
    
    [sLock lock];
    MatchaViewRegistrationBlock block = sDict[name];
    if (block != nil) {
        child = block(viewNode);
    }
    [sLock unlock];

    return child;
}

UIViewController<MatchaChildViewController> *MatchaViewControllerWithNode(MatchaNode *node, MatchaViewNode *viewNode) {
    NSString *name = node.nativeViewName;
    UIViewController<MatchaChildViewController> *child = nil;
    if ([name isEqual:@"gomatcha.io/matcha/view/tabscreen"]) {
        child = [[MatchaTabScreen alloc] initWithViewNode:viewNode];
    } else if ([name isEqual:@"gomatcha.io/matcha/view/stacknav"]) {
        child = [[MatchaStackScreen alloc] initWithViewNode:viewNode];
    } else if ([name isEqual:@"gomatcha.io/matcha/view/stacknav Bar"]) {
        child = [[MatchaStackBar alloc] initWithViewNode:viewNode];
    }
    return child;
}

@implementation MatchaViewNode

- (id)initWithParent:(MatchaViewNode *)node rootVC:(MatchaViewController *)rootVC {
    if ((self = [super init])) {
        self.parent = node;
        self.rootVC = rootVC; 
    }
    return self;
}

- (void)setNode:(MatchaNode *)node {
    NSAssert(self.node == nil || [self.node.nativeViewName isEqual:node.nativeViewName], @"Node with different name");
    
    if (self.view == nil && self.viewController == nil) {
        self.view = MatchaViewWithNode(node, self);
        self.viewController = MatchaViewControllerWithNode(node, self);
        if (self.view == nil && self.viewController == nil) {
            NSLog(@"Cannot find corresponding view or view controller for node: %@", node.nativeViewName);
        }
    }
    
    // Build children
    NSDictionary<NSNumber *, MatchaViewNode *> *children = nil;
    NSMutableArray *addedKeys = [NSMutableArray array];
    NSMutableArray *removedKeys = [NSMutableArray array];
    NSMutableArray *unmodifiedKeys = [NSMutableArray array];
    if (![node.buildId isEqual:self.node.buildId]) {
        for (NSNumber *i in self.children) {
            MatchaNode *child = node.nodeChildren[i];
            if (child == nil) {
                [removedKeys addObject:i];
            }
        }
        for (NSNumber *i in node.nodeChildren) {
            MatchaViewNode *prevChild = self.children[i];
            if (prevChild == nil) {
                [addedKeys addObject:i];
            } else {
                [unmodifiedKeys addObject:i];
            }
        }
        
        // Add/remove child nodes
        NSMutableDictionary<NSNumber *, MatchaViewNode *> *mutChildren = [NSMutableDictionary dictionary];
        for (NSNumber *i in addedKeys) {
            mutChildren[i] = [[MatchaViewNode alloc] initWithParent:self rootVC:self.rootVC];
        }
        for (NSNumber *i in unmodifiedKeys) {
            mutChildren[i] = self.children[i];
        }
        children = mutChildren;
    } else {
        children = self.children;
    }
    
    // Update children
    for (NSNumber *i in children) {
        MatchaViewNode *child = children[i];
        child.node = node.nodeChildren[i];
    }
    
    if (![node.buildId isEqual:self.node.buildId]) {
        // Update the views with native values
        if (self.view) {
            self.view.node = node;
        } else if (self.viewController) {
            self.viewController.node = node;
            
            NSMutableDictionary<NSNumber *, UIViewController *> *childVCs = [NSMutableDictionary dictionary];
            for (NSNumber *i in children) {
                MatchaViewNode *child = children[i];
                childVCs[i] = child.wrappedViewController;
            }
            self.viewController.matchaChildViewControllers = childVCs;
        }
        
        // Add/remove subviews
        for (NSNumber *i in addedKeys) {
            MatchaViewNode *child = children[i];
            child.view.node = node.nodeChildren[i];
            
            if (self.viewController) {
                // no-op. The view controller will handle this itself.
            } else if (child.view) {
                [self.materializedView addSubview:child.view];
            } else if (child.viewController) {
//                [self.materializedViewController addChildViewController:child.viewController]; // TODO(KD): Why can't I add as a child view controller?
                [self.materializedView addSubview:child.viewController.view];
            }
        }
        for (NSNumber *i in removedKeys) {
            MatchaViewNode *child = self.children[i];
            if (self.viewController) {
                // no-op
            } else if (child.view) {
                [child.view removeFromSuperview];
            } else if (child.viewController) {
                [child.view removeFromSuperview];
                [child.viewController removeFromParentViewController];
            }
        }
    }
    
    // Update gesture recognizers
    {
        if (self.view) {
            NSMutableArray *addedKeys = [NSMutableArray array];
            NSMutableArray *removedKeys = [NSMutableArray array];
            NSMutableArray *unmodifiedKeys = [NSMutableArray array];
            for (NSNumber *i in self.node.touchRecognizers) {
                GPBAny *child = node.touchRecognizers[i];
                if (child == nil) {
                    [removedKeys addObject:i];
                }
            }
            for (NSNumber *i in node.touchRecognizers) {
                GPBAny *prevChild = self.node.touchRecognizers[i];
                if (prevChild == nil) {
                    [addedKeys addObject:i];
                } else {
                    [unmodifiedKeys addObject:i];
                }
            }
            
            NSMutableDictionary *touchRecognizers = [NSMutableDictionary dictionary];
            for (NSNumber *i in removedKeys) {
                UIGestureRecognizer *recognizer = self.touchRecognizers[i];
                [(id)recognizer disable];
                [self.view removeGestureRecognizer:recognizer];
            }
            for (NSNumber *i in addedKeys) {
                UIGestureRecognizer *recognizer = MatchaGestureRecognizerWithPB(node.identifier.longLongValue, node.touchRecognizers[i], self);
                [self.view addGestureRecognizer:recognizer];
                touchRecognizers[i] = recognizer;
            }
            for (NSNumber *i in unmodifiedKeys) {
                UIGestureRecognizer *recognizer = self.touchRecognizers[i];
                [(id)recognizer updateWithProtobuf:node.touchRecognizers[i]];
                touchRecognizers[i] = recognizer;
            }
            self.touchRecognizers = touchRecognizers;
        }
    }

    // Layout subviews
    if (![node.layoutId isEqual:self.node.layoutId]) {
        if (self.view) {
            NSArray *sortedKeys = [[children allKeys] sortedArrayUsingComparator:^NSComparisonResult(NSNumber *obj1, NSNumber *obj2) {
                return node.nodeChildren[obj1].guide.zIndex > node.nodeChildren[obj2].guide.zIndex;
            }];
            
            for (NSInteger i = 0; i < sortedKeys.count; i++) {
                NSNumber *key = sortedKeys[i];
                UIView *subview = children[key].view;
                if ([self.view.subviews indexOfObject:subview] != i) {
                    [self.view insertSubview:subview atIndex:i];
                }
            }
        }
        
        
        if ([self.parent.view isKindOfClass:[MatchaScrollView class]]) {
            MatchaScrollView *scrollView = (MatchaScrollView *)self.parent.view;
            bool scrollEvents = scrollView.scrollEvents;
            scrollView.scrollEvents = false;
            
            CGRect frame = node.guide.frame;
            frame.origin = CGPointZero;
            self.materializedView.frame = frame;
            self.materializedView.autoresizingMask = UIViewAutoresizingNone;
            [scrollView setContentOffset:node.guide.frame.origin];
            
            scrollView.scrollEvents = scrollEvents;
            
        } else if (self.parent.viewController == nil) {
            // let view controllers do their own layout
            self.materializedView.frame = node.guide.frame;
            self.materializedView.autoresizingMask = UIViewAutoresizingNone;
        }
    }
    
    // Paint view
    if (![node.paintId isEqual:self.node.paintId]) {
        MatchaPaintOptions *paintOptions = node.paintOptions;
        self.view.alpha = 1 - paintOptions.transparency;
        self.view.backgroundColor = paintOptions.backgroundColor;
        self.view.layer.borderColor = paintOptions.borderColor.CGColor;
        self.view.layer.borderWidth = paintOptions.borderWidth;
        self.view.layer.cornerRadius = paintOptions.cornerRadius;
        self.view.layer.shadowRadius = paintOptions.shadowRadius;
        self.view.layer.shadowOffset = paintOptions.shadowOffset;
        self.view.layer.shadowColor = paintOptions.shadowColor.CGColor;
        self.view.layer.shadowOpacity = paintOptions.shadowColor == nil ? 0 : 1;
        if (paintOptions.cornerRadius != 0) {
            self.view.clipsToBounds = YES; // TODO(KD): Be better about this...
        }
    }
    
    _node = node;
    self.children = children;
}

- (UIViewController *)materializedViewController {
    UIViewController *vc = nil;
    MatchaViewNode *viewNode = self;
    while (vc == nil && viewNode != nil) {
        viewNode = self.parent;
        vc = viewNode.viewController;
    }
    if (vc == nil) {
        vc = self.rootVC;
    }
    return vc;
}

- (UIViewController *)wrappedViewController {
    if (_wrappedViewController) {
        return _wrappedViewController;
    }
    
    if (self.viewController) {
        _wrappedViewController = self.viewController;
        return _wrappedViewController;
    }
    _wrappedViewController = [[UIViewController alloc] initWithNibName:nil bundle:nil];
    _wrappedViewController.view = self.view;
    MatchaConfigureChildViewController(_wrappedViewController);
    return _wrappedViewController;
}

- (UIView *)materializedView {
    return self.viewController.view ?: self.view;
}

@end
