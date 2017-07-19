#import "MatchaSlider.h"
#import "MatchaViewController.h"

@implementation MatchaSlider

+ (void)load {
    MatchaRegisterView(@"gomatcha.io/matcha/view/slider", ^(MatchaViewNode *node){
        return [[MatchaSlider alloc] initWithViewNode:node];
    });
}

- (id)initWithViewNode:(MatchaViewNode *)viewNode {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewNode = viewNode;
        [self addTarget:self action:@selector(onChange:forEvent:) forControlEvents:UIControlEventValueChanged];
        
    }
    return self;
}

- (void)setNode:(MatchaBuildNode *)value {
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

- (void)onChange:(id)sender forEvent:(UIEvent *)e {
    MatchaSliderPbEvent *event = [[MatchaSliderPbEvent alloc] init];
    event.value = self.value;
    
    NSData *data = [event data];
    MatchaGoValue *value = [[MatchaGoValue alloc] initWithData:data];
    
    [self.viewNode.rootVC call:@"OnValueChange" viewId:self.node.identifier.longLongValue args:@[value]];
    
    UITouch *touchEvent = [[e allTouches] anyObject];
    if (touchEvent.phase == UITouchPhaseEnded) {
        [self.viewNode.rootVC call:@"OnSubmit" viewId:self.node.identifier.longLongValue args:@[value]];
    }
}

@end
