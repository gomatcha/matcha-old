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
        NSLog(@"children:%@", _node.nodeChildren);
        for (MochiNode *i in _node.nodeChildren.objectEnumerator) {
            NSLog(@"Child %@",i);
            // NSString *name = i.bridgeName;
            MochiView *child = nil;
            // if ([name isEqual:@""]) {
                child = [[MochiView alloc] init];
            // } else if ([name isEqual:@"github.com/overcyn/mochi TextView"]) {
            //     child = [[MochiTextView alloc] init];
            // }
            child.node = i;
            [self addSubview:child];
        }
    }
}

@end

// @interface MochiTextView ()
// @property (nonatomic, strong) UILabel *label;
// @end

// @implementation MochiTextView

// - (id)initWithFrame:(CGRect)frame {
//     if ((self = [super initWithFrame:frame])) {
//         self.label = [[UILabel alloc] init];
//         [self addSubview:self.label];
//     }
//     return self;
// }

// - (void)setNode:(MochiNode *)node {
//     [super setNode:node];
//     BridgeValue *state = node.bridgeState;
//     // BridgeValue *text = state[@"Text"];
//     // BridgeValue *format = state[@"Format"];
//     BridgeValue *formattedText = state[@"FormattedText"];
//     if (!formattedText.isNil) {
//         NSAttributedString *attrString = [[NSAttributedString alloc] initWithBridgeValue:formattedText];
//         self.label.attributedText = attrString;
//     }
// }

// - (void)layoutSubviews {
//     self.label.frame = self.bounds;
// }

// @end


