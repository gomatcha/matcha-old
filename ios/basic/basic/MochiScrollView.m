#import "MochiScrollView.h"
#import "MochiProtobuf.h"

@interface MochiScrollView ()
@property (nonatomic, weak) MochiViewNode *viewNode;
@property (nonatomic, strong) MochiNode *node;
@end

@implementation MochiScrollView

- (id)initWithViewNode:(MochiViewNode *)viewNode {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewNode = viewNode;
    }
    return self;
}

- (void)setNode:(MochiNode *)value {
    _node = value;
    
    if (self.subviews.count > 0) {
        self.contentSize = ((UIView *)self.subviews[0]).frame.size;
    }
    
    GPBAny *state = value.nativeViewState;
    NSError *error = nil;
    MochiPBScrollView *pbscrollview = (id)[state unpackMessageClass:[MochiPBScrollView class] error:&error];
    if (pbscrollview != nil) {
        self.scrollEnabled = pbscrollview.scrollEnabled;
        self.showsVerticalScrollIndicator = pbscrollview.showsVerticalScrollIndicator;
        self.showsHorizontalScrollIndicator = pbscrollview.showsHorizontalScrollIndicator;
        self.alwaysBounceVertical = true;
    }
}

@end
