#import "MatchaSegmentView.h"
#import "MatchaProtobuf.h"
#import "MatchaViewController.h"

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
    GPBAny *state = value.nativeViewState;
    NSError *error = nil;
    MatchaSegmentViewPbView *view = (id)[state unpackMessageClass:[MatchaSegmentViewPbView class] error:&error];
    if (view != nil) {
        self.selectedSegmentIndex = view.value;
        self.enabled = view.enabled;
        self.momentary = view.momentary;
    }
}

- (void)onChange:(id)sender {
    MatchaPBSwitchViewEvent *event = [[MatchaPBSwitchViewEvent alloc] init];
    event.value = self.selectedSegmentIndex;
    
    NSData *data = [event data];
    MatchaGoValue *value = [[MatchaGoValue alloc] initWithData:data];
    
    [self.viewNode.rootVC call:@"OnChange" viewId:self.node.identifier.longLongValue args:@[value]];
}

@end
