#import "MatchaNode.h"
#import "MatchaProtobuf.h"

@interface MatchaNodeRoot ()
@property (nonatomic, strong) MatchaNode *node;
@end

@implementation MatchaNodeRoot
- (id)initWithProtobuf:(MatchaViewPBRoot *)pbroot {
    if ((self = [super init])) {
        if (pbroot.node) {
            self.node = [[MatchaNode alloc] initWithProtobuf:pbroot.node];
        }
    }
    return self;
}
@end

@interface MatchaNode ()
@property (nonatomic, strong) NSDictionary *nodeChildren;
@property (nonatomic, strong) MatchaGoValue *goValue;
@property (nonatomic, strong) MatchaLayoutGuide *guide;
@property (nonatomic, strong) MatchaPaintOptions *paintOptions;
@property (nonatomic, strong) NSNumber *identifier;
@property (nonatomic, strong) NSNumber *buildId;
@property (nonatomic, strong) NSNumber *layoutId;
@property (nonatomic, strong) NSNumber *paintId;
@property (nonatomic, strong) NSString *nativeViewName;
@property (nonatomic, strong) GPBAny *nativeViewState;
@property (nonatomic, strong) NSMutableDictionary<NSString*, GPBAny*> *nativeValues;
@property (nonatomic, strong) NSMutableDictionary<NSNumber *, GPBAny *> *touchRecognizers;
@end

@implementation MatchaNode

- (id)initWithProtobuf:(MatchaViewPBNode *)node {
    if ((self = [super init])) {
        self.identifier = @(node.id_p);
        self.buildId = @(node.buildId);
        self.layoutId = @(node.layoutId);
        self.paintId = @(node.paintId);
        self.paintOptions = [[MatchaPaintOptions alloc] initWithProtobuf:node.paintStyle];
        self.guide = [[MatchaLayoutGuide alloc] initWithProtobuf:node.layoutGuide];
        self.nativeViewName = node.bridgeName;
        self.nativeViewState = node.bridgeValue;
        self.nativeValues = node.values;
        
        NSMutableDictionary *children = [NSMutableDictionary dictionary];
        for (MatchaViewPBNode *i in node.childrenArray) {
            MatchaNode *child = [[MatchaNode alloc] initWithProtobuf:i];
            children[child.identifier] = child;
        }
        self.nodeChildren = children;
        
        GPBAny *any = self.nativeValues[@"gomatcha.io/matcha/touch"];
        NSError *error = nil;
        MatchaPBTouchRecognizerList *recognizerList = (id)[any unpackMessageClass:[MatchaPBTouchRecognizerList class] error:&error];
        if (error == nil) {
            NSMutableDictionary *touchRecognizers = [NSMutableDictionary dictionary];
            for (MatchaPBTouchRecognizer *i in recognizerList.recognizersArray) {
                touchRecognizers[@(i.id_p)] = i.recognizer;
            }
            self.touchRecognizers = touchRecognizers;
        }
    }
    return self;
}

@end

@interface MatchaPaintOptions ()
@property (nonatomic, assign) CGFloat transparency;
@property (nonatomic, strong) UIColor *backgroundColor;
@property (nonatomic, strong) UIColor *borderColor;
@property (nonatomic, assign) CGFloat borderWidth;
@property (nonatomic, assign) CGFloat cornerRadius;
@property (nonatomic, assign) CGFloat shadowRadius;
@property (nonatomic, assign) CGSize shadowOffset;
@property (nonatomic, strong) UIColor *shadowColor;
@end

@implementation MatchaPaintOptions

- (id)initWithProtobuf:(MatchaPaintPBStyle *)style {
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


@interface MatchaLayoutGuide ()
@property (nonatomic, assign) CGRect frame;
@property (nonatomic, assign) UIEdgeInsets insets;
@property (nonatomic, assign) NSInteger zIndex;
@end

@implementation MatchaLayoutGuide

- (id)initWithProtobuf:(MatchaLayoutPBGuide *)guide {
    if (self = [super init]) {
        self.frame = guide.frame.toCGRect;
        self.insets = guide.insets.toUIEdgeInsets;
        self.zIndex = guide.zIndex;
    }
    return self;
}

@end
