//
//  MochiObjcRoot.h
//  basic
//
//  Created by Kevin Dang on 4/19/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import <UIKit/UIKit.h>
@import Mochi;

@interface MochiRoot : NSObject
- (MochiGoValue *)sizeForAttributedString:(NSData *)data;
- (void)updateId:(NSInteger)identifier withProtobuf:(NSData *)protobuf;
@end
