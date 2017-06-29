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
    [MatchaDeadlockLogger sharedLogger]; // Initialize
    
    static CADisplayLink *displayLink = nil;
    if (displayLink == nil) {
        displayLink = [CADisplayLink displayLinkWithTarget:self selector:@selector(screenUpdate)];
//        displayLink.preferredFramesPerSecond = 1;
        [displayLink addToRunLoop:[NSRunLoop mainRunLoop] forMode:NSRunLoopCommonModes];
    }
}

- (MatchaGoValue *)sizeForAttributedString:(NSData *)protobuf {
    MatchaPBSizeFunc *func = [[MatchaPBSizeFunc alloc] initWithData:protobuf error:nil];
    
    NSAttributedString *attrStr = [[NSAttributedString alloc] initWithProtobuf:func.text];
    CGRect rect = [attrStr boundingRectWithSize:func.maxSize.toCGSize options:NSStringDrawingUsesLineFragmentOrigin|NSStringDrawingUsesFontLeading context:nil];
    
    MatchaLayoutPBPoint *point = [[MatchaLayoutPBPoint alloc] initWithCGSize:CGSizeMake(ceil(rect.size.width), ceil(rect.size.height))];
    return [[MatchaGoValue alloc] initWithData:point.data];
}

- (void)screenUpdate {
    static MatchaGoValue *updateFunc = nil;
    if (updateFunc == nil) {
        updateFunc = [[MatchaGoValue alloc] initWithFunc:@"github.com/gomatcha/matcha/animate screenUpdate"];
    }
    [updateFunc call:nil args:nil];
}

- (void)updateId:(NSInteger)identifier withProtobuf:(NSData *)protobuf {
    MatchaViewPBRoot *pbroot = [[MatchaViewPBRoot alloc] initWithData:protobuf error:nil];
    MatchaNodeRoot *root = [[MatchaNodeRoot alloc] initWithProtobuf:pbroot];
    
    MatchaViewController *vc = [MatchaViewController viewControllerWithIdentifier:identifier];
    [vc update:root.node];
}

- (NSString *)assetsDir {
     return [[NSBundle mainBundle] resourcePath];
}

- (MatchaGoValue *)imageForResource:(NSString *)path {
    UIImage *image = [UIImage imageNamed:path];
    if (image == nil) {
        return nil;
    }
    NSData *data = UIImagePNGRepresentation(image);
    return [[MatchaGoValue alloc] initWithData:data];
}

- (MatchaGoValue *)propertiesForResource:(NSString *)path {
    UIImage *image = [UIImage imageNamed:path];
    if (image == nil) {
        return nil;
    }
    MatchaPBImageProperties *props = [[MatchaPBImageProperties alloc] init];
    props.width = ceil(image.size.width * image.scale);
    props.height = ceil(image.size.height * image.scale);
    props.scale = image.scale;
    return [[MatchaGoValue alloc] initWithData:props.data];
}

@end
