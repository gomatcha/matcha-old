//
//  MatchaObjcRoot.h
//  basic
//
//  Created by Kevin Dang on 4/19/17.
//  Copyright © 2017 Matcha. All rights reserved.
//

#import <UIKit/UIKit.h>
@import Matcha;

@interface MatchaObjcBridge (Extensions)
- (void)configure;
- (MatchaGoValue *)sizeForAttributedString:(NSData *)data;
- (void)updateId:(NSInteger)identifier withProtobuf:(NSData *)protobuf;
- (NSString *)assetsDir;
- (MatchaGoValue *)sizeForResource:(NSString *)path;
- (MatchaGoValue *)imageForResource:(NSString *)path;
@end
