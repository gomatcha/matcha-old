#import <UIKit/UIKit.h>
#import "MatchaView.h"

@interface MatchaScrollView : UIScrollView <MatchaChildView, UIScrollViewDelegate>
@property (nonatomic, weak) MatchaViewNode *viewNode;
@property (nonatomic, strong) MatchaNode *node;
@property (nonatomic, assign) BOOL scrollEvents;
@end
