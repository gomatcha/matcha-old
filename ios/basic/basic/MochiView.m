#import "MochiView.h"
#import "MochiProtobuf.h"
#import "MochiTapGestureRecognizer.h"
#import "MochiPressGestureRecognizer.h"
#import "MochiTabBarController.h"

@interface MochiViewConfig : NSObject
@property (nonatomic, strong) NSDictionary<NSNumber *, UIView *> *childViews;
@property (nonatomic, strong) NSDictionary<NSNumber *, UIGestureRecognizer *> *touchRecognizers;
@property (nonatomic, strong) MochiNode *node;
@property (nonatomic, strong) NSNumber *buildId;
@property (nonatomic, strong) NSNumber *layoutId;
@property (nonatomic, strong) NSNumber *paintId;
@end

@implementation MochiViewConfig
@end

@interface MochiBasicView ()
@property (nonatomic, strong) MochiViewConfig *config;
@end

@implementation MochiBasicView

- (id)initWithViewRoot:(MochiViewController *)viewRoot parentVC:(UIViewController *)parentVC {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewRoot = viewRoot;
        self.parentVC = parentVC;
        self.config = [[MochiViewConfig alloc] init];
    }
    return self;
}

- (void)setNode:(MochiNode *)value {
    MochiConfigureViewWithNode(self, value, self.config, self.viewRoot);
}

@end

@interface MochiTextView ()
@property (nonatomic, strong) MochiViewConfig *config;
@end

@implementation MochiTextView

- (id)initWithViewRoot:(MochiViewController *)viewRoot parentVC:(UIViewController *)parentVC {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewRoot = viewRoot;
        self.parentVC = parentVC;
        self.config = [[MochiViewConfig alloc] init];
    }
    return self;
}

- (void)setNode:(MochiNode *)value {
    bool update = MochiConfigureViewWithNode(self, value, self.config, self.viewRoot);
    if (update) {
        GPBAny *state = value.nativeViewState;
        NSError *error = nil;
        MochiPBText *text = (id)[state unpackMessageClass:[MochiPBText class] error:&error];
        if (text != nil) {
            NSAttributedString *attrString = [[NSAttributedString alloc] initWithProtobuf:text];
            self.attributedText = attrString;
        }
    }
}

@end

@interface MochiImageView ()
@property (nonatomic, weak) MochiViewNode *viewNode;
@property (nonatomic, strong) MochiNode *node;
@end

@implementation MochiImageView

- (id)initWithViewNode:(MochiViewNode *)viewNode {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewNode = viewNode;
    }
    return self;
}

- (void)setNode:(MochiNode *)value {
    _node = value;
    GPBAny *state = value.nativeViewState;
    NSError *error = nil;
    MochiPBImageView *pbimageview = (id)[state unpackMessageClass:[MochiPBImageView class] error:&error];
    
    self.image = [[UIImage alloc] initWithProtobuf:pbimageview.image];
    switch (pbimageview.resizeMode) {
    case MochiPBResizeMode_Fit:
        self.contentMode = UIViewContentModeScaleAspectFit;
        break;
    case MochiPBResizeMode_Fill:
        self.contentMode = UIViewContentModeScaleAspectFill;
        break;
    case MochiPBResizeMode_Stretch:
        self.contentMode = UIViewContentModeScaleToFill;
        break;
    case MochiPBResizeMode_Center:
        self.contentMode = UIViewContentModeCenter;
        break;
    }
}

@end

@interface MochiButton ()
@property (nonatomic, strong) UIButton *button;
@property (nonatomic, weak) MochiViewNode *viewNode;
@property (nonatomic, strong) MochiNode *node;
@end

@implementation MochiButton

- (id)initWithViewNode:(MochiViewNode *)viewNode {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewNode = viewNode;
        self.button = [UIButton buttonWithType:UIButtonTypeSystem];
        [self.button addTarget:self action:@selector(onPress) forControlEvents:UIControlEventTouchUpInside];
        [self addSubview:self.button];
    }
    return self;
}

- (void)setNode:(MochiNode *)value {
    _node = value;
    GPBAny *state = value.nativeViewState;
    NSError *error = nil;
    MochiPBButtonButton *pbbutton = (id)[state unpackMessageClass:[MochiPBButtonButton class] error:&error];
    
    NSAttributedString *string = [[NSAttributedString alloc] initWithProtobuf:pbbutton.text];
    [self.button setAttributedTitle:string forState:UIControlStateNormal];
}

