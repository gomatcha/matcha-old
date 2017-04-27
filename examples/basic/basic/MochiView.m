//
//  MochiView.m
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiView.h"

@interface MochiView ()
@property (nonatomic, strong) NSArray *childViews;
@end

NSArray *MochiConfigureViewWithNode(UIView *view,  MochiNode *node, NSArray *prevChildren);
NSArray *MochiConfigureViewWithNode(UIView *view,  MochiNode *node, NSArray *prevChildren) {
    for (UIView *i in prevChildren) {
        [i removeFromSuperview];
    }
    
    view.backgroundColor = node.paintOptions.backgroundColor;
    view.frame = node.guide.frame;
    
    NSMutableArray *array = [NSMutableArray array];
    for (MochiNode *i in node.nodeChildren.objectEnumerator) {
        NSString *name = i.bridgeName;
        MochiView *child = nil;
        if ([name isEqual:@""]) {
            child = [[MochiView alloc] init];
        } else if ([name isEqual:@"github.com/overcyn/mochi/view/textview"]) {
            child = (id)[[MochiTextView alloc] init];
        } else if ([name isEqual:@"github.com/overcyn/mochi/view/imageview"]) {
            child = (id)[[MochiImageView alloc] init];
        } else if ([name isEqual:@"github.com/overcyn/mochi/view/button"]) {
            child = [[MochiButton alloc] init];
        } else if ([name isEqual:@"github.com/overcyn/mochi/view/scrollview"]) {
            child = (id)[[MochiScrollView alloc] init];
        }
        child.node = i;
        [array addObject:child];
    }
    [array sortUsingComparator:^NSComparisonResult(MochiView *obj1, MochiView *obj2) {
        return obj1.node.guide.zIndex > obj2.node.guide.zIndex;
    }];
    for (UIView *i in array) {
        [view addSubview:i];
    }
    return array;
}

@implementation MochiView

- (void)setNode:(MochiNode *)value {
    if (_node != value) {
        _node = value;
        self.childViews = MochiConfigureViewWithNode(self, _node, self.childViews);
    }
}

@end

@interface MochiTextView ()
@property (nonatomic, strong) MochiNode *node;
@property (nonatomic, strong) NSArray *childViews;
@end

@implementation MochiTextView

- (void)setNode:(MochiNode *)value {
    if (_node != value) {
        _node = value;
        self.childViews = MochiConfigureViewWithNode(self, _node, self.childViews);
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
@property (nonatomic, strong) MochiNode *node;
@property (nonatomic, strong) NSArray *childViews;
@end

@implementation MochiImageView

- (void)setNode:(MochiNode *)value {
    if (_node != value) {
        _node = value;
        self.childViews = MochiConfigureViewWithNode(self, _node, self.childViews);
        MochiGoValue *state = _node.bridgeState;
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
@end

@implementation MochiButton

- (id)initWithFrame:(CGRect)frame {
    if ((self = [super initWithFrame:frame])) {
        self.button = [UIButton buttonWithType:UIButtonTypeSystem];
        [self.button addTarget:self action:@selector(onPress) forControlEvents:UIControlEventTouchUpInside];
        [self addSubview:self.button];
    }
    return self;
}

- (void)setNode:(MochiNode *)node {
    [super setNode:node];
    MochiGoValue *state = node.bridgeState[@"Text"];
    NSAttributedString *string = [[NSAttributedString alloc] initWithGoValue:state];
    [self.button setAttributedTitle:string forState:UIControlStateNormal];
}

- (void)layoutSubviews {
    self.button.frame = self.bounds;
}

- (void)onPress {
    MochiGoValue *onPress = self.node.bridgeState[@"OnPress"];
    if (!onPress.isNil) {
        [onPress call:nil args:nil];
    }
}

@end

@interface MochiScrollView ()
@property (nonatomic, strong) MochiNode *node;
@property (nonatomic, strong) NSArray *childViews;
@end

@implementation MochiScrollView

- (void)setNode:(MochiNode *)value {
    if (_node != value) {
        _node = value;
        self.childViews = MochiConfigureViewWithNode(self, _node, self.childViews);
        
        if (self.childViews.count > 0) {
            self.contentSize = ((UIView *)self.childViews[0]).frame.size;
        }
        MochiGoValue *state = _node.bridgeState;
//        MochiGoValue *imageData = state[@"Bytes"];
    }
}
@end
