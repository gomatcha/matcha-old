#import <UIKit/UIKit.h>
#import <MatchaBridge/MatchaBridge.h>
@class MatchaPaintOptions;
@class MatchaLayoutGuide;
@class MatchaNode;
@class MatchaViewPBNode;
@class MatchaViewPBRoot;
@class MatchaLayoutPBGuide;
@class MatchaPaintPBStyle;
@class MatchaPBRecognizer;
@class MatchaViewPBLayoutPaintNode;
@class MatchaLayoutPaintNode;
@class GPBInt64ObjectDictionary;
@class MatchaViewPBLayoutPaintNode;
@class GPBAny;

@interface MatchaNodeRoot : NSObject // view.root
- (id)initWithProtobuf:(MatchaViewPBRoot *)data;
@property (nonatomic, readonly) MatchaNode *node;
@property (nonatomic, readonly) GPBInt64ObjectDictionary *layoutPaintNodes;
@end

@interface MatchaNode : NSObject // view.node
- (id)initWithProtobuf:(MatchaViewPBNode *)node;
@property (nonatomic, readonly) NSDictionary<NSNumber *, MatchaNode *> *nodeChildren;
@property (nonatomic, readonly) NSMutableDictionary<NSString*, GPBAny*> *nativeValues;
@property (nonatomic, readonly) NSString *nativeViewName;
@property (nonatomic, readonly) GPBAny *nativeViewState;
@property (nonatomic, readonly) NSNumber *identifier;
@property (nonatomic, readonly) NSNumber *buildId;
@property (nonatomic, readonly) NSDictionary<NSNumber *, GPBAny *> *touchRecognizers;
@end

@interface MatchaLayoutPaintNode : NSObject
- (id)initWithProtobuf:(MatchaViewPBLayoutPaintNode *)node;
@property (nonatomic, readonly) NSNumber *identifier;
@property (nonatomic, readonly) NSNumber *layoutId;
@property (nonatomic, readonly) NSNumber *paintId;
@property (nonatomic, readonly) MatchaLayoutGuide *guide;
@property (nonatomic, readonly) MatchaPaintOptions *paintOptions;
@end

@interface MatchaPaintOptions : NSObject
- (id)initWithProtobuf:(MatchaPaintPBStyle *)style;
@property (nonatomic, readonly) CGFloat transparency;
@property (nonatomic, readonly) UIColor *backgroundColor;
@property (nonatomic, readonly) UIColor *borderColor;
@property (nonatomic, readonly) CGFloat borderWidth;
@property (nonatomic, readonly) CGFloat cornerRadius;
@property (nonatomic, readonly) CGFloat shadowRadius;
@property (nonatomic, readonly) CGSize shadowOffset;
@property (nonatomic, readonly) UIColor *shadowColor;
@end

@interface MatchaLayoutGuide : NSObject
- (id)initWithProtobuf:(MatchaLayoutPBGuide *)guide;
@property (nonatomic, readonly) CGRect frame;
@property (nonatomic, readonly) UIEdgeInsets insets;
@property (nonatomic, readonly) NSInteger zIndex;
@end
