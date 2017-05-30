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

@interface MochiPressGestureRecognizer ()
@property (nonatomic, assign) int64_t funcId;
@property (nonatomic, assign) int64_t viewId;
@property (nonatomic, weak) MochiViewRoot *viewRoot;
@property (nonatomic, strong) NSDate *startTime;
@end

@implementation MochiPressGestureRecognizer

- (id)initWitViewRoot:(MochiViewRoot *)viewRoot viewId:(int64_t)viewId protobuf:(GPBAny *)pb {
    NSError *error = nil;
    MochiPBPressRecognizer *pbTapRecognizer = (id)[pb unpackMessageClass:[MochiPBPressRecognizer class] error:&error];
    if (pbTapRecognizer == nil) {
        return nil;
    }
    if ((self = [super initWithTarget:self action:@selector(action:)])) {
        self.minimumPressDuration = pbTapRecognizer.minDuration.timeInterval;
        self.viewRoot = viewRoot;
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

- (void)action:(id)sender {
    CGPoint point = [self locationInView:self.view];
    
    MochiPBPressEvent *event = [[MochiPBPressEvent alloc] init];
    event.position = [[MochiPBPoint alloc] initWithCGPoint:point];
    event.timestamp = [[GPBTimestamp alloc] initWithDate:[NSDate date]];
    if (self.state == UIGestureRecognizerStateBegan) {
        event.kind = MochiPBEventKind_EventKindBegan;
        self.startTime = [NSDate date];
    } else if (self.state == UIGestureRecognizerStateChanged) {
        event.kind = MochiPBEventKind_EventKindChanged;
    } else if (self.state == UIGestureRecognizerStateEnded) {
        event.kind = MochiPBEventKind_EventKindEnded;
    } else if (self.state == UIGestureRecognizerStateCancelled) {
        event.kind = MochiPBEventKind_EventKindCancelled;
    } else {
        return;
    }
    event.duration = [[GPBDuration alloc] initWithTimeInterval:-self.startTime.timeIntervalSinceNow];
    
    NSData *data = [event data];
    MochiGoValue *value = [[MochiGoValue alloc] initWithData:data];
    
    NSLog(@"KD:%s balh", __FUNCTION__);
    [self.viewRoot call:self.funcId viewId:self.viewId args:@[value]];
    NSLog(@"KD:%s blah2", __FUNCTION__);
}


@end
