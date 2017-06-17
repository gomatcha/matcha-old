//
//  MatchaObjcRoot.m
//  basic
//
//  Created by Kevin Dang on 4/19/17.
//  Copyright © 2017 Matcha. All rights reserved.
//

#import "MatchaObjcBridge.h"
#import "MatchaBridge.h"
#import "MatchaNode.h"
#import "MatchaViewController.h"
#import "MatchaDeadlockLogger.h"
#import "MatchaProtobuf.h"

@implementation MatchaObjcBridge (Extensions)

- (void)configure {
//    [MatchaDeadlockLogger sharedLogger]; // Initialize
    
    static CADisplayLink *displayLink = nil;
    if (displayLink == nil) {
        displayLink = [CADisplayLink displayLinkWithTarget:self selector:@selector(screenUpdate)];
//        displayLink.preferredFramesPerSecond = 2;
        [displayLink addToRunLoop:[NSRunLoop mainRunLoop] forMode:NSRunLoopCommonModes];
    }
}

- (MatchaGoValue *)sizeForAttributedString:(NSData *)protobuf {
    MatchaPBSizeFunc *func = [[MatchaPBSizeFunc alloc] initWithData:protobuf error:nil];
    
    NSAttributedString *attrStr = [[NSAttributedString alloc] initWithProtobuf:func.text];
    CGRect rect = [attrStr boundingRectWithSize:func.maxSize.toCGSize options:NSStringDrawingUsesLineFragmentOrigin|NSStringDrawingUsesFontLeading context:nil];
    
    MatchaPBPoint *point = [[MatchaPBPoint alloc] initWithCGSize:CGSizeMake(ceil(rect.size.width), ceil(rect.size.height))];
    return [[MatchaGoValue alloc] initWithData:point.data];
}

- (void)screenUpdate {
    static MatchaGoValue *updateFunc = nil;
    if (updateFunc == nil) {
        updateFunc = [[MatchaGoValue alloc] initWithFunc:@"github.com/overcyn/matcha/animate screenUpdate"];
    }
    [updateFunc call:nil args:nil];
}

- (void)updateId:(NSInteger)identifier withProtobuf:(NSData *)protobuf {
    MatchaPBRoot *pbroot = [[MatchaPBRoot alloc] initWithData:protobuf error:nil];
    MatchaNodeRoot *root = [[MatchaNodeRoot alloc] initWithProtobuf:pbroot];
    
    MatchaViewController *vc = [MatchaViewController viewControllerWithIdentifier:identifier];
    [vc update:root.node];
}

- (NSString *)assetsDir {
     return [[NSBundle mainBundle] resourcePath];
}

- (MatchaGoValue *)sizeForResource:(NSString *)path {
    UIImage *image = [UIImage imageNamed:path];
    
    MatchaPBPoint *point = [[MatchaPBPoint alloc] initWithCGSize:CGSizeMake(ceil(image.size.width / image.scale), ceil(image.size.height / image.scale))];
    return [[MatchaGoValue alloc] initWithData:point.data];
}

@end