- (void)layoutSubviews {
    self.button.frame = self.bounds;
}

- (void)onPress {
    MochiGoValue *identifier = [[MochiGoValue alloc] initWithLongLong:self.node.identifier.longLongValue];
    [[[MochiGoValue alloc] initWithFunc:@"github.com/overcyn/mochi/view/button OnPress"] call:nil args:@[identifier]];
}

@end

@interface MochiScrollView ()
@property (nonatomic, strong) MochiViewConfig *config;
@end

@implementation MochiScrollView

- (id)initWithViewRoot:(MochiViewController *)viewRoot parentVC:(UIViewController *)parentVC {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewRoot = viewRoot;
        self.config = [[MochiViewConfig alloc] init];
        self.parentVC = parentVC;
    }
    return self;
}

- (void)setNode:(MochiNode *)value {
    bool update = MochiConfigureViewWithNode(self, value, self.config, self.viewRoot);
    if (update) {
        if (self.config.childViews.count > 0) {
            self.contentSize = ((UIView *)self.config.childViews.allValues[0]).frame.size;
        }

        GPBAny *state = value.nativeViewState;
        NSError *error = nil;
        MochiPBScrollView *pbscrollview = (id)[state unpackMessageClass:[MochiPBScrollView class] error:&error];
        if (pbscrollview != nil) {
            self.scrollEnabled = pbscrollview.scrollEnabled;
            self.showsVerticalScrollIndicator = pbscrollview.showsVerticalScrollIndicator;
            self.showsHorizontalScrollIndicator = pbscrollview.showsHorizontalScrollIndicator;
        }
    }
}
@end

bool MochiConfigureViewWithNode(UIView *view, MochiNode *node, MochiViewConfig *config, MochiViewController *viewRoot) {
    return TRUE;
    // if (![node.buildId isEqual:config.node.buildId]) {
    //     {
    //         // Rebuild children
    //         NSMutableArray *addedKeys = [NSMutableArray array];
    //         NSMutableArray *removedKeys = [NSMutableArray array];
    //         NSMutableArray *unmodifiedKeys = [NSMutableArray array];
            
    //         for (NSNumber *i in config.node.nodeChildren) {
    //             MochiNode *child = node.nodeChildren[i];
    //             if (child == nil) {
    //                 [removedKeys addObject:i];
    //             }
    //         }
    //         for (NSNumber *i in node.nodeChildren) {
    //             MochiNode *prevChild = config.node.nodeChildren[i];
    //             if (prevChild == nil) {
    //                 [addedKeys addObject:i];
    //             } else {
    //                 [unmodifiedKeys addObject:i];
    //             }
    //         }
            
    //         NSMutableDictionary *childViews = [NSMutableDictionary dictionary];
    //         for (NSNumber *i in removedKeys) {
    //             [config.childViews[i] removeFromSuperview];
    //         }
    //         for (NSNumber *i in addedKeys) {
    //             MochiBasicView *childView = MochiViewWithNode(node.nodeChildren[i], viewRoot, nil);
    //             [view addSubview:childView];
    //             childViews[i] = childView;
    //         }
    //         for (NSNumber *i in unmodifiedKeys) {
    //             MochiBasicView *childView = (id)config.childViews[i];
    //             childViews[i] = childView;
    //         }
    //         config.childViews = childViews;
    //     }
    
    //     // Update gesture recognizers
    //     {
    //         NSMutableArray *addedKeys = [NSMutableArray array];
    //         NSMutableArray *removedKeys = [NSMutableArray array];
    //         NSMutableArray *unmodifiedKeys = [NSMutableArray array];
            
    //         for (NSNumber *i in config.node.touchRecognizers) {
    //             GPBAny *child = node.touchRecognizers[i];
    //             if (child == nil) {
    //                 [removedKeys addObject:i];
    //             }
    //         }
    //         for (NSNumber *i in node.touchRecognizers) {
    //             GPBAny *prevChild = config.node.touchRecognizers[i];
    //             if (prevChild == nil) {
    //                 [addedKeys addObject:i];
    //             } else {
    //                 [unmodifiedKeys addObject:i];
    //             }
    //         }
            
    //         NSMutableDictionary *touchRecognizers = [NSMutableDictionary dictionary];
    //         for (NSNumber *i in removedKeys) {
    //             UIGestureRecognizer *recognizer = config.touchRecognizers[i];
    //             [(id)recognizer disable];
    //             [view removeGestureRecognizer:recognizer];
    //         }
    //         for (NSNumber *i in addedKeys) {
    //             UIGestureRecognizer *recognizer = MochiGestureRecognizerWithPB(node.identifier.longLongValue, node.touchRecognizers[i], viewRoot);
    //             [view addGestureRecognizer:recognizer];
    //             touchRecognizers[i] = recognizer;
    //         }
    //         for (NSNumber *i in unmodifiedKeys) {
    //             UIGestureRecognizer *recognizer = (id)config.touchRecognizers[i];
    //             [(id)recognizer updateWithProtobuf:node.touchRecognizers[i]];
    //             touchRecognizers[i] = recognizer;
    //         }
    //         config.touchRecognizers = touchRecognizers;
    //     }
    // }
    // if (![node.layoutId isEqual:config.node.layoutId]) {
    //     NSArray *sortedKeys = [[config.childViews allKeys] sortedArrayUsingComparator:^NSComparisonResult(NSNumber *obj1, NSNumber *obj2) {
    //         return node.nodeChildren[obj1].guide.zIndex > node.nodeChildren[obj2].guide.zIndex;
    //     }];
    //     NSArray *subviews = view.subviews;
    //     for (NSInteger i = 0; i < sortedKeys.count; i++) {
    //         NSNumber *key = sortedKeys[i];
    //         UIView *subview = config.childViews[key];
    //         if ([subviews indexOfObject:subview] != i) {
    //             [view insertSubview:subview atIndex:i];
    //         }
    //     }
    //     view.frame = node.guide.frame;
    // }
    // if (![node.paintId isEqual:config.node.paintId]) {
    //     MochiPaintOptions *paintOptions = node.paintOptions;
    //     view.alpha = 1 - paintOptions.transparency;
    //     view.backgroundColor = paintOptions.backgroundColor;
    //     view.layer.borderColor = paintOptions.borderColor.CGColor;
    //     view.layer.borderWidth = paintOptions.borderWidth;
    //     view.layer.cornerRadius = paintOptions.cornerRadius;
    //     view.layer.shadowRadius = paintOptions.shadowRadius;
    //     view.layer.shadowOffset = paintOptions.shadowOffset;
    //     view.layer.shadowColor = paintOptions.shadowColor.CGColor;
    //     view.layer.shadowOpacity = view.layer.shadowColor == nil ? 0 : 1;
    // }
    
    // for (NSNumber *i in config.childViews) {
    //     ((MochiBasicView *)config.childViews[i]).node = node.nodeChildren[i];
    // }
    
    // bool update = ![node.buildId isEqual:config.node.buildId];
    // config.node = node;
    // return update;
}

