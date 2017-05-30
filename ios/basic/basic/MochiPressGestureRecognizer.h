//
//  MochiPressGestureRecognizer.h
//  basic
//
//  Created by Kevin Dang on 5/30/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import <UIKit/UIKit.h>
@class GPBAny;
@class MochiViewRoot;

@interface MochiPressGestureRecognizer : UILongPressGestureRecognizer
- (id)initWitViewRoot:(MochiViewRoot *)viewRoot viewId:(int64_t)viewId protobuf:(GPBAny *)pb;
- (void)updateWithProtobuf:(GPBAny *)pb;
- (void)disable;
@end
