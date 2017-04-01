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
        
    }
}

@end
