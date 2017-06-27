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

@interface MatchaBasicView ()
@property (nonatomic, weak) MatchaViewNode *viewNode;
@property (nonatomic, strong) MatchaNode *node;
@end

@implementation MatchaBasicView

- (id)initWithViewNode:(MatchaViewNode *)viewNode {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewNode = viewNode;
    }
    return self;
}

- (void)setNode:(MatchaNode *)value {
    _node = value;
}

@end

@interface MatchaTextView ()
@property (nonatomic, weak) MatchaViewNode *viewNode;
@property (nonatomic, strong) MatchaNode *node;
@end

@implementation MatchaTextView

- (id)initWithViewNode:(MatchaViewNode *)viewNode {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewNode = viewNode;
    }
    return self;
}

- (void)setNode:(MatchaNode *)value {
    _node = value;
    GPBAny *state = value.nativeViewState;
    NSError *error = nil;
    MatchaPBStyledText *text = (id)[state unpackMessageClass:[MatchaPBStyledText class] error:&error];
    if (text != nil) {
        NSAttributedString *attrString = [[NSAttributedString alloc] initWithProtobuf:text];
        self.attributedText = attrString;
        self.numberOfLines = 0;
    }
}

@end

@interface MatchaImageView ()
@property (nonatomic, weak) MatchaViewNode *viewNode;
@property (nonatomic, strong) MatchaNode *node;
@end

@implementation MatchaImageView

- (id)initWithViewNode:(MatchaViewNode *)viewNode {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewNode = viewNode;
    }
    return self;
}

- (void)setNode:(MatchaNode *)value {
    _node = value;
    GPBAny *state = value.nativeViewState;
    NSError *error = nil;
    MatchaImageViewPBView *pbimageview = (id)[state unpackMessageClass:[MatchaImageViewPBView class] error:&error];
    
    UIImage *image = [[UIImage alloc] initWithImageOrResourceProtobuf:pbimageview.image];
    
    switch (pbimageview.resizeMode) {
    case MatchaImageViewPBResizeMode_Fit:
        self.contentMode = UIViewContentModeScaleAspectFit;
        break;
    case MatchaImageViewPBResizeMode_Fill:
        self.contentMode = UIViewContentModeScaleAspectFill;
        break;
    case MatchaImageViewPBResizeMode_Stretch:
        self.contentMode = UIViewContentModeScaleToFill;
        break;
    case MatchaImageViewPBResizeMode_Center:
        self.contentMode = UIViewContentModeCenter;
        break;
    }
    if (pbimageview.hasTint) {
        self.tintColor = [[UIColor alloc] initWithProtobuf:pbimageview.tint];
        image = [image imageWithRenderingMode:UIImageRenderingModeAlwaysTemplate];
    }
    
    if (![self.image isEqual:image]) {
        self.image = image;
    }
}

@end

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
    if ([name isEqual:@""]) {
        child = [[MatchaBasicView alloc] initWithViewNode:viewNode];
    } else if ([name isEqual:@"github.com/overcyn/matcha/view/textview"]) {
        child = [[MatchaTextView alloc] initWithViewNode:viewNode];
    } else if ([name isEqual:@"github.com/overcyn/matcha/view/imageview"]) {
        child = [[MatchaImageView alloc] initWithViewNode:viewNode];
    } else if ([name isEqual:@"github.com/overcyn/matcha/view/button"]) {
        child = [[MatchaButton alloc] initWithViewNode:viewNode];
    } else if ([name isEqual:@"github.com/overcyn/matcha/view/scrollview"]) {
        child = [[MatchaScrollView alloc] initWithViewNode:viewNode];
    } else if ([name isEqual:@"github.com/overcyn/matcha/view/switch"]) {
        child = [[MatchaSwitchView alloc] initWithViewNode:viewNode];
    } else if ([name isEqual:@"github.com/overcyn/matcha/view/textinput"]) {
        child = [[MatchaTextInput alloc] initWithViewNode:viewNode];
    } else if ([name isEqual:@"github.com/overcyn/matcha/view/slider"]) {
        child = [[MatchaSlider alloc] initWithViewNode:viewNode];
    }
    return child;
}

UIViewController<MatchaChildViewController> *MatchaViewControllerWithNode(MatchaNode *node, MatchaViewNode *viewNode) {
    NSString *name = node.nativeViewName;
    UIViewController<MatchaChildViewController> *child = nil;
    if ([name isEqual:@"github.com/overcyn/matcha/view/tabscreen"]) {
        child = [[MatchaTabScreen alloc] initWithViewNode:viewNode];
    } else if ([name isEqual:@"github.com/overcyn/matcha/view/stacknav"]) {
        child = [[MatchaStackScreen alloc] initWithViewNode:viewNode];
    } else if ([name isEqual:@"github.com/overcyn/matcha/view/stacknav Bar"]) {
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
            NSLog(@"Cannot find corresponding view or view controller for node:%@", node.nativeViewName);
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
        
        // let view controllers do their own layout
        if (self.parent.viewController == nil) {
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
