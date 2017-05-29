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

@interface MochiViewRoot ()
@end

@implementation MochiViewRoot
- (id)initWithGoValue:(MochiGoValue *)value {
    if ((self = [super init])) {
        self.value = value;
    }
    return self;
}

- (NSArray<MochiGoValue *> *)call:(int64_t)funcId viewId:(int64_t)viewId args:(NSArray<MochiGoValue *> *)args {
    MochiGoValue *goValue = [[MochiGoValue alloc] initWithLongLong:funcId];
    MochiGoValue *goViewId = [[MochiGoValue alloc] initWithLongLong:viewId];
    MochiGoValue *goArgs = [[MochiGoValue alloc] initWithArray:args];
    return [self.value call:@"Call" args:@[goValue, goViewId, goArgs]];
}

@end

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

@end
