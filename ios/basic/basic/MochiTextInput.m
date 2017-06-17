#import "MochiTextInput.h"
#import "MochiProtobuf.h"
#import "MochiViewController.h"

@implementation MochiTextInput

- (id)initWithViewNode:(MochiViewNode *)viewNode {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewNode = viewNode;
        self.delegate = self;
    }
    return self;
}

- (void)setNode:(MochiNode *)value {
    _node = value;
    GPBAny *state = value.nativeViewState;
    NSError *error = nil;
    MochiPBTextInputView *view = (id)[state unpackMessageClass:[MochiPBTextInputView class] error:&error];

    self.funcId = view.onUpdate;
    NSAttributedString *attrString = [[NSAttributedString alloc] initWithProtobuf:view.styledText];
    dispatch_after(dispatch_time(DISPATCH_TIME_NOW, 1 * NSEC_PER_SEC), dispatch_get_main_queue(), ^{
        self.attributedText = attrString;
    });
}

- (void)textViewDidChange:(UITextView *)textView {
    MochiPBTextInputEvent *event = [[MochiPBTextInputEvent alloc] init];
    // event.value = self.attributedString;
    
    NSData *data = [event data];
    MochiGoValue *value = [[MochiGoValue alloc] initWithData:data];
    
    [self.viewNode.rootVC call:self.funcId viewId:self.node.identifier.longLongValue args:@[value]];
}

@end
