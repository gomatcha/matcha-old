#import "MatchaSlider.h"
#import "MatchaViewController.h"

@implementation MatchaSlider

- (id)initWithViewNode:(MatchaViewNode *)viewNode {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewNode = viewNode;
        [self addTarget:self action: @selector(onChange:) forControlEvents: UIControlEventValueChanged];
        
    }
    return self;
}

- (void)setNode:(MatchaNode *)value {
    _node = value;
    GPBAny *state = value.nativeViewState;
    NSError *error = nil;
    MatchaSliderPbView *view = (id)[state unpackMessageClass:[MatchaSliderPbView class] error:&error];
    if (view != nil) {
        self.enabled = view.enabled;
        self.value = view.value;
        self.maximumValue = view.maxValue;
        self.minimumValue = view.minValue;
    }
}

- (void)onChange:(id)sender {
    MatchaSliderPbEvent *event = [[MatchaSliderPbEvent alloc] init];
    event.value = self.value;
    
    NSData *data = [event data];
    MatchaGoValue *value = [[MatchaGoValue alloc] initWithData:data];
    
    [self.viewNode.rootVC call:@"OnValueChange" viewId:self.node.identifier.longLongValue args:@[value]];
}

@end
