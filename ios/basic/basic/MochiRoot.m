//
//  MochiObjcRoot.m
//  basic
//
//  Created by Kevin Dang on 4/19/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiRoot.h"
#import "MochiBridge.h"
#import "MochiNode.h"
#import "MochiViewController.h"
#import "MochiDeadlockLogger.h"
#import "MochiProtobuf.h"

@interface MochiRoot ()
@property (nonatomic, strong) CADisplayLink *displayLink;
@property (nonatomic, strong) MochiGoValue *screenUpdateFunc;
@end

@implementation MochiRoot

- (id)init {
    if ((self = [super init])) {
//        [MochiDeadlockLogger sharedLogger]; // Initialize
        
        self.displayLink = [CADisplayLink displayLinkWithTarget:self selector:@selector(screenUpdate)];
        [self.displayLink addToRunLoop:[NSRunLoop mainRunLoop] forMode:NSRunLoopCommonModes];
//        self.displayLink.preferredFramesPerSecond = 2;
        self.screenUpdateFunc = [[MochiGoValue alloc] initWithFunc:@"github.com/overcyn/mochi/animate screenUpdate"];
    }
    return self;
}

- (MochiGoValue *)sizeForAttributedString:(NSData *)protobuf {
    MochiPBSizeFunc *func = [[MochiPBSizeFunc alloc] initWithData:protobuf error:nil];
    
    NSAttributedString *attrStr = [[NSAttributedString alloc] initWithProtobuf:func.text];
    CGRect rect = [attrStr boundingRectWithSize:func.maxSize.toCGSize options:NSStringDrawingUsesLineFragmentOrigin|NSStringDrawingUsesFontLeading context:nil];
    
    MochiPBPoint *point = [[MochiPBPoint alloc] initWithCGSize:rect.size];
    return [[MochiGoValue alloc] initWithData:point.data];
}

- (void)screenUpdate {
    [self.screenUpdateFunc call:nil args:nil];
}

- (void)updateId:(NSInteger)identifier withProtobuf:(NSData *)protobuf {
    MochiPBRoot *pbroot = [[MochiPBRoot alloc] initWithData:protobuf error:nil];
    MochiNodeRoot *root = [[MochiNodeRoot alloc] initWithProtobuf:pbroot];
    
    MochiViewController *vc = [MochiViewController viewControllerWithIdentifier:identifier];
    [vc update:root.node];
}

@end
