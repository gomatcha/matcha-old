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

- (NSDictionary<BridgeValue *, MochiNode *> *)nodeChildren {
    NSDictionary<BridgeValue *, BridgeValue *> *children = self.bridgeValue[@"NodeChildren"].toDictionary;
    NSMutableDictionary *nodeChildren = [NSMutableDictionary dictionary];
    for (BridgeValue *i in children.allKeys) {
        nodeChildren[i] = [[MochiNode alloc] initWithBridgeValue:children[i]];
    }
    return nodeChildren;
}

- (MochiLayoutGuide *)guide {
    return [[MochiLayoutGuide alloc] initWithBridgeValue:self.bridgeValue[@"LayoutGuide"]];
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
