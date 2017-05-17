//
//  MochiNode.m
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiNode.h"
#import "MochiBridge.h"

@interface MochiNode ()
@property (nonatomic, strong) NSDictionary *nodeChildren;
@property (nonatomic, strong) MochiGoValue *goValue;
@property (nonatomic, strong) MochiLayoutGuide *guide;
@property (nonatomic, strong) MochiPaintOptions *paintOptions;
@property (nonatomic, strong) NSNumber *identifier;
@property (nonatomic, strong) NSNumber *buildId;
@property (nonatomic, strong) NSNumber *layoutId;
@property (nonatomic, strong) NSNumber *paintId;
@end

@implementation MochiNode

- (id)initWithGoValue:(MochiGoValue *)value {
    if (self = [super init]) {
        self.goValue = value;
        self.identifier = @(value[@"Id"].toLongLong);
        self.buildId = @(value[@"BuildId"].toLongLong);
        self.layoutId = @(value[@"LayoutId"].toLongLong);
        self.paintId = @(value[@"PaintId"].toLongLong);
    }
    return self;
}

- (MochiPaintOptions *)paintOptions {
    if (_paintOptions == nil) {
        _paintOptions = [[MochiPaintOptions alloc] initWithGoValue:self.goValue[@"PaintOptions"]];
    }
    return _paintOptions;
}

- (MochiLayoutGuide *)guide {
    if (_guide == nil) {
        _guide = [[MochiLayoutGuide alloc] initWithGoValue:self.goValue[@"LayoutGuide"]];
    }
    return _guide;
}

- (NSDictionary<NSNumber *, MochiNode *> *)nodeChildren {
    if (_nodeChildren == nil) {
        NSMapTable *children = self.goValue[@"Children"].toMapTable;
        NSMutableDictionary<NSNumber *, MochiNode *> *nodeChildren = [NSMutableDictionary dictionary];
        for (MochiGoValue *i in children) {
            nodeChildren[@(i.toLongLong)] = [[MochiNode alloc] initWithGoValue:children[i]];
        }
        _nodeChildren = nodeChildren;
    }
    return _nodeChildren;
}

- (NSString *)bridgeName {
    return self.goValue[@"BridgeName"].toString;
}

- (MochiGoValue *)bridgeState {
    return self.goValue[@"BridgeState"];
}

@end

@interface MochiPaintOptions ()
@property (nonatomic, strong) MochiGoValue *goValue;
@end

@implementation MochiPaintOptions

- (id)initWithGoValue:(MochiGoValue *)value {
    if (self = [super init]) {
        self.goValue = value;
    }
    return self;
}

- (UIColor *)backgroundColor {
    MochiGoValue *value = self.goValue[@"BackgroundColor"];
    if (!value.isNil) {
        return [[UIColor alloc] initWithGoValue:value];
    }
    return nil;
}

@end


@interface MochiLayoutGuide ()
@property (nonatomic, assign) CGRect frame;
@property (nonatomic, assign) UIEdgeInsets insets;
@property (nonatomic, assign) NSInteger zIndex;
@end

@implementation MochiLayoutGuide

- (id)initWithGoValue:(MochiGoValue *)value {
    if (value.isNil) {
        return nil;
    }
    if (self = [super init]) {
        self.frame = value[@"Frame"].toCGRect;
        self.insets = value[@"Insets"].toUIEdgeInsets;
        self.zIndex = value[@"ZIndex"].toLongLong;
    }
    return self;
}

@end
