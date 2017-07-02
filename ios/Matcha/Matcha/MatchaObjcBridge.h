//
//  MatchaObjcRoot.h
//  basic
//
//  Created by Kevin Dang on 4/19/17.
//  Copyright © 2017 Matcha. All rights reserved.
//

#import <UIKit/UIKit.h>
@import MatchaBridge;

@interface MatchaObjcBridge (Extensions)
- (void)configure;
- (MatchaGoValue *)sizeForAttributedString:(NSData *)data;
- (void)updateId:(NSInteger)identifier withProtobuf:(NSData *)protobuf;
- (NSString *)assetsDir;
- (MatchaGoValue *)imageForResource:(NSString *)path;
- (MatchaGoValue *)propertiesForResource:(NSString *)path;
@end
