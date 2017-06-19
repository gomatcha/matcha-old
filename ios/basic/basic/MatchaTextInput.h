#import <UIKit/UIKit.h>
#import "MatchaView.h"

@interface MatchaTextInput : UITextView <MatchaChildView, UITextViewDelegate>

// Private
@property (nonatomic, weak) MatchaViewNode *viewNode;
@property (nonatomic, strong) MatchaNode *node;
@property (nonatomic, assign) bool hasFocus;
@end
