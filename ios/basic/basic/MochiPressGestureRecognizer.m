//
//  MochiPressGestureRecognizer.m
//  basic
//
//  Created by Kevin Dang on 5/30/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiPressGestureRecognizer.h"
#import "MochiProtobuf.h"
#import "MochiNode.h"
#import "MochiViewController.h"

@interface MochiPressGestureRecognizer ()
@property (nonatomic, assign) int64_t funcId;
@property (nonatomic, assign) int64_t viewId;
@property (nonatomic, weak) MochiViewController *viewController;
@property (nonatomic, strong) NSDate *startTime;
@property (nonatomic, assign) BOOL disabled;
@end

@implementation MochiPressGestureRecognizer

- (id)initWithMochiVC:(MochiViewController *)viewController viewId:(int64_t)viewId protobuf:(GPBAny *)pb {
    NSError *error = nil;
    MochiPBPressRecognizer *pbTapRecognizer = (id)[pb unpackMessageClass:[MochiPBPressRecognizer class] error:&error];
    if (pbTapRecognizer == nil) {
        return nil;
    }
    if ((self = [super initWithTarget:self action:@selector(action:)])) {
        self.minimumPressDuration = pbTapRecognizer.minDuration.timeInterval;
        self.viewController = viewController;
        self.funcId = pbTapRecognizer.funcId;
        self.viewId = viewId;
    }
    return self;
}

- (void)updateWithProtobuf:(GPBAny *)pb {
    NSError *error = nil;
    MochiPBPressRecognizer *pbTapRecognizer = (id)[pb unpackMessageClass:[MochiPBPressRecognizer class] error:&error];
    if (pbTapRecognizer == nil) {
        return;
    }
    self.funcId = pbTapRecognizer.funcId;
}

- (void)disable {
    self.disabled = false;
}

- (void)action:(id)sender {
    if (self.disabled) {
        return;
    }
    
    CGPoint point = [self locationInView:self.view];
    
    MochiPBPressEvent *event = [[MochiPBPressEvent alloc] init];
    event.position = [[MochiPBPoint alloc] initWithCGPoint:point];
    event.timestamp = [[GPBTimestamp alloc] initWithDate:[NSDate date]];
    if (self.state == UIGestureRecognizerStateBegan) {
        event.kind = MochiPBEventKind_EventKindChanged;
        self.startTime = [NSDate date];
    } else if (self.state == UIGestureRecognizerStateChanged) {
        event.kind = MochiPBEventKind_EventKindChanged;
    } else if (self.state == UIGestureRecognizerStateEnded) {
        event.kind = MochiPBEventKind_EventKindRecognized;
    } else if (self.state == UIGestureRecognizerStateCancelled) {
        event.kind = MochiPBEventKind_EventKindFailed;
    } else {
        return;
    }
    event.duration = [[GPBDuration alloc] initWithTimeInterval:-self.startTime.timeIntervalSinceNow];
    
    NSData *data = [event data];
    MochiGoValue *value = [[MochiGoValue alloc] initWithData:data];
    
    [self.viewController call:self.funcId viewId:self.viewId args:@[value]];
}

@end
