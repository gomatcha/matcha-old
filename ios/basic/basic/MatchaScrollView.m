#import "MatchaScrollView.h"
#import "MatchaProtobuf.h"

@interface MatchaScrollView ()
@property (nonatomic, weak) MatchaViewNode *viewNode;
@property (nonatomic, strong) MatchaNode *node;
@end

@implementation MatchaScrollView

- (id)initWithViewNode:(MatchaViewNode *)viewNode {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewNode = viewNode;
    }
    return self;
}

- (void)setNode:(MatchaNode *)value {
    _node = value;
    
    if (self.subviews.count > 0) {
        self.contentSize = ((UIView *)self.subviews[0]).frame.size;
    }
    
    GPBAny *state = value.nativeViewState;
    NSError *error = nil;
    MatchaPBScrollView *pbscrollview = (id)[state unpackMessageClass:[MatchaPBScrollView class] error:&error];
    if (pbscrollview != nil) {
        self.scrollEnabled = pbscrollview.scrollEnabled;
        self.showsVerticalScrollIndicator = pbscrollview.showsVerticalScrollIndicator;
        self.showsHorizontalScrollIndicator = pbscrollview.showsHorizontalScrollIndicator;
        self.alwaysBounceVertical = true;
    }
}

@end
