//
//  ViewController.m
//  basic
//
//  Created by Kevin Dang on 3/30/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import "ViewController.h"
#import "MochiButtonGestureRecognizer2.h"

@interface ViewController ()

@end

@implementation ViewController

- (id)initWithNibName:(NSString *)nibNameOrNil bundle:(NSBundle *)nibBundleOrNil {
    if ((self = [super initWithNibName:nil bundle:nil])) {
        self.view.backgroundColor = [UIColor redColor];
        
        UIView *subview = [[UIView alloc] initWithFrame:CGRectMake(0, 0, 100, 100)];
        subview.backgroundColor = [UIColor greenColor];
        subview.autoresizingMask = UIViewAutoresizingNone;
        [self.view addSubview:subview];
        
        MochiButtonGestureRecognizer2 *recognizer = [[MochiButtonGestureRecognizer2 alloc] initWithTarget:self action:@selector(action:)];
        [subview addGestureRecognizer:recognizer];
    }
    return self;
}

- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view.
}

- (void)didReceiveMemoryWarning {
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}

- (void)action:(MochiButtonGestureRecognizer2 *)recognizer {
    NSLog(@"action,%@,%@",recognizer,@(recognizer.state));
}

@end
