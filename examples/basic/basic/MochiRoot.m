//
//  MochiObjcRoot.m
//  basic
//
//  Created by Kevin Dang on 4/19/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiRoot.h"
#import "MochiBridge.h"

@implementation MochiRoot

- (MochiGoValue *)sizeForAttributedString:(MochiGoValue *)string minSize:(MochiGoValue *)minSize maxSize:(MochiGoValue *)maxSize {
    NSAttributedString *attrStr = [[NSAttributedString alloc] initWithGoValue:string];
    CGRect rect = [attrStr boundingRectWithSize:maxSize.toCGSize options:NSStringDrawingUsesLineFragmentOrigin|NSStringDrawingUsesFontLeading context:nil];
    MochiGoValue *value = [[MochiGoValue alloc] initWithCGSize:rect.size];
    return value;
}

@end
