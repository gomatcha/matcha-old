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
        for (MochiNode *i in _node.nodeChildren.allValues) {
            MochiView *child = [[MochiView alloc] init];
            [self addSubview:child];
            child.node = i;
        }
    }
}

@end
