#import "MochiTabBarController.h"
#import "MochiView.h"

@implementation MochiTabBarController

- (id)initWithViewNode:(MochiViewNode *)viewNode {
    if ((self = [super init])) {
        self.viewNode = viewNode;
    }
    return self;
}

@end
