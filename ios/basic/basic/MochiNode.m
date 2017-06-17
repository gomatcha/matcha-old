//
//  MochiNode.m
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiNode.h"
#import "MochiProtobuf.h"

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
@property (nonatomic, strong) NSMutableDictionary<NSNumber *, GPBAny *> *touchRecognizers;
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
        
        GPBAny *any = self.nativeValues[@"github.com/overcyn/mochi/touch"];
        NSError *error = nil;
        MochiPBTouchRecognizerList *recognizerList = (id)[any unpackMessageClass:[MochiPBTouchRecognizerList class] error:&error];
        if (error == nil) {
            NSMutableDictionary *touchRecognizers = [NSMutableDictionary dictionary];
            for (MochiPBTouchRecognizer *i in recognizerList.recognizersArray) {
                touchRecognizers[@(i.id_p)] = i.recognizer;
            }
            self.touchRecognizers = touchRecognizers;
        }
    }
    return self;
}

@end

@interface MochiPaintOptions ()
@property (nonatomic, assign) CGFloat transparency;
@property (nonatomic, strong) UIColor *backgroundColor;
@property (nonatomic, strong) UIColor *borderColor;
@property (nonatomic, assign) CGFloat borderWidth;
@property (nonatomic, assign) CGFloat cornerRadius;
@property (nonatomic, assign) CGFloat shadowRadius;
@property (nonatomic, assign) CGSize shadowOffset;
@property (nonatomic, strong) UIColor *shadowColor;
@end

@implementation MochiPaintOptions

- (id)initWithProtobuf:(MochiPBPaintStyle *)style {
    if (self = [super init]) {
        self.transparency = style.transparency;
        self.backgroundColor = [[UIColor alloc] initWithProtobuf:style.backgroundColor];
        self.borderColor = [[UIColor alloc] initWithProtobuf:style.borderColor];
        self.borderWidth = style.borderWidth;
        self.cornerRadius = style.cornerRadius;
        self.shadowRadius = style.shadowRadius;
        self.shadowOffset = style.shadowOffset.toCGSize;
        self.shadowColor = [[UIColor alloc] initWithProtobuf:style.shadowColor];
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
