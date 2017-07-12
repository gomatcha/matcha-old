#import "MatchaBasicView.h"

@interface MatchaBasicView ()
@property (nonatomic, weak) MatchaViewNode *viewNode;
@property (nonatomic, strong) MatchaNode *node;
@end

@implementation MatchaBasicView

+ (void)load {
    MatchaRegisterView(@"", ^(MatchaViewNode *node){
        return [[MatchaBasicView alloc] initWithViewNode:node];
    });
}

- (id)initWithViewNode:(MatchaViewNode *)viewNode {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewNode = viewNode;
    }
    return self;
}

- (void)setNode:(MatchaNode *)value {
    _node = value;
}

@end

