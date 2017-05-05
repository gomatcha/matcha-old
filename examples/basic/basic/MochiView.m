//
//  MochiView.m
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiView.h"

bool MochiConfigureViewWithNode(UIView *view, MochiNode *node, MochiViewConfig *config);
MochiView *MochiViewWithNode(MochiNode *node);

@interface MochiViewConfig : NSObject
@property (nonatomic, strong) NSDictionary<NSNumber *, UIView *> *childViews;
@property (nonatomic, strong) MochiNode *node;
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
        MochiGoValue *state = value.bridgeState;
        MochiGoValue *formattedText = state[@"Text"];
        if (!formattedText.isNil) {
            NSAttributedString *attrString = [[NSAttributedString alloc] initWithGoValue:formattedText];
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
        MochiGoValue *state = value.bridgeState;
        MochiGoValue *imageData = state[@"Bytes"];
        if (!imageData.isNil) {
            NSData *data = imageData.toData;
            UIImage *image = [[UIImage alloc] initWithData:data];
            self.image = image;
        } else {
            self.image = nil;
        }
        
        switch (state[@"ResizeMode"].toLongLong) {
        case 1:
            self.contentMode = UIViewContentModeScaleAspectFit;
        case 2:
            self.contentMode = UIViewContentModeScaleAspectFill;
        case 3:
            self.contentMode = UIViewContentModeScaleToFill;
        case 4:
            self.contentMode = UIViewContentModeCenter;
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
        MochiGoValue *state = value.bridgeState[@"Text"];
        NSAttributedString *string = [[NSAttributedString alloc] initWithGoValue:state];
        [self.button setAttributedTitle:string forState:UIControlStateNormal];
    }
}

- (void)layoutSubviews {
    self.button.frame = self.bounds;
}

- (void)onPress {
    MochiGoValue *onPress = self.config.node.bridgeState[@"OnPress"];
    if (!onPress.isNil) {
        [onPress call:nil args:nil];
    }
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

        MochiGoValue *state = value.bridgeState;
        self.scrollEnabled = state[@"ScrollEnabled"].toBool;
        self.showsVerticalScrollIndicator = state[@"ShowsVerticalScrollIndicator"].toBool;
        self.showsHorizontalScrollIndicator = state[@"ShowsHorizontalScrollIndicator"].toBool;
    }
}
@end

bool MochiConfigureViewWithNode(UIView *view, MochiNode *node, MochiViewConfig *config) {
    // bool update = node.updateId != config.node.updateId;
    bool update = true;
    
    // Update layout and paint options
    view.backgroundColor = node.paintOptions.backgroundColor ?: [UIColor clearColor];
    view.frame = node.guide.frame;
    
    // Rebuild children
    if (update) {
        NSMutableArray *addedKeys = [NSMutableArray array];
        NSMutableArray *removedKeys = [NSMutableArray array];
        NSMutableArray *unmodifiedKeys = [NSMutableArray array];
        
        for (NSNumber *i in config.node.nodeChildren) {
            MochiNode *child = node.nodeChildren[i];
            if (child.guide == nil) { // Ignore nodes without a guide
                continue;
            }            
            if (child == nil) {
                [removedKeys addObject:i];
            }
        }
        for (NSNumber *i in node.nodeChildren) {
            MochiNode *prevChild = config.node.nodeChildren[i];
            MochiNode *child = node.nodeChildren[i];
            if (child.guide == nil) { // Ignore nodes without a guide
                continue;
            }
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
            childView.node = node.nodeChildren[i];
            childViews[i] = childView;
            [childView removeFromSuperview];
        }
        
        NSArray *sortedKeys = [[childViews allKeys] sortedArrayUsingComparator:^NSComparisonResult(NSNumber *obj1, NSNumber *obj2) {
            return node.nodeChildren[obj1].guide.zIndex > node.nodeChildren[obj2].guide.zIndex;
        }];
        NSArray *subviews = view.subviews;
        for (NSInteger i = 0; i < sortedKeys.count; i++) {
            NSNumber *key = sortedKeys[i];
            UIView *subview = childViews[key];
            if ([subviews indexOfObject:subview] != i) {
                [view insertSubview:subview atIndex:i];
            }
        }
        config.childViews = childViews;
    }
    
    config.node = node;
    return update;
}

MochiView *MochiViewWithNode(MochiNode *node) {
    NSString *name = node.bridgeName;
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
    child.node = node;
    return child;
}
