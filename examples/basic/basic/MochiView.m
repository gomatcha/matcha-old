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

- (void)setNode:(BridgeValue *)value {
    if (_node != value) {
        _node = value;
        // self.backgroundColor = [[UIColor alloc] initWithBridgeValue:value];
    }
}

@end
