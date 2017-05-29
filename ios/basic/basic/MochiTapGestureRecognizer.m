//
//  MochiTapGestureRecognizer.m
//  basic
//
//  Created by Kevin Dang on 5/26/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiTapGestureRecognizer.h"
#import <Mochi/mochigo.h>
#import "MochiNode.h"

@implementation MochiTapGestureRecognizer

- (id)initWitViewRoot:(MochiViewRoot *)viewRoot viewId:(int64_t)viewId protobuf:(GPBAny *)pb {
    NSError *error = nil;
    MochiPBTapRecognizer *pbTapRecognizer = (id)[pb unpackMessageClass:[MochiPBTapRecognizer class] error:&error];
    if (pbTapRecognizer == nil) {
        return nil;
    }
    if ((self = [super initWithTarget:self action:@selector(action:)])) {
        self.numberOfTapsRequired = pbTapRecognizer.count;
        self.viewRoot = viewRoot;
        self.funcId = pbTapRecognizer.recognizedFunc;
        self.viewId = viewId;
    }
    return self;
}

- (void)action:(id)sender {
    NSLog(@"blah,%@,%@", @(self.funcId), self.viewRoot);
    [self.viewRoot call:self.funcId viewId:self.viewId args:nil];
}

@end
