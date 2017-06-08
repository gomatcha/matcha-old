#import "MochiSwitchView.h"
#import "MochiProtobuf.h"
#import "MochiViewController.h"

@implementation MochiSwitchView

- (id)initWithViewNode:(MochiViewNode *)viewNode {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewNode = viewNode;
        [self addTarget:self action: @selector(onChange:) forControlEvents: UIControlEventValueChanged];

    }
    return self;
}

- (void)setNode:(MochiNode *)value {
    _node = value;
    GPBAny *state = value.nativeViewState;
    NSError *error = nil;
    MochiPBSwitchViewView *view = (id)[state unpackMessageClass:[MochiPBSwitchViewView class] error:&error];
    if (view != nil) {
        self.on = view.value;
    }
    self.funcId = view.onValueChange;
}

- (void)onChange:(id)sender {
    MochiPBSwitchViewEvent *event = [[MochiPBSwitchViewEvent alloc] init];
    event.value = self.on;
    
    NSData *data = [event data];
    MochiGoValue *value = [[MochiGoValue alloc] initWithData:data];
    
    [self.viewNode.rootVC call:self.funcId viewId:self.node.identifier.longLongValue args:@[value]];

}

@end
