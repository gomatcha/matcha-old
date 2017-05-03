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
@property (nonatomic, strong) NSMapTable *nodeChildren;
@property (nonatomic, strong) MochiGoValue *goValue;
@property (nonatomic, strong) MochiLayoutGuide *guide;
@property (nonatomic, assign) NSInteger buildId;
@property (nonatomic, assign) NSInteger updateId;
@end

@implementation MochiNode

- (id)initWithGoValue:(MochiGoValue *)value {
    if (self = [super init]) {
        self.goValue = value;
        self.buildId = value[@"BuildId"].toLongLong;
        self.updateId = value[@"UpdateId"].toLongLong;
    }
    return self;
}

- (MochiPaintOptions *)paintOptions {
    return [[MochiPaintOptions alloc] initWithGoValue:self.goValue[@"PaintOptions"]];
}

- (NSMapTable<MochiGoValue *, MochiNode *> *)nodeChildren {
    if (_nodeChildren == nil) {
        NSMapTable *children = self.goValue[@"Children"].toMapTable;
        NSMapTable<MochiGoValue *, MochiNode *> *nodeChildren = [NSMapTable strongToStrongObjectsMapTable];
        for (MochiGoValue *i in children) {
            nodeChildren[i] = [[MochiNode alloc] initWithGoValue:children[i]];
        }
        _nodeChildren = nodeChildren;
    }
    return _nodeChildren;
}

- (MochiLayoutGuide *)guide {
    if (_guide == nil) {
        _guide = [[MochiLayoutGuide alloc] initWithGoValue:self.goValue[@"LayoutGuide"]];
    }
    return _guide;
}

- (NSString *)bridgeName {
    return self.goValue[@"Bridge"][@"Name"].toString;
}

- (MochiGoValue *)bridgeState {
    return self.goValue[@"Bridge"][@"State"];
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
    if (self = [super init]) {
        self.frame = value[@"Frame"].toCGRect;
        self.insets = value[@"Insets"].toUIEdgeInsets;
        self.zIndex = value[@"ZIndex"].toLongLong;
    }
    return self;
}

@end
