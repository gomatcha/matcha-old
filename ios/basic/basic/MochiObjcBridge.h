//
//  MochiObjcRoot.h
//  basic
//
//  Created by Kevin Dang on 4/19/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import <UIKit/UIKit.h>
@import Mochi;

@interface MochiObjcBridge (Extensions)
- (void)configure;
- (MochiGoValue *)sizeForAttributedString:(NSData *)data;
- (void)updateId:(NSInteger)identifier withProtobuf:(NSData *)protobuf;
- (NSString *)assetsDir;
- (MochiGoValue *)sizeForResource:(NSString *)path;
@end
