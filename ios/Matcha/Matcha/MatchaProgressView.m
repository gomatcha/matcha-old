#import "MatchaProgressView.h"
#import "MatchaProtobuf.h"

@implementation MatchaProgressView

- (id)initWithViewNode:(MatchaViewNode *)viewNode {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewNode = viewNode;        
    }
    return self;
}

- (void)setNode:(MatchaNode *)value {
    _node = value;
    GPBAny *state = value.nativeViewState;
    NSError *error = nil;
    MatchaProgressViewPBView *view = (id)[state unpackMessageClass:[MatchaProgressViewPBView class] error:&error];
    if (view != nil) {
        self.progress = view.progress;
    }
}

@end
