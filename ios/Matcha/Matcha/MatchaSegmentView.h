#import <UIKit/UIKit.h>
#import "MatchaView.h"

@interface MatchaSegmentView : UISegmentedControl <MatchaChildView>
@property (nonatomic, weak) MatchaViewNode *viewNode;
@property (nonatomic, strong) MatchaNode *node;
@end
