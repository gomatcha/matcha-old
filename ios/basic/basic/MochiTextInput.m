#import "MochiTextInput.h"
#import "MochiProtobuf.h"

@implementation MochiTextInput

- (id)initWithViewNode:(MochiViewNode *)viewNode {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewNode = viewNode;
        
    }
    return self;
}

- (void)setNode:(MochiNode *)value {
    _node = value;
    GPBAny *state = value.nativeViewState;
    NSError *error = nil;
    MochiPBTextInputView *view = (id)[state unpackMessageClass:[MochiPBTextInputView class] error:&error];
    if (view != nil) {
//        self.funcId = view.onValueChange;
    }
}


@end