UIGestureRecognizer *MochiGestureRecognizerWithPB(int64_t viewId, GPBAny *any, MochiViewController *viewRoot) {
    if ([any.typeURL isEqual:@"type.googleapis.com/mochi.touch.TapRecognizer"]) {
        return [[MochiTapGestureRecognizer alloc] initWithMochiVC:viewRoot viewId:viewId protobuf:any];
    } else if ([any.typeURL isEqual:@"type.googleapis.com/mochi.touch.PressRecognizer"]) {
        return [[MochiPressGestureRecognizer alloc] initWithMochiVC:viewRoot viewId:viewId protobuf:any];
    }
    return nil;
}

UIView<MochiChildView> *MochiViewWithNode(MochiNode *node, MochiViewController *root, UIViewController *parentVC, MochiViewNode *viewNode) {
    NSString *name = node.nativeViewName;
    MochiBasicView *child = nil;
    if ([name isEqual:@""]) {
        child = [[MochiBasicView alloc] initWithViewRoot:root parentVC:parentVC];
    } else if ([name isEqual:@"github.com/overcyn/mochi/view/textview"]) {
        child = (id)[[MochiTextView alloc] initWithViewRoot:root parentVC:parentVC];
    } else if ([name isEqual:@"github.com/overcyn/mochi/view/imageview"]) {
        child = (id)[[MochiImageView alloc] initWithViewNode:viewNode];
    } else if ([name isEqual:@"github.com/overcyn/mochi/view/button"]) {
        child = (id)[[MochiButton alloc] initWithViewNode:viewNode];
    } else if ([name isEqual:@"github.com/overcyn/mochi/view/scrollview"]) {
        child = (id)[[MochiScrollView alloc] initWithViewRoot:root parentVC:parentVC];
    }
    return child;
}

