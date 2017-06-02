#import "MochiView.h"
#import "MochiProtobuf.h"
#import "MochiTapGestureRecognizer.h"
#import "MochiPressGestureRecognizer.h"
#import "MochiTabBarController.h"

@interface MochiBasicView ()
@property (nonatomic, weak) MochiViewNode *viewNode;
@property (nonatomic, strong) MochiNode *node;
@end

@implementation MochiBasicView

- (id)initWithViewNode:(MochiViewNode *)viewNode {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewNode = viewNode;
    }
    return self;
}

- (void)setNode:(MochiNode *)value {
    _node = value;
}

@end

@interface MochiTextView ()
@property (nonatomic, weak) MochiViewNode *viewNode;
@property (nonatomic, strong) MochiNode *node;
@end

@implementation MochiTextView

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
    MochiPBText *text = (id)[state unpackMessageClass:[MochiPBText class] error:&error];
    if (text != nil) {
        NSAttributedString *attrString = [[NSAttributedString alloc] initWithProtobuf:text];
        self.attributedText = attrString;
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
@property (nonatomic, weak) MochiViewNode *viewNode;
@property (nonatomic, strong) MochiNode *node;
@end

@implementation MochiScrollView

- (id)initWithViewNode:(MochiViewNode *)viewNode {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewNode = viewNode;
    }
    return self;
}

- (void)setNode:(MochiNode *)value {
    _node = value;

    if (self.subviews.count > 0) {
        self.contentSize = ((UIView *)self.subviews[0]).frame.size;
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
@end

UIGestureRecognizer *MochiGestureRecognizerWithPB(int64_t viewId, GPBAny *any, MochiViewNode *viewNode) {
    if ([any.typeURL isEqual:@"type.googleapis.com/mochi.touch.TapRecognizer"]) {
        return [[MochiTapGestureRecognizer alloc] initWithMochiVC:viewNode.rootVC viewId:viewId protobuf:any];
    } else if ([any.typeURL isEqual:@"type.googleapis.com/mochi.touch.PressRecognizer"]) {
        return [[MochiPressGestureRecognizer alloc] initWithMochiVC:viewNode.rootVC viewId:viewId protobuf:any];
    }
    return nil;
}

UIView<MochiChildView> *MochiViewWithNode(MochiNode *node, MochiViewNode *viewNode) {
    NSString *name = node.nativeViewName;
    UIView<MochiChildView> *child = nil;
    if ([name isEqual:@""]) {
        child = [[MochiBasicView alloc] initWithViewNode:viewNode];
    } else if ([name isEqual:@"github.com/overcyn/mochi/view/textview"]) {
        child = [[MochiTextView alloc] initWithViewNode:viewNode];
    } else if ([name isEqual:@"github.com/overcyn/mochi/view/imageview"]) {
        child = [[MochiImageView alloc] initWithViewNode:viewNode];
    } else if ([name isEqual:@"github.com/overcyn/mochi/view/button"]) {
        child = [[MochiButton alloc] initWithViewNode:viewNode];
    } else if ([name isEqual:@"github.com/overcyn/mochi/view/scrollview"]) {
        child = [[MochiScrollView alloc] initWithViewNode:viewNode];
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
        self.view = MochiViewWithNode(node, self);
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
    
    // Update gesture recognizers
    {
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
            UIGestureRecognizer *recognizer = MochiGestureRecognizerWithPB(node.identifier.longLongValue, node.touchRecognizers[i], self);
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
