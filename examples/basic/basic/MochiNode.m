//
//  MochiNode.m
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiNode.h"

@interface MochiNode ()
@property (nonatomic, strong) BridgeValue *bridgeValue;
@end

@implementation MochiNode

- (id)initWithBridgeValue:(BridgeValue *)value {
    if (self = [super init]) {
        self.bridgeValue = value;
    }
    return self;
}

- (MochiPaintOptions *)paintOptions {
    return [[MochiPaintOptions alloc] initWithBridgeValue:self.bridgeValue[@"PaintOptions"]];
}

- (NSArray<MochiNode *> *)nodeChildren {
    NSArray<BridgeValue *> *children = self.bridgeValue[@"NodeChildren"].toArray;
    NSMutableArray *nodeChildren = [NSMutableArray array];
    for (BridgeValue *i in children) {
        [nodeChildren addObject:[[MochiNode alloc] initWithBridgeValue:i]];
    }
    return nodeChildren;
}

@end

@interface MochiPaintOptions ()
@property (nonatomic, strong) BridgeValue *bridgeValue;
@end

@implementation MochiPaintOptions

- (id)initWithBridgeValue:(BridgeValue *)value {
    if (self = [super init]) {
        self.bridgeValue = value;
    }
    return self;
}

- (UIColor *)backgroundColor {
    return [[UIColor alloc] initWithBridgeValue:self.bridgeValue[@"BackgroundColor"]];
}

@end


@interface MochiLayoutGuide ()
@property (nonatomic, assign) CGRect frame;
@property (nonatomic, assign) UIEdgeInsets insets;
@end

@implementation MochiLayoutGuide

- (id)initWithBridgeValue:(BridgeValue *)value {
    if (self = [super init]) {
        self.frame = value[@"Frame"].toCGRect;
        self.insets = value[@"Insets"].toUIEdgeInsets;
    }
    return self;
}

@end
