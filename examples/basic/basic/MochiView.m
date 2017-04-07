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
        for (MochiNode *i in _node.nodeChildren.objectEnumerator) {
            NSString *name = i.bridgeName;
            MochiView *child = nil;
            if ([name isEqual:@""]) {
                child = [[MochiView alloc] init];
            } else if ([name isEqual:@"github.com/overcyn/mochi TextView"]) {
                child = [[MochiTextView alloc] init];
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
    BridgeValue *state = node.bridgeState;
    BridgeValue *text = state[@"Text"];
    self.label.text = text.toString;
    self.label.textColor = [[UIColor alloc] initWithBridgeValue:state[@"TextColor"]];
}

- (void)layoutSubviews {
    self.label.frame = self.bounds;
}

@end
