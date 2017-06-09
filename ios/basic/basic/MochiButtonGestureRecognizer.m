#import "MochiButtonGestureRecognizer.h"

@interface MochiButtonGestureRecognizer ()
@property (nonatomic, assign) int64_t funcId;
@property (nonatomic, assign) int64_t viewId;
@property (nonatomic, weak) MochiViewController *viewController;
@property (nonatomic, strong) NSDate *startTime;
@property (nonatomic, assign) BOOL disabled;
@property (nonatomic, assign) BOOL inside;
@end


@implementation MochiButtonGestureRecognizer

- (id)initWithMochiVC:(MochiViewController *)viewController viewId:(int64_t)viewId protobuf:(GPBAny *)pbany {
    NSError *error = nil;
    MochiPBTouchButtonRecognizer *pb = (id)[pbany unpackMessageClass:[MochiPBTouchButtonRecognizer class] error:&error];
    if (pb == nil) {
        return nil;
    }
    if ((self = [super initWithTarget:self action:@selector(action:)])) {
        self.viewController = viewController;
        self.funcId = pb.onEvent;
        self.viewId = viewId;
        self.allowableMovement = 10000;
    }
    return self;
}

- (void)updateWithProtobuf:(GPBAny *)pbany {
    NSError *error = nil;
    MochiPBTouchButtonRecognizer *pb = (id)[pbany unpackMessageClass:[MochiPBTouchButtonRecognizer class] error:&error];
    if (pb == nil) {
        return;
    }
    self.funcId = pb.onEvent;
}

- (void)disable {
    self.disabled = false;
}

- (void)action:(id)sender {
    if (self.disabled) {
        return;
    }
    
    CGPoint point = [self locationInView:self.view];
    BOOL prevInside = self.inside;
    self.inside = CGRectContainsPoint(self.view.bounds, point);
    
    MochiPBTouchButtonEvent *event = [[MochiPBTouchButtonEvent alloc] init];
    event.position = [[MochiPBPoint alloc] initWithCGPoint:point];
    event.timestamp = [[GPBTimestamp alloc] initWithDate:[NSDate date]];
    if (self.state == UIGestureRecognizerStateBegan) {
        event.kind = MochiPBEventKind_EventKindBegan;
        self.startTime = [NSDate date];
    } else if (self.state == UIGestureRecognizerStateChanged && self.inside != prevInside) { // Only update if inside has changed
        event.kind = MochiPBEventKind_EventKindChanged;
    } else if (self.state == UIGestureRecognizerStateEnded) {
        if (self.inside) {
            event.kind = MochiPBEventKind_EventKindEnded;
        } else {
            event.kind = MochiPBEventKind_EventKindCancelled;
        }
    } else if (self.state == UIGestureRecognizerStateCancelled) {
        event.kind = MochiPBEventKind_EventKindCancelled;
    } else {
        return;
    }
    
    NSData *data = [event data];
    MochiGoValue *value = [[MochiGoValue alloc] initWithData:data];
    
    [self.viewController call:self.funcId viewId:self.viewId args:@[value]];
}

@end
