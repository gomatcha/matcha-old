#import "MochiButtonGestureRecognizer.h"

@interface MochiButtonGestureRecognizer () <UIGestureRecognizerDelegate>
@property (nonatomic, assign) int64_t funcId;
@property (nonatomic, assign) int64_t viewId;
@property (nonatomic, weak) MochiViewController *viewController;
@property (nonatomic, strong) NSDate *startTime;
@property (nonatomic, assign) BOOL disabled;
@property (nonatomic, assign) BOOL inside;
@property (nonatomic, assign) BOOL ignoresScroll;
@end


@implementation MochiButtonGestureRecognizer

- (id)initWithMochiVC:(MochiViewController *)viewController viewId:(int64_t)viewId protobuf:(GPBAny *)pbany {
    NSError *error = nil;
    MochiPBTouchButtonRecognizer *pb = (id)[pbany unpackMessageClass:[MochiPBTouchButtonRecognizer class] error:&error];
    if (pb == nil) {
        return nil;
    }
    if ((self = [super initWithTarget:self action:@selector(action:)])) {
        self.delegate = self;
        self.viewController = viewController;
        self.funcId = pb.onEvent;
        self.viewId = viewId;
        self.minimumPressDuration = 0;
        self.allowableMovement = 10000;
        self.ignoresScroll = pb.ignoresScroll;
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
    event.timestamp = [[GPBTimestamp alloc] initWithDate:[NSDate date]];
    event.inside = self.inside;
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

#pragma mark - UIGestureRecognizerDelegate
//
//- (BOOL)canBePreventedByGestureRecognizer:(UIGestureRecognizer *)preventingGestureRecognizer {
//    NSLog(@"%@",preventingGestureRecognizer);
//    return [preventingGestureRecognizer.view isKindOfClass:UIScrollView.class];
//}
//}

//- (BOOL)gestureRecognizer:(UIGestureRecognizer *)gestureRecognizer shouldRequireFailureOfGestureRecognizer:(UIGestureRecognizer *)otherGestureRecognizer {
//    return [otherGestureRecognizer.view isKindOfClass:UIScrollView.class];
//}

//- (BOOL)gestureRecognizer:(UIPanGestureRecognizer *)gestureRecognizer shouldRecognizeSimultaneouslyWithGestureRecognizer:(UISwipeGestureRecognizer *)otherGestureRecognizer {
//    return YES;
//}

@end
