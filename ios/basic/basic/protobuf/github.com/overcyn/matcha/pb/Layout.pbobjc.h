// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: github.com/overcyn/matcha/pb/layout.proto

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

@class MatchaPBInsets;
@class MatchaPBPoint;
@class MatchaPBRect;

NS_ASSUME_NONNULL_BEGIN

#pragma mark - MatchaPBLayoutRoot

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
@interface MatchaPBLayoutRoot : GPBRootObject
@end

#pragma mark - MatchaPBPoint

typedef GPB_ENUM(MatchaPBPoint_FieldNumber) {
  MatchaPBPoint_FieldNumber_X = 1,
  MatchaPBPoint_FieldNumber_Y = 2,
};

@interface MatchaPBPoint : GPBMessage

@property(nonatomic, readwrite) double x;

@property(nonatomic, readwrite) double y;

@end

#pragma mark - MatchaPBRect

typedef GPB_ENUM(MatchaPBRect_FieldNumber) {
  MatchaPBRect_FieldNumber_Min = 1,
  MatchaPBRect_FieldNumber_Max = 2,
};

@interface MatchaPBRect : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) MatchaPBPoint *min;
/** Test to see if @c min has been set. */
@property(nonatomic, readwrite) BOOL hasMin;

@property(nonatomic, readwrite, strong, null_resettable) MatchaPBPoint *max;
/** Test to see if @c max has been set. */
@property(nonatomic, readwrite) BOOL hasMax;

@end

#pragma mark - MatchaPBInsets

typedef GPB_ENUM(MatchaPBInsets_FieldNumber) {
  MatchaPBInsets_FieldNumber_Top = 1,
  MatchaPBInsets_FieldNumber_Left = 2,
  MatchaPBInsets_FieldNumber_Bottom = 3,
  MatchaPBInsets_FieldNumber_Right = 4,
};

@interface MatchaPBInsets : GPBMessage

@property(nonatomic, readwrite) double top;

@property(nonatomic, readwrite) double left;

@property(nonatomic, readwrite) double bottom;

@property(nonatomic, readwrite) double right;

@end

#pragma mark - MatchaPBGuide

typedef GPB_ENUM(MatchaPBGuide_FieldNumber) {
  MatchaPBGuide_FieldNumber_Frame = 1,
  MatchaPBGuide_FieldNumber_Insets = 2,
  MatchaPBGuide_FieldNumber_ZIndex = 3,
};

@interface MatchaPBGuide : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) MatchaPBRect *frame;
/** Test to see if @c frame has been set. */
@property(nonatomic, readwrite) BOOL hasFrame;

@property(nonatomic, readwrite, strong, null_resettable) MatchaPBInsets *insets;
/** Test to see if @c insets has been set. */
@property(nonatomic, readwrite) BOOL hasInsets;

@property(nonatomic, readwrite) int64_t zIndex;

@end

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)