//
//  MochiView.m
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiView.h"

@interface MochiView ()
@end

@implementation MochiView

- (void)setNode:(MochiNode *)value {
    if (_node != value) {
        _node = value;
        self.backgroundColor = _node.paintOptions.backgroundColor;
        self.frame = _node.guide.frame;
        self.layer.zPosition = _node.guide.zIndex;
        for (MochiNode *i in _node.nodeChildren.objectEnumerator) {
            NSString *name = i.bridgeName;
            MochiView *child = nil;
            if ([name isEqual:@""]) {
                child = [[MochiView alloc] init];
            } else if ([name isEqual:@"github.com/overcyn/mochi TextView"]) {
                child = [[MochiTextView alloc] init];
            } else if ([name isEqual:@"github.com/overcyn/mochi ImageView"]) {
                child = [[MochiImageView alloc] init];
            }
            child.node = i;
            [self addSubview:child];
        }
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
    if (!state.isNil) {
        NSData *data = state.elem.toData;
        UIImage *image = [[UIImage alloc] initWithData:data];
        self.imageView.image = image;
    } else {
        self.imageView.image = nil;
    }
}

- (void)layoutSubviews {
    self.imageView.frame = self.bounds;
}

@end
