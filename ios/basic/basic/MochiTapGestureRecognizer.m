//
//  MochiTapGestureRecognizer.m
//  basic
//
//  Created by Kevin Dang on 5/26/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiTapGestureRecognizer.h"
#import <Mochi/mochigo.h>

@implementation MochiTapGestureRecognizer

- (id)initWithViewId:(int64_t)viewId recognizerId:(int64_t)recognizerId protobuf:(GPBAny *)pb {
    NSError *error = nil;
    MochiPBTapRecognizer *pbTapRecognizer = (id)[pb unpackMessageClass:[MochiPBTapRecognizer class] error:&error];
    if (pbTapRecognizer == nil) {
        return nil;
    }
    if ((self = [super initWithTarget:self action:@selector(action:)])) {
        self.numberOfTapsRequired = pbTapRecognizer.count;
        self.viewId = viewId;
        self.recognizerId = recognizerId;
    }
    return self;
}

- (void)action:(id)sender {
    MochiGoValue *viewId = [[MochiGoValue alloc] initWithLongLong:self.viewId];
    MochiGoValue *recognizerId = [[MochiGoValue alloc] initWithLongLong:self.recognizerId];
    [[[MochiGoValue alloc] initWithFunc:@"github.com/overcyn/mochi/touch TapRecognizer.Recognized"] call:nil args:@[viewId, recognizerId]];
}

@end
