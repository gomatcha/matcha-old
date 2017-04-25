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

@implementation MochiView

- (void)setNode:(MochiNode *)value {
    if (_node != value) {
        for (UIView *i in self.childViews) {
            [i removeFromSuperview];
        }
        
        _node = value;
        self.backgroundColor = _node.paintOptions.backgroundColor;
        self.frame = _node.guide.frame;
        
        NSMutableArray *array = [NSMutableArray array];
        for (MochiNode *i in _node.nodeChildren.objectEnumerator) {
            NSString *name = i.bridgeName;
            MochiView *child = nil;
            if ([name isEqual:@""]) {
                child = [[MochiView alloc] init];
            } else if ([name isEqual:@"github.com/overcyn/mochi TextView"]) {
                child = [[MochiTextView alloc] init];
            } else if ([name isEqual:@"github.com/overcyn/mochi ImageView"]) {
                child = [[MochiImageView alloc] init];
            } else if ([name isEqual:@"github.com/overcyn/mochi/view/button Button"]) {
                child = [[MochiButton alloc] init];
            }
            child.node = i;
            [array addObject:child];
        }
        [array sortUsingComparator:^NSComparisonResult(MochiView *obj1, MochiView *obj2) {
            return obj1.node.guide.zIndex > obj2.node.guide.zIndex;
        }];
        for (UIView *i in array) {
            [self addSubview:i];
        }
        self.childViews = array;
    }
}

@end

@interface MochiTextView ()
@property (nonatomic, strong) UILabel *label;
@end

@implementation MochiTextView

- (id)initWithFrame:(CGRect)frame {
    if ((self = [super initWithFrame:frame])) {
        self.label = [[UILabel alloc] init];
        [self addSubview:self.label];
    }
    return self;
}

- (void)setNode:(MochiNode *)node {
    [super setNode:node];
    MochiGoValue *state = node.bridgeState;
    MochiGoValue *formattedText = state[@"FormattedText"];
    if (!formattedText.isNil) {
        NSAttributedString *attrString = [[NSAttributedString alloc] initWithGoValue:formattedText];
        self.label.attributedText = attrString;
    }
}

- (void)layoutSubviews {
    self.label.frame = self.bounds;
}

@end

@interface MochiImageView ()
@property (nonatomic, strong) UIImageView *imageView;
@end

@implementation MochiImageView

- (id)initWithFrame:(CGRect)frame {
    if ((self = [super initWithFrame:frame])) {
        self.imageView = [[UIImageView alloc] init];
        [self addSubview:self.imageView];
    }
    return self;
}

- (void)setNode:(MochiNode *)node {
    [super setNode:node];
    MochiGoValue *state = node.bridgeState;
    MochiGoValue *imageData = state[@"Bytes"];
    if (!imageData.isNil) {
        NSData *data = imageData.toData;
        UIImage *image = [[UIImage alloc] initWithData:data];
        self.imageView.image = image;
    } else {
        self.imageView.image = nil;
    }
    
    switch (state[@"ResizeMode"].toLongLong) {
    case 1:
        self.imageView.contentMode = UIViewContentModeScaleAspectFit;
    case 2:
        self.imageView.contentMode = UIViewContentModeScaleAspectFill;
    case 3:
        self.imageView.contentMode = UIViewContentModeScaleToFill;
    case 4:
        self.imageView.contentMode = UIViewContentModeCenter;
    }
}

- (void)layoutSubviews {
    self.imageView.frame = self.bounds;
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
    MochiGoValue *state = node.bridgeState[@"FormattedText"];
    NSAttributedString *string = [[NSAttributedString alloc] initWithGoValue:state];
    // [self.button setAttributedTitle:string forState:UIControlStateNormal]; 
    [self.button setTitle:string.string forState:UIControlStateNormal]; 
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
