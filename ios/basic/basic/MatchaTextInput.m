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
    MatchaTextInputPBView *view = (id)[state unpackMessageClass:[MatchaTextInputPBView class] error:&error];
    
    NSAttributedString *attrString = [[NSAttributedString alloc] initWithProtobuf:view.styledText];
    if (![attrString.string isEqual:self.attributedText.string]) { // TODO(KD): Better comparison.
        self.attributedText = attrString;
    }
    self.attrStr2 = attrString;
    self.hasFocus = view.focused;
    
    if (self.hasFocus && !self.isFirstResponder) {
        [self becomeFirstResponder];
    } else if (!self.hasFocus && self.isFirstResponder) {
        [self resignFirstResponder];
    }
}

- (void)textViewDidChange:(UITextView *)textView {
    if ([self.attributedText isEqual:self.attrStr2] || self.attributedText == self.attrStr2) {
        return;
    }
    MatchaTextInputPBEvent *event = [[MatchaTextInputPBEvent alloc] init];
    event.styledText = self.attributedText.protobuf;
    
    NSData *data = [event data];
    MatchaGoValue *value = [[MatchaGoValue alloc] initWithData:data];
    
    [self.viewNode.rootVC call:@"OnChange" viewId:self.node.identifier.longLongValue args:@[value]];
}

- (void)textViewDidBeginEditing:(UITextView *)textView {
    [self focusDidChange];
}

- (void)textViewDidEndEditing:(UITextView *)textView {
    [self focusDidChange];
}

- (void)focusDidChange {
    if ((self.hasFocus && !self.isFirstResponder) || (!self.hasFocus && self.isFirstResponder)) {
        MatchaTextInputPBFocusEvent *event = [[MatchaTextInputPBFocusEvent alloc] init];
        event.focused = self.isFirstResponder;
        
        MatchaGoValue *value = [[MatchaGoValue alloc] initWithData:event.data];
        [self.viewNode.rootVC call:@"OnFocus" viewId:self.node.identifier.longLongValue args:@[value]];
    }
}

@end
