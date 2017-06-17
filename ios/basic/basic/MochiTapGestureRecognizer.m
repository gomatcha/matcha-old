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
#import "MochiProtobuf.h"
#import "MochiBridge.h"
#import "MochiViewController.h"

@interface MochiTapGestureRecognizer ()
@property (nonatomic, assign) int64_t funcId;
@property (nonatomic, assign) int64_t viewId;
@property (nonatomic, weak) MochiViewController *viewController;
@property (nonatomic, assign) bool disabled;
@end

@implementation MochiTapGestureRecognizer

- (id)initWithMochiVC:(MochiViewController *)viewController viewId:(int64_t)viewId protobuf:(GPBAny *)pb {
    NSError *error = nil;
    MochiPBTouchTapRecognizer *pbTapRecognizer = (id)[pb unpackMessageClass:[MochiPBTouchTapRecognizer class] error:&error];
    if (pbTapRecognizer == nil) {
        return nil;
    }
    if ((self = [super initWithTarget:self action:@selector(action:)])) {
        self.numberOfTapsRequired = pbTapRecognizer.count;
        self.viewController = viewController;
        self.funcId = pbTapRecognizer.recognizedFunc;
        self.viewId = viewId;
    }
    return self;
}

- (void)disable {
    self.disabled = true;
}

- (void)updateWithProtobuf:(GPBAny *)pb {
    NSError *error = nil;
    MochiPBTouchTapRecognizer *pbTapRecognizer = (id)[pb unpackMessageClass:[MochiPBTouchTapRecognizer class] error:&error];
    if (pbTapRecognizer == nil) {
        return;
    }
    self.funcId = pbTapRecognizer.recognizedFunc;
}

- (void)action:(id)sender {
    if (self.disabled) {
        return;
    }
    
    CGPoint point = [self locationInView:self.view];
    
    MochiPBTouchTapEvent *event = [[MochiPBTouchTapEvent alloc] init];
    event.position = [[MochiPBPoint alloc] initWithCGPoint:point];
    event.timestamp = [[GPBTimestamp alloc] initWithDate:[NSDate date]];
    
    NSData *data = [event data];
    MochiGoValue *value = [[MochiGoValue alloc] initWithData:data];
    
    [self.viewController call:self.funcId viewId:self.viewId args:@[value]];
}

@end
