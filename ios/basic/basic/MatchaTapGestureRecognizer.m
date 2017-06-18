//
//  MatchaTapGestureRecognizer.m
//  basic
//
//  Created by Kevin Dang on 5/26/17.
//  Copyright © 2017 Matcha. All rights reserved.
//

#import "MatchaTapGestureRecognizer.h"
#import <Matcha/matchago.h>
#import "MatchaNode.h"
#import "MatchaProtobuf.h"
#import "MatchaBridge.h"
#import "MatchaViewController.h"

@interface MatchaTapGestureRecognizer ()
@property (nonatomic, assign) int64_t funcId;
@property (nonatomic, assign) int64_t viewId;
@property (nonatomic, weak) MatchaViewController *viewController;
@property (nonatomic, assign) bool disabled;
@end

@implementation MatchaTapGestureRecognizer

- (id)initWithMatchaVC:(MatchaViewController *)viewController viewId:(int64_t)viewId protobuf:(GPBAny *)pb {
    NSError *error = nil;
    MatchaPBTouchTapRecognizer *pbTapRecognizer = (id)[pb unpackMessageClass:[MatchaPBTouchTapRecognizer class] error:&error];
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
    MatchaPBTouchTapRecognizer *pbTapRecognizer = (id)[pb unpackMessageClass:[MatchaPBTouchTapRecognizer class] error:&error];
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
    
    MatchaPBTouchTapEvent *event = [[MatchaPBTouchTapEvent alloc] init];
    event.position = [[MatchaLayoutPBPoint alloc] initWithCGPoint:point];
    event.timestamp = [[GPBTimestamp alloc] initWithDate:[NSDate date]];
    
    NSData *data = [event data];
    MatchaGoValue *value = [[MatchaGoValue alloc] initWithData:data];
    
    [self.viewController call:self.funcId viewId:self.viewId args:@[value]];
}

@end
