// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: github.com/overcyn/matcha/pb/view/textinput/textinput.proto

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

@class MatchaPBStyledText;

NS_ASSUME_NONNULL_BEGIN

#pragma mark - MatchaTextInputPBTextinputRoot

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
@interface MatchaTextInputPBTextinputRoot : GPBRootObject
@end

#pragma mark - MatchaTextInputPBView

typedef GPB_ENUM(MatchaTextInputPBView_FieldNumber) {
  MatchaTextInputPBView_FieldNumber_StyledText = 1,
  MatchaTextInputPBView_FieldNumber_OnUpdate = 2,
  MatchaTextInputPBView_FieldNumber_OnFocus = 3,
  MatchaTextInputPBView_FieldNumber_Focused = 4,
};

@interface MatchaTextInputPBView : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) MatchaPBStyledText *styledText;
/** Test to see if @c styledText has been set. */
@property(nonatomic, readwrite) BOOL hasStyledText;

@property(nonatomic, readwrite) BOOL focused;

@property(nonatomic, readwrite) int64_t onUpdate;

@property(nonatomic, readwrite) int64_t onFocus;

@end

#pragma mark - MatchaTextInputPBEvent

typedef GPB_ENUM(MatchaTextInputPBEvent_FieldNumber) {
  MatchaTextInputPBEvent_FieldNumber_StyledText = 1,
};

@interface MatchaTextInputPBEvent : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) MatchaPBStyledText *styledText;
/** Test to see if @c styledText has been set. */
@property(nonatomic, readwrite) BOOL hasStyledText;

@end

#pragma mark - MatchaTextInputPBFocusEvent

typedef GPB_ENUM(MatchaTextInputPBFocusEvent_FieldNumber) {
  MatchaTextInputPBFocusEvent_FieldNumber_Focused = 1,
};

@interface MatchaTextInputPBFocusEvent : GPBMessage

@property(nonatomic, readwrite) BOOL focused;

@end

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
