// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: gomatcha.io/matcha/pb/view/view.proto

// This CPP symbol can be defined to use imports that match up to the framework
// imports needed when using CocoaPods.
#if !defined(GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS)
 #define GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS 0
#endif

#if GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS
 #import <Protobuf/GPBProtocolBuffers.h>
#else
 #import "GPBProtocolBuffers.h"
#endif

#if GOOGLE_PROTOBUF_OBJC_VERSION < 30002
#error This file was generated by a newer version of protoc which is incompatible with your Protocol Buffer library sources.
#endif
#if 30002 < GOOGLE_PROTOBUF_OBJC_MIN_SUPPORTED_VERSION
#error This file was generated by an older version of protoc which is incompatible with your Protocol Buffer library sources.
#endif

// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

CF_EXTERN_C_BEGIN

@class GPBAny;
@class MatchaPaintPBStyle;
@class MatchaViewPBBuildNode;
@class MatchaViewPBLayoutPaintNode;

NS_ASSUME_NONNULL_BEGIN

#pragma mark - MatchaViewPBViewRoot

/**
 * Exposes the extension registry for this file.
 *
 * The base class provides:
 * @code
 *   + (GPBExtensionRegistry *)extensionRegistry;
 * @endcode
 * which is a @c GPBExtensionRegistry that includes all the extensions defined by
 * this file and all files that it depends on.
 **/
@interface MatchaViewPBViewRoot : GPBRootObject
@end

#pragma mark - MatchaViewPBBuildNode

typedef GPB_ENUM(MatchaViewPBBuildNode_FieldNumber) {
  MatchaViewPBBuildNode_FieldNumber_Id_p = 1,
  MatchaViewPBBuildNode_FieldNumber_BuildId = 2,
  MatchaViewPBBuildNode_FieldNumber_BridgeName = 3,
  MatchaViewPBBuildNode_FieldNumber_BridgeValue = 4,
  MatchaViewPBBuildNode_FieldNumber_Values = 5,
  MatchaViewPBBuildNode_FieldNumber_ChildrenArray = 6,
};

@interface MatchaViewPBBuildNode : GPBMessage

@property(nonatomic, readwrite) int64_t id_p;

@property(nonatomic, readwrite) int64_t buildId;

@property(nonatomic, readwrite, copy, null_resettable) NSString *bridgeName;

@property(nonatomic, readwrite, strong, null_resettable) GPBAny *bridgeValue;
/** Test to see if @c bridgeValue has been set. */
@property(nonatomic, readwrite) BOOL hasBridgeValue;

@property(nonatomic, readwrite, strong, null_resettable) NSMutableDictionary<NSString*, GPBAny*> *values;
/** The number of items in @c values without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger values_Count;

@property(nonatomic, readwrite, strong, null_resettable) GPBInt64Array *childrenArray;
/** The number of items in @c childrenArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger childrenArray_Count;

@end

#pragma mark - MatchaViewPBLayoutPaintNode

typedef GPB_ENUM(MatchaViewPBLayoutPaintNode_FieldNumber) {
  MatchaViewPBLayoutPaintNode_FieldNumber_Id_p = 1,
  MatchaViewPBLayoutPaintNode_FieldNumber_LayoutId = 2,
  MatchaViewPBLayoutPaintNode_FieldNumber_PaintId = 3,
  MatchaViewPBLayoutPaintNode_FieldNumber_Minx = 4,
  MatchaViewPBLayoutPaintNode_FieldNumber_Miny = 5,
  MatchaViewPBLayoutPaintNode_FieldNumber_Maxx = 6,
  MatchaViewPBLayoutPaintNode_FieldNumber_Maxy = 7,
  MatchaViewPBLayoutPaintNode_FieldNumber_ZIndex = 8,
  MatchaViewPBLayoutPaintNode_FieldNumber_ChildOrderArray = 9,
  MatchaViewPBLayoutPaintNode_FieldNumber_PaintStyle = 10,
};

@interface MatchaViewPBLayoutPaintNode : GPBMessage

@property(nonatomic, readwrite) int64_t id_p;

@property(nonatomic, readwrite) int64_t layoutId;

@property(nonatomic, readwrite) int64_t paintId;

/**
 * matcha.layout.Guide layoutGuide = 4;
 * Guide
 **/
@property(nonatomic, readwrite) double minx;

@property(nonatomic, readwrite) double miny;

@property(nonatomic, readwrite) double maxx;

@property(nonatomic, readwrite) double maxy;

@property(nonatomic, readwrite) int64_t zIndex;

@property(nonatomic, readwrite, strong, null_resettable) GPBInt64Array *childOrderArray;
/** The number of items in @c childOrderArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger childOrderArray_Count;

/**
 * PaintStyle
 * double transparency = 1;
 * matcha.Color backgroundColor = 2;
 * matcha.Color borderColor = 3;
 * double borderWidth = 4;
 * double cornerRadius = 5;
 * double shadowRadius = 7;
 * matcha.layout.Point shadowOffset = 8;
 * matcha.Color shadowColor = 9;
 **/
@property(nonatomic, readwrite, strong, null_resettable) MatchaPaintPBStyle *paintStyle;
/** Test to see if @c paintStyle has been set. */
@property(nonatomic, readwrite) BOOL hasPaintStyle;

@end

#pragma mark - MatchaViewPBRoot

typedef GPB_ENUM(MatchaViewPBRoot_FieldNumber) {
  MatchaViewPBRoot_FieldNumber_LayoutPaintNodes = 2,
  MatchaViewPBRoot_FieldNumber_BuildNodes = 3,
};

@interface MatchaViewPBRoot : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) GPBInt64ObjectDictionary<MatchaViewPBLayoutPaintNode*> *layoutPaintNodes;
/** The number of items in @c layoutPaintNodes without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger layoutPaintNodes_Count;

@property(nonatomic, readwrite, strong, null_resettable) GPBInt64ObjectDictionary<MatchaViewPBBuildNode*> *buildNodes;
/** The number of items in @c buildNodes without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger buildNodes_Count;

@end

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
