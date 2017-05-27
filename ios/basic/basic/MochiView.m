//
//  MochiView.m
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiView.h"
#import "View.pbobjc.h"
#import "Layout.pbobjc.h"
#import "Text.pbobjc.h"
#import "Scrollview.pbobjc.h"
#import "Imageview.pbobjc.h"
#import "Button.pbobjc.h"

bool MochiConfigureViewWithNode(UIView *view, MochiNode *node, MochiViewConfig *config);
MochiView *MochiViewWithNode(MochiNode *node);

@interface MochiViewConfig : NSObject
@property (nonatomic, strong) NSDictionary<NSNumber *, UIView *> *childViews;
@property (nonatomic, strong) MochiNode *node;
@property (nonatomic, strong) NSNumber *buildId;
@property (nonatomic, strong) NSNumber *layoutId;
@property (nonatomic, strong) NSNumber *paintId;
@end

@implementation MochiViewConfig
@end

@interface MochiView ()
@property (nonatomic, strong) MochiViewConfig *config;
@end

@implementation MochiView

- (id)initWithFrame:(CGRect)frame {
    if ((self = [super initWithFrame:frame])) {
        self.config = [[MochiViewConfig alloc] init];
    }
    return self;
}

- (void)setNode:(MochiNode *)value {
    MochiConfigureViewWithNode(self, value, self.config);
}

@end

@interface MochiTextView ()
@property (nonatomic, strong) MochiViewConfig *config;
@end

@implementation MochiTextView

- (id)initWithFrame:(CGRect)frame {
    if ((self = [super initWithFrame:frame])) {
        self.config = [[MochiViewConfig alloc] init];
    }
    return self;
}

- (void)setNode:(MochiNode *)value {
    bool update = MochiConfigureViewWithNode(self, value, self.config);
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
@property (nonatomic, strong) MochiViewConfig *config;
@end

@implementation MochiImageView

- (id)initWithFrame:(CGRect)frame {
    if ((self = [super initWithFrame:frame])) {
        self.config = [[MochiViewConfig alloc] init];
    }
    return self;
}

- (void)setNode:(MochiNode *)value {
    bool update = MochiConfigureViewWithNode(self, value, self.config);
    if (update) {
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
}

@end

@interface MochiButton ()
@property (nonatomic, strong) UIButton *button;
@property (nonatomic, strong) MochiViewConfig *config;
@end

@implementation MochiButton

- (id)initWithFrame:(CGRect)frame {
    if ((self = [super initWithFrame:frame])) {
        self.config = [[MochiViewConfig alloc] init];
        self.button = [UIButton buttonWithType:UIButtonTypeSystem];
        [self.button addTarget:self action:@selector(onPress) forControlEvents:UIControlEventTouchUpInside];
        [self addSubview:self.button];
    }
    return self;
}

- (void)setNode:(MochiNode *)value {
    bool update = MochiConfigureViewWithNode(self, value, self.config);
    if (update) {
        GPBAny *state = value.nativeViewState;
        NSError *error = nil;
        MochiPBButton *pbbutton = (id)[state unpackMessageClass:[MochiPBButton class] error:&error];
        
        NSAttributedString *string = [[NSAttributedString alloc] initWithProtobuf:pbbutton.text];
        [self.button setAttributedTitle:string forState:UIControlStateNormal];
    }
}

- (void)layoutSubviews {
    self.button.frame = self.bounds;
}

- (void)onPress {
    MochiGoValue *identifier = [[MochiGoValue alloc] initWithLongLong:self.config.node.identifier.longLongValue];
    [[[MochiGoValue alloc] initWithFunc:@"github.com/overcyn/mochi/view/button OnPress"] call:nil args:@[identifier]];
}

@end

@interface MochiScrollView ()
@property (nonatomic, strong) MochiViewConfig *config;
@end

@implementation MochiScrollView

- (id)initWithFrame:(CGRect)frame {
    if ((self = [super initWithFrame:frame])) {
        self.config = [[MochiViewConfig alloc] init];
    }
    return self;
}

- (void)setNode:(MochiNode *)value {
    bool update = MochiConfigureViewWithNode(self, value, self.config);
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

bool MochiConfigureViewWithNode(UIView *view, MochiNode *node, MochiViewConfig *config) {
    if (![node.buildId isEqual:config.node.buildId]) {
        // Rebuild children
        NSMutableArray *addedKeys = [NSMutableArray array];
        NSMutableArray *removedKeys = [NSMutableArray array];
        NSMutableArray *unmodifiedKeys = [NSMutableArray array];
        
        for (NSNumber *i in config.node.nodeChildren) {
            MochiNode *child = node.nodeChildren[i];
            if (child == nil) {
                [removedKeys addObject:i];
            }
        }
        for (NSNumber *i in node.nodeChildren) {
            MochiNode *prevChild = config.node.nodeChildren[i];
            if (prevChild == nil) {
                [addedKeys addObject:i];
            } else {
                [unmodifiedKeys addObject:i];
            }
        }
        
        NSMutableDictionary *childViews = [NSMutableDictionary dictionary];
        for (NSNumber *i in removedKeys) {
            [config.childViews[i] removeFromSuperview];
        }
        for (NSNumber *i in addedKeys) {
            MochiView *childView = MochiViewWithNode(node.nodeChildren[i]);
            [view addSubview:childView];
            childViews[i] = childView;
        }
        for (NSNumber *i in unmodifiedKeys) {
            MochiView *childView = (id)config.childViews[i];
            childViews[i] = childView;
        }
        config.childViews = childViews;
    }
    if (![node.layoutId isEqual:config.node.layoutId]) {
        NSArray *sortedKeys = [[config.childViews allKeys] sortedArrayUsingComparator:^NSComparisonResult(NSNumber *obj1, NSNumber *obj2) {
            return node.nodeChildren[obj1].guide.zIndex > node.nodeChildren[obj2].guide.zIndex;
        }];
        NSArray *subviews = view.subviews;
        for (NSInteger i = 0; i < sortedKeys.count; i++) {
            NSNumber *key = sortedKeys[i];
            UIView *subview = config.childViews[key];
            if ([subviews indexOfObject:subview] != i) {
                [view insertSubview:subview atIndex:i];
            }
        }
        view.frame = node.guide.frame;
    }
    if (![node.paintId isEqual:config.node.paintId]) {
        view.backgroundColor = node.paintOptions.backgroundColor ?: [UIColor clearColor];
    }
    
    for (NSNumber *i in config.childViews) {
        ((MochiView *)config.childViews[i]).node = node.nodeChildren[i];
    }
    
    bool update = ![node.buildId isEqual:config.node.buildId];
    config.node = node;
    return update;
}

MochiView *MochiViewWithNode(MochiNode *node) {
    NSString *name = node.nativeViewName;
    MochiView *child = nil;
    if ([name isEqual:@""]) {
        child = [[MochiView alloc] initWithFrame:CGRectZero];
    } else if ([name isEqual:@"github.com/overcyn/mochi/view/textview"]) {
        child = (id)[[MochiTextView alloc] initWithFrame:CGRectZero];
    } else if ([name isEqual:@"github.com/overcyn/mochi/view/imageview"]) {
        child = (id)[[MochiImageView alloc] initWithFrame:CGRectZero];
    } else if ([name isEqual:@"github.com/overcyn/mochi/view/button"]) {
        child = (id)[[MochiButton alloc] initWithFrame:CGRectZero];
    } else if ([name isEqual:@"github.com/overcyn/mochi/view/scrollview"]) {
        child = (id)[[MochiScrollView alloc] initWithFrame:CGRectZero];
    }
    return child;
}
