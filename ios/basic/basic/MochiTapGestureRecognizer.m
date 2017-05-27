//
//  MochiTapGestureRecognizer.m
//  basic
//
//  Created by Kevin Dang on 5/26/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiTapGestureRecognizer.h"

@implementation MochiTapGestureRecognizer

- (id)initWithProtobuf:(GPBAny *)pb {
    NSError *error = nil;
    MochiPBTapRecognizer *pbTapRecognizer = (id)[pb unpackMessageClass:[MochiPBTapRecognizer class] error:&error];
    if (pbTapRecognizer == nil) {
        return nil;
    }
    if ((self = [super initWithTarget:self action:@selector(action:)])) {
        self.numberOfTapsRequired = pbTapRecognizer.count;
    }
    return self;
}

- (void)action:(id)sender {
    NSLog(@"ACTION,%@", self);
}

@end