UIViewController *MochiViewControllerWithNode(MochiNode *node, MochiViewController *root) {
    NSString *name = node.nativeViewName;
    UIViewController *child = nil;
    if ([name isEqual:@"github.com/overcyn/mochi/view/tabnavigator"]) {
        child = [[MochiTabBarController alloc] initWithViewRoot:root];
    }
    return child;
}

@implementation MochiViewNode

- (id)initWithParent:(MochiViewNode *)node rootVC:(MochiViewController *)rootVC {
    if ((self = [super init])) {
        self.parent = node;
        self.rootVC = rootVC; 
    }
    return self;
}

- (void)setNode:(MochiNode *)node {
    NSAssert(self.node == nil || [self.node.nativeViewName isEqual:node.nativeViewName], @"Node with different name");
    
    if (self.view == nil) {
        self.view = MochiViewWithNode(node, self.rootVC, nil, self);
    }
    
    // Build children
    NSDictionary<NSNumber *, MochiViewNode *> *children = nil;
    NSMutableArray *addedKeys = [NSMutableArray array];
    NSMutableArray *removedKeys = [NSMutableArray array];
    NSMutableArray *unmodifiedKeys = [NSMutableArray array];
    if (![node.buildId isEqual:self.node.buildId]) {
        for (NSNumber *i in self.children) {
            MochiNode *child = node.nodeChildren[i];
            if (child == nil) {
                [removedKeys addObject:i];
            }
        }
        for (NSNumber *i in node.nodeChildren) {
            MochiViewNode *prevChild = self.children[i];
            if (prevChild == nil) {
                [addedKeys addObject:i];
            } else {
                [unmodifiedKeys addObject:i];
            }
        }
        
        // Add/remove child nodes
        NSMutableDictionary<NSNumber *, MochiViewNode *> *mutChildren = [NSMutableDictionary dictionary];
        for (NSNumber *i in addedKeys) {
            mutChildren[i] = [[MochiViewNode alloc] initWithParent:self rootVC:self.rootVC];
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
        MochiViewNode *child = children[i];
        child.node = node.nodeChildren[i];
    }
    
    // Add/remove subviews
    if (![node.buildId isEqual:self.node.buildId]) {   
        for (NSNumber *i in addedKeys) {
            MochiViewNode *child = children[i];
            child.view.node = node.nodeChildren[i];
            [self.view addSubview:child.view];
        }
        for (NSNumber *i in removedKeys) {
            MochiViewNode *child = self.children[i];
            [child.view removeFromSuperview];
        }
        for (NSNumber *i in unmodifiedKeys) {
            MochiViewNode *child = children[i];
            child.view.node = node.nodeChildren[i];
        }
    }
    
    // Layout subviews
    if (![node.layoutId isEqual:self.node.layoutId]) {
        NSArray *sortedKeys = [[self.children allKeys] sortedArrayUsingComparator:^NSComparisonResult(NSNumber *obj1, NSNumber *obj2) {
            return node.nodeChildren[obj1].guide.zIndex > node.nodeChildren[obj2].guide.zIndex;
        }];
        
        for (NSInteger i = 0; i < sortedKeys.count; i++) {
            NSNumber *key = sortedKeys[i];
            UIView *subview = self.children[key].view;
            if ([self.view.subviews indexOfObject:subview] != i) {
                [self.view insertSubview:subview atIndex:i];
            }
        }
        self.view.frame = node.guide.frame;
    }
    
    // Paint view
    if (![node.paintId isEqual:self.node.paintId]) {
        MochiPaintOptions *paintOptions = node.paintOptions;
        self.view.alpha = 1 - paintOptions.transparency;
        self.view.backgroundColor = paintOptions.backgroundColor;
        self.view.layer.borderColor = paintOptions.borderColor.CGColor;
        self.view.layer.borderWidth = paintOptions.borderWidth;
        self.view.layer.cornerRadius = paintOptions.cornerRadius;
        self.view.layer.shadowRadius = paintOptions.shadowRadius;
        self.view.layer.shadowOffset = paintOptions.shadowOffset;
        self.view.layer.shadowColor = paintOptions.shadowColor.CGColor;
        self.view.layer.shadowOpacity = paintOptions.shadowColor == nil ? 0 : 1;
    }
    
    _node = node;
    self.children = children;
}

@end
