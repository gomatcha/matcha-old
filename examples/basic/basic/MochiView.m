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
@property (nonatomic, strong) NSMapTable *childViewsTable;
@property (nonatomic, strong) NSArray *childViews;
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
            self.contentSize = ((UIView *)self.config.childViews[0]).frame.size;
        }

        MochiGoValue *state = value.bridgeState;
        self.scrollEnabled = state[@"ScrollEnabled"].toBool;
        self.showsVerticalScrollIndicator = state[@"ShowsVerticalScrollIndicator"].toBool;
        self.showsHorizontalScrollIndicator = state[@"ShowsHorizontalScrollIndicator"].toBool;
    }
}
@end

bool MochiConfigureViewWithNode(UIView *view, MochiNode *node, MochiViewConfig *config) {
    bool update = node.updateId != config.node.updateId;
    update = true;
    
    // Update layout and paint options
    view.backgroundColor = node.paintOptions.backgroundColor ?: [UIColor clearColor];
    view.frame = node.guide.frame;
    
    // Rebuild children
    if (update) {
        NSMutableArray *addedKeys = [NSMutableArray array];
        NSMutableArray *removedKeys = [NSMutableArray array];
        NSMutableArray *rebuiltKeys = [NSMutableArray array];
        NSMutableArray *rebuiltKeys2 = [NSMutableArray array];
        NSMutableArray *unmodifiedKeys = [NSMutableArray array];
        NSMutableArray *unmodifiedKeys2 = [NSMutableArray array];
        
        for (MochiGoValue *i in config.node.nodeChildren.keyEnumerator) {
            MochiNode *prevChild = config.node.nodeChildren[i];
            MochiNode *child = nil;
            MochiGoValue *key = nil;
            for (MochiGoValue *j in node.nodeChildren.keyEnumerator) {
                if ([i isEqual:j]) {
                    child = node.nodeChildren[j];
                    key = j;
                    break;
                }
            }
            if (child.guide == nil) { // Ignore nodes without a guide
                continue;
            }
            
            if (child == nil) {
                [removedKeys addObject:i];
            } else if (child.buildId != prevChild.buildId) {
                [rebuiltKeys addObject:i];
                [rebuiltKeys2 addObject:key];
            }
        }
        for (MochiGoValue *i in node.nodeChildren.keyEnumerator) {
            MochiGoValue *prevKey = nil;
            MochiNode *prevChild = nil;
            MochiNode *child = node.nodeChildren[i];
            for (MochiGoValue *j in config.node.nodeChildren.keyEnumerator) {
                if ([i isEqual:j]) {
                    prevChild = config.node.nodeChildren[j];
                    prevKey = j;
                    break;
                }
            }
            if (child.guide == nil) { // Ignore nodes without a guide
                continue;
            }
           
            if (prevChild == nil) {
                [addedKeys addObject:i];
            } else if (child.buildId == prevChild.buildId) {
                [unmodifiedKeys addObject:i];
                [unmodifiedKeys2 addObject:prevKey];
            }
        }
        
        NSMapTable *childViewsTable = [NSMapTable strongToStrongObjectsMapTable];
        for (MochiGoValue *i in removedKeys) {
            [config.childViewsTable[i] removeFromSuperview];
        }
        for (NSInteger i = 0; i < rebuiltKeys.count; i++) {
            MochiGoValue *prevKey = rebuiltKeys[i];
            MochiGoValue *key = rebuiltKeys2[i];
            
            [config.childViewsTable[prevKey] removeFromSuperview];
            MochiView *childView = MochiViewWithNode(node.nodeChildren[key]);
            childViewsTable[key] = childView;
        }
        for (MochiGoValue *i in addedKeys) {
            MochiView *childView = MochiViewWithNode(node.nodeChildren[i]);
            childViewsTable[i] = childView;
        }
        for (NSInteger i = 0; i < unmodifiedKeys.count; i++) {
            MochiGoValue *prevKey = unmodifiedKeys2[i];
            MochiGoValue *key = unmodifiedKeys[i];
            
            MochiView *childView = config.childViewsTable[prevKey];
            childView.node = node.nodeChildren[key];
            childViewsTable[key] = childView;
            [childView removeFromSuperview];
        }
        
        NSMutableArray *childViews = [NSMutableArray array];
        for (UIView *i in childViewsTable.objectEnumerator) {
            [childViews addObject:i];
        }
        [childViews sortUsingComparator:^NSComparisonResult(MochiView *obj1, MochiView *obj2) {
            return obj1.config.node.guide.zIndex > obj2.config.node.guide.zIndex;
        }];
        for (UIView *i in childViews) {
            [view addSubview:i];
        }
        
        config.childViews = childViews;
        config.childViewsTable = childViewsTable;
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
