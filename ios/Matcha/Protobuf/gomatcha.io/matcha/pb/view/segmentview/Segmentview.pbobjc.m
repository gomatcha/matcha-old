// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: gomatcha.io/matcha/pb/view/segmentview/segmentview.proto

// This CPP symbol can be defined to use imports that match up to the framework
// imports needed when using CocoaPods.
#if !defined(GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS)
 #define GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS 0
#endif

#if GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS
 #import <Protobuf/GPBProtocolBuffers_RuntimeSupport.h>
#else
 #import "GPBProtocolBuffers_RuntimeSupport.h"
#endif

 #import "gomatcha.io/matcha/pb/view/segmentview/Segmentview.pbobjc.h"
// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

#pragma mark - MatchaSegmentViewPbSegmentviewRoot

@implementation MatchaSegmentViewPbSegmentviewRoot

// No extensions in the file and no imports, so no need to generate
// +extensionRegistry.

@end

#pragma mark - MatchaSegmentViewPbSegmentviewRoot_FileDescriptor

static GPBFileDescriptor *MatchaSegmentViewPbSegmentviewRoot_FileDescriptor(void) {
  // This is called by +initialize so there is no need to worry
  // about thread safety of the singleton.
  static GPBFileDescriptor *descriptor = NULL;
  if (!descriptor) {
    GPB_DEBUG_CHECK_RUNTIME_VERSIONS();
    descriptor = [[GPBFileDescriptor alloc] initWithPackage:@"segmentview"
                                                 objcPrefix:@"MatchaSegmentViewPb"
                                                     syntax:GPBFileSyntaxProto3];
  }
  return descriptor;
}

#pragma mark - MatchaSegmentViewPbView

@implementation MatchaSegmentViewPbView

@dynamic value;
@dynamic titlesArray, titlesArray_Count;
@dynamic momentary;
@dynamic enabled;

typedef struct MatchaSegmentViewPbView__storage_ {
  uint32_t _has_storage_[1];
  NSMutableArray *titlesArray;
  int64_t value;
} MatchaSegmentViewPbView__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "value",
        .dataTypeSpecific.className = NULL,
        .number = MatchaSegmentViewPbView_FieldNumber_Value,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaSegmentViewPbView__storage_, value),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "titlesArray",
        .dataTypeSpecific.className = NULL,
        .number = MatchaSegmentViewPbView_FieldNumber_TitlesArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(MatchaSegmentViewPbView__storage_, titlesArray),
        .flags = GPBFieldRepeated,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "momentary",
        .dataTypeSpecific.className = NULL,
        .number = MatchaSegmentViewPbView_FieldNumber_Momentary,
        .hasIndex = 1,
        .offset = 2,  // Stored in _has_storage_ to save space.
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeBool,
      },
      {
        .name = "enabled",
        .dataTypeSpecific.className = NULL,
        .number = MatchaSegmentViewPbView_FieldNumber_Enabled,
        .hasIndex = 3,
        .offset = 4,  // Stored in _has_storage_ to save space.
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeBool,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaSegmentViewPbView class]
                                     rootClass:[MatchaSegmentViewPbSegmentviewRoot class]
                                          file:MatchaSegmentViewPbSegmentviewRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaSegmentViewPbView__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MatchaSegmentViewPbEvent

@implementation MatchaSegmentViewPbEvent

@dynamic value;

typedef struct MatchaSegmentViewPbEvent__storage_ {
  uint32_t _has_storage_[1];
  int64_t value;
} MatchaSegmentViewPbEvent__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "value",
        .dataTypeSpecific.className = NULL,
        .number = MatchaSegmentViewPbEvent_FieldNumber_Value,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaSegmentViewPbEvent__storage_, value),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeInt64,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaSegmentViewPbEvent class]
                                     rootClass:[MatchaSegmentViewPbSegmentviewRoot class]
                                          file:MatchaSegmentViewPbSegmentviewRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaSegmentViewPbEvent__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end


#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
