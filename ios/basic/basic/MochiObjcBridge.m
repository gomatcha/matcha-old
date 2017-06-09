//
//  MochiObjcRoot.m
//  basic
//
//  Created by Kevin Dang on 4/19/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiObjcBridge.h"
#import "MochiBridge.h"
#import "MochiNode.h"
#import "MochiViewController.h"
#import "MochiDeadlockLogger.h"
#import "MochiProtobuf.h"

@implementation MochiObjcBridge (Extensions)

- (void)configure {
//    [MochiDeadlockLogger sharedLogger]; // Initialize
    
    static CADisplayLink *displayLink = nil;
    if (displayLink == nil) {
        displayLink = [CADisplayLink displayLinkWithTarget:self selector:@selector(screenUpdate)];
//        displayLink.preferredFramesPerSecond = 2;
        [displayLink addToRunLoop:[NSRunLoop mainRunLoop] forMode:NSRunLoopCommonModes];
    }
}

- (MochiGoValue *)sizeForAttributedString:(NSData *)protobuf {
    MochiPBSizeFunc *func = [[MochiPBSizeFunc alloc] initWithData:protobuf error:nil];
    
    NSAttributedString *attrStr = [[NSAttributedString alloc] initWithProtobuf:func.text];
    CGRect rect = [attrStr boundingRectWithSize:func.maxSize.toCGSize options:NSStringDrawingUsesLineFragmentOrigin|NSStringDrawingUsesFontLeading context:nil];
    
    MochiPBPoint *point = [[MochiPBPoint alloc] initWithCGSize:rect.size];
    return [[MochiGoValue alloc] initWithData:point.data];
}

- (void)screenUpdate {
    static MochiGoValue *updateFunc = nil;
    if (updateFunc == nil) {
        updateFunc = [[MochiGoValue alloc] initWithFunc:@"github.com/overcyn/mochi/animate screenUpdate"];
    }
    [updateFunc call:nil args:nil];
}

- (void)updateId:(NSInteger)identifier withProtobuf:(NSData *)protobuf {
    MochiPBRoot *pbroot = [[MochiPBRoot alloc] initWithData:protobuf error:nil];
    MochiNodeRoot *root = [[MochiNodeRoot alloc] initWithProtobuf:pbroot];
    
    MochiViewController *vc = [MochiViewController viewControllerWithIdentifier:identifier];
    [vc update:root.node];
}

- (NSString *)assetsDir {
     return [[NSBundle mainBundle] resourcePath];
}

@end
