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
@property (nonatomic, strong) MochiView *mochiView;
@property (nonatomic, strong) MochiGoValue *goVC;
@property (nonatomic, strong) MochiViewRoot *viewRoot;
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

- (id)initWithMochiViewRoot:(MochiViewRoot *)root {
    if ((self = [super initWithNibName:nil bundle:nil])) {
        self.identifier = [root.value call:@"Id" args:nil][0].toLongLong;
        [[MochiViewController viewControllers] addPointer:(__bridge void *)self];
        
        self.viewRoot = root;
    }
    return self;
}

- (void)loadView {
    self.mochiView = [[MochiView alloc] initWithViewRoot:self];
    self.view = self.mochiView;
}

- (void)viewDidLayoutSubviews {
    if (!CGRectEqualToRect(self.lastFrame, self.view.frame)) {
        self.lastFrame = self.view.frame;
        
        [self.viewRoot.value call:@"SetSize" args:@[[[MochiGoValue alloc] initWithCGPoint:CGPointMake(self.view.frame.size.width, self.view.frame.size.height)]]];
    }
}

- (void)update:(MochiNode *)node {
    self.mochiView.node = node;
}

@end
