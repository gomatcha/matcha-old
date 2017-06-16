#import <UIKit/UIKit.h>
#import "MochiView.h"

@interface MochiTextInput : UIView <MochiChildView>

// Private
@property (nonatomic, weak) MochiViewNode *viewNode;
@property (nonatomic, strong) MochiNode *node;
@property (nonatomic, assign) int64_t funcId;
@end
