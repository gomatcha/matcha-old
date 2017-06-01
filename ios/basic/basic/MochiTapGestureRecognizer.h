//
//  MochiTapGestureRecognizer.h
//  basic
//
//  Created by Kevin Dang on 5/26/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import <UIKit/UIKit.h>
#import "Touch.pbobjc.h"
@class MochiViewRoot;
@class MochiViewController;

@interface MochiTapGestureRecognizer : UITapGestureRecognizer
- (id)initWitViewRoot:(MochiViewController *)viewRoot viewId:(int64_t)viewId protobuf:(GPBAny *)pb;
- (void)disable;
- (void)updateWithProtobuf:(GPBAny *)pb;
@end
