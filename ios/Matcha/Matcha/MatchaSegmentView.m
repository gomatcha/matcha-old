#import "MatchaSegmentView.h"

@implementation MatchaSegmentView

- (id)initWithViewNode:(MatchaViewNode *)viewNode {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewNode = viewNode;
        [self addTarget:self action: @selector(onChange:) forControlEvents:UIControlEventValueChanged];
    }
    return self;
}

- (void)setNode:(MatchaNode *)value {
    _node = value;

    NSError *error = nil;
    MatchaSegmentViewPbView *view = (id)[value.nativeViewState unpackMessageClass:[MatchaSegmentViewPbView class] error:&error];
    if (error != nil) {
        NSLog(@"Error:%@", error);
    }
    
    self.selectedSegmentIndex = view.value;
    self.enabled = view.enabled;
    self.momentary = view.momentary;
    [self removeAllSegments];
    for (NSInteger i = 0; i < view.titlesArray.count; i++) {
        [self insertSegmentWithTitle:view.titlesArray[i] atIndex:i animated:NO];
    }
}

- (void)onChange:(id)sender {
    MatchaSegmentViewPbEvent *event = [[MatchaSegmentViewPbEvent alloc] init];
    event.value = self.selectedSegmentIndex;
    MatchaGoValue *value = [[MatchaGoValue alloc] initWithData:event.data];
    
    [self.viewNode.rootVC call:@"OnChange" viewId:self.node.identifier.longLongValue args:@[value]];
}

@end
