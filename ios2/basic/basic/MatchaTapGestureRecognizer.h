//
//  MatchaTapGestureRecognizer.h
//  basic
//
//  Created by Kevin Dang on 5/26/17.
//  Copyright © 2017 Matcha. All rights reserved.
//

#import <UIKit/UIKit.h>
#import "MatchaProtobuf.h"
@class MatchaViewRoot;
@class MatchaViewController;

@interface MatchaTapGestureRecognizer : UITapGestureRecognizer
- (id)initWithMatchaVC:(MatchaViewController *)viewRoot viewId:(int64_t)viewId protobuf:(GPBAny *)pb;
- (void)disable;
- (void)updateWithProtobuf:(GPBAny *)pb;
@end
