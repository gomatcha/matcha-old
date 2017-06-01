//
//  MochiViewController.m
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "MochiViewController.h"
#import "MochiView.h"
#import "MochiBridge.h"
#import "MochiNode.h"

@interface MochiViewController ()
@property (nonatomic, assign) NSInteger identifier;
@property (nonatomic, strong) MochiBasicView *mochiView;
@property (nonatomic, strong) MochiGoValue *goValue;
@property (nonatomic, assign) CGRect lastFrame;
@end

@implementation MochiViewController

+ (NSPointerArray *)viewControllers {
    static NSPointerArray *sPointerArray;
    static dispatch_once_t sOnce;
    dispatch_once(&sOnce, ^{
        sPointerArray = [NSPointerArray weakObjectsPointerArray];
    });
    return sPointerArray;
}

+ (MochiViewController *)viewControllerWithIdentifier:(NSInteger)identifier {
    for (MochiViewController *i in [self viewControllers]) {
        if (i.identifier == identifier) {
            return i;
        }
    }
    return nil;
}
- (id)initWithGoValue:(MochiGoValue *)value {
    if ((self = [super initWithNibName:nil bundle:nil])) {
        self.goValue = value;
        self.identifier = [value call:@"Id" args:nil][0].toLongLong;
        [[MochiViewController viewControllers] addPointer:(__bridge void *)self];
    }
    return self;
}

- (void)loadView {
    self.mochiView = [[MochiBasicView alloc] initWithViewRoot:self parentVC:self];
    self.view = self.mochiView;
}

- (void)viewDidLayoutSubviews {
    if (!CGRectEqualToRect(self.lastFrame, self.view.frame)) {
        self.lastFrame = self.view.frame;
        
        [self.goValue call:@"SetSize" args:@[[[MochiGoValue alloc] initWithCGPoint:CGPointMake(self.view.frame.size.width, self.view.frame.size.height)]]];
    }
}

- (NSArray<MochiGoValue *> *)call:(int64_t)funcId viewId:(int64_t)viewId args:(NSArray<MochiGoValue *> *)args {
    MochiGoValue *goValue = [[MochiGoValue alloc] initWithLongLong:funcId];
    MochiGoValue *goViewId = [[MochiGoValue alloc] initWithLongLong:viewId];
    MochiGoValue *goArgs = [[MochiGoValue alloc] initWithArray:args];
    return [self.goValue call:@"Call" args:@[goValue, goViewId, goArgs]];
}


- (void)update:(MochiNode *)node {
    self.mochiView.node = node;
}

@end
