//
//  MochiTapGestureRecognizer.h
//  basic
//
//  Created by Kevin Dang on 5/26/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import <UIKit/UIKit.h>
#import "Touch.pbobjc.h"

@interface MochiTapGestureRecognizer : UITapGestureRecognizer
- (id)initWithViewId:(int64_t)viewId recognizerId:(int64_t)recognizerId protobuf:(GPBAny *)pb;
@property (nonatomic, assign) long long viewId;
@property (nonatomic, assign) long long recognizerId;
@end
