#import "MatchaTextView.h"

@interface MatchaTextView ()
@property (nonatomic, weak) MatchaViewNode *viewNode;
@property (nonatomic, strong) MatchaNode *node;
@end

@implementation MatchaTextView

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
    MatchaPBStyledText *text = (id)[state unpackMessageClass:[MatchaPBStyledText class] error:&error];
    if (text != nil) {
        NSAttributedString *attrString = [[NSAttributedString alloc] initWithProtobuf:text];
        self.attributedText = attrString;
        self.numberOfLines = 0;
    }
}

@end
