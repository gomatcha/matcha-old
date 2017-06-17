#import "MatchaTextInput.h"
#import "MatchaProtobuf.h"
#import "MatchaViewController.h"

@implementation MatchaTextInput

- (id)initWithViewNode:(MatchaViewNode *)viewNode {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewNode = viewNode;
        self.delegate = self;
    }
    return self;
}

- (void)setNode:(MatchaNode *)value {
    _node = value;
    GPBAny *state = value.nativeViewState;
    NSError *error = nil;
    MatchaPBTextInputView *view = (id)[state unpackMessageClass:[MatchaPBTextInputView class] error:&error];

    self.funcId = view.onUpdate;
    NSAttributedString *attrString = [[NSAttributedString alloc] initWithProtobuf:view.styledText];
    dispatch_after(dispatch_time(DISPATCH_TIME_NOW, 1 * NSEC_PER_SEC), dispatch_get_main_queue(), ^{
        self.attributedText = attrString;
    });
}

- (void)textViewDidChange:(UITextView *)textView {
    MatchaPBTextInputEvent *event = [[MatchaPBTextInputEvent alloc] init];
    // event.value = self.attributedString;
    
    NSData *data = [event data];
    MatchaGoValue *value = [[MatchaGoValue alloc] initWithData:data];
    
    [self.viewNode.rootVC call:self.funcId viewId:self.node.identifier.longLongValue args:@[value]];
}

@end
