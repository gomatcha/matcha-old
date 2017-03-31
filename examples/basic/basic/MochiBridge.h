//
//  MochiBridge.h
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import <UIKit/UIKit.h>
#import "Bridge+Extensions.h"

@interface UIColor (Mochi)
- (id)initWithBridgeValue:(BridgeValue *)value;
@end
