//
//  MochiNode.m
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiNode.h"
#import "MochiBridge.h"
#import "View.pbobjc.h"
#import "Layout.pbobjc.h"
#import "Text.pbobjc.h"
#import "Paint.pbobjc.h"

@interface MochiNodeRoot ()
@property (nonatomic, strong) MochiNode *node;
@end

@implementation MochiNodeRoot
- (id)initWithProtobuf:(MochiPBRoot *)pbroot {
    if ((self = [super init])) {
        if (pbroot.node) {
            self.node = [[MochiNode alloc] initWithProtobuf:pbroot.node];
        }
    }
    return self;
}
@end

@interface MochiNode ()
@property (nonatomic, strong) NSDictionary *nodeChildren;
@property (nonatomic, strong) MochiGoValue *goValue;
@property (nonatomic, strong) MochiLayoutGuide *guide;
@property (nonatomic, strong) MochiPaintOptions *paintOptions;
@property (nonatomic, strong) NSNumber *identifier;
@property (nonatomic, strong) NSNumber *buildId;
@property (nonatomic, strong) NSNumber *layoutId;
@property (nonatomic, strong) NSNumber *paintId;
@property (nonatomic, strong) NSString *nativeViewName;
@property (nonatomic, strong) GPBAny *nativeViewState;
@property (nonatomic, strong) NSMutableDictionary<NSString*, GPBAny*> *nativeValues;
@end

@implementation MochiNode

- (id)initWithProtobuf:(MochiPBNode *)node {
    if ((self = [super init])) {
        self.identifier = @(node.id_p);
        self.buildId = @(node.buildId);
        self.layoutId = @(node.layoutId);
        self.paintId = @(node.paintId);
        self.paintOptions = [[MochiPaintOptions alloc] initWithProtobuf:node.paintStyle];
        self.guide = [[MochiLayoutGuide alloc] initWithProtobuf:node.layoutGuide];
        self.nativeViewName = node.bridgeName;
        self.nativeViewState = node.bridgeValue;
        self.nativeValues = node.values;
        
        NSMutableDictionary *children = [NSMutableDictionary dictionary];
        for (MochiPBNode *i in node.childrenArray) {
            MochiNode *child = [[MochiNode alloc] initWithProtobuf:i];
            children[child.identifier] = child;
        }
        self.nodeChildren = children;
    }
    return self;
}

- (id)initWithGoValue:(MochiGoValue *)value {
    if ((self = [super init])) {
        self.goValue = value;
        self.identifier = @(value[@"Id"].toLongLong);
        self.buildId = @(value[@"BuildId"].toLongLong);
        self.layoutId = @(value[@"LayoutId"].toLongLong);
        self.paintId = @(value[@"PaintId"].toLongLong);
        self.paintOptions = [[MochiPaintOptions alloc] initWithGoValue:self.goValue[@"PaintOptions"]];
        self.guide = [[MochiLayoutGuide alloc] initWithGoValue:self.goValue[@"LayoutGuide"]];
        self.nativeViewName = self.goValue[@"BridgeName"].toString;
    }
    return self;
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

// - (NSString *)description {
//     return [NSString stringWithFormat:@"<MochiNode id:%@,%@,%@,%@
// }

@end

@interface MochiPaintOptions ()
@property (nonatomic, strong) MochiGoValue *goValue;
@property (nonatomic, strong) UIColor *backgroundColor;
@end

@implementation MochiPaintOptions

- (id)initWithProtobuf:(MochiPBPaintStyle *)style {
    if (self = [super init]) {
        self.backgroundColor = [[UIColor alloc] initWithProtobuf:style.backgroundColor];
    }
    return self;
}

- (id)initWithGoValue:(MochiGoValue *)value {
    if (self = [super init]) {
        self.goValue = value;
        
        MochiGoValue *value = self.goValue[@"BackgroundColor"];
        if (!value.isNil) {
            self.backgroundColor = [[UIColor alloc] initWithGoValue:value];
        }
    }
    return self;
}

@end


@interface MochiLayoutGuide ()
@property (nonatomic, assign) CGRect frame;
@property (nonatomic, assign) UIEdgeInsets insets;
@property (nonatomic, assign) NSInteger zIndex;
@end

@implementation MochiLayoutGuide

- (id)initWithProtobuf:(MochiPBGuide *)guide {
    if (self = [super init]) {
        self.frame = guide.frame.toCGRect;
        self.insets = guide.insets.toUIEdgeInsets;
        self.zIndex = guide.zIndex;
    }
    return self;
}

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
