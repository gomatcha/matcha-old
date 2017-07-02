#import <UIKit/UIKit.h>
#import "MatchaView.h"

@interface MatchaSwitchView : UISwitch <MatchaChildView>

// Private
@property (nonatomic, weak) MatchaViewNode *viewNode;
@property (nonatomic, strong) MatchaNode *node;
@end
