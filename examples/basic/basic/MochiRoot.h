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
- (MochiGoValue *)sizeForAttributedString:(MochiGoValue *)string minSize:(MochiGoValue *)minSize maxSize:(MochiGoValue *)maxSize;
@end
