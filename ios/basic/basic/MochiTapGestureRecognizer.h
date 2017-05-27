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
- (id)initWithProtobuf:(GPBAny *)pb;
@end
