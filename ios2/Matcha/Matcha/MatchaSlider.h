//
//  MatchaSlider.h
//  basic
//
//  Created by Kevin Dang on 6/27/17.
//  Copyright © 2017 Matcha. All rights reserved.
//

#import <UIKit/UIKit.h>
#import "MatchaView.h"
#import "MatchaProtobuf.h"

@interface MatchaSlider : UISlider <MatchaChildView>
@property (nonatomic, weak) MatchaViewNode *viewNode;
@property (nonatomic, strong) MatchaNode *node;
@end
