//
//  MochiButton.m
//  basic
//
//  Created by Kevin Dang on 6/17/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiButton.h"
#import "MochiViewController.h"

@interface MochiButton ()
@property (nonatomic, strong) UIButton *button;
@property (nonatomic, weak) MochiViewNode *viewNode;
@property (nonatomic, strong) MochiNode *node;
@property (nonatomic, assign) int64_t funcId;
@end

@implementation MochiButton

- (id)initWithViewNode:(MochiViewNode *)viewNode {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewNode = viewNode;
        self.button = [UIButton buttonWithType:UIButtonTypeSystem];
        [self.button addTarget:self action:@selector(onPress) forControlEvents:UIControlEventTouchUpInside];
        [self addSubview:self.button];
    }
    return self;
}

- (void)setNode:(MochiNode *)value {
    _node = value;
    GPBAny *state = value.nativeViewState;
    NSError *error = nil;
    MochiPBButtonButton *pbbutton = (id)[state unpackMessageClass:[MochiPBButtonButton class] error:&error];
    
    NSAttributedString *string = [[NSAttributedString alloc] initWithProtobuf:pbbutton.styledText];
    [self.button setAttributedTitle:string forState:UIControlStateNormal];
    self.funcId = pbbutton.onPress;
}

- (void)layoutSubviews {
    self.button.frame = self.bounds;
}

- (void)onPress {
    [self.viewNode.rootVC call:self.funcId viewId:self.node.identifier.longLongValue args:@[]];
}

@end
