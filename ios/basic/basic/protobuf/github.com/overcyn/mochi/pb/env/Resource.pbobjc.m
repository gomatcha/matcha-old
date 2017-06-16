// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: github.com/overcyn/mochi/pb/env/resource.proto

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

 #import "github.com/overcyn/mochi/pb/env/Resource.pbobjc.h"
// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

#pragma mark - MochiPBEnvResourceRoot

@implementation MochiPBEnvResourceRoot

// No extensions in the file and no imports, so no need to generate
// +extensionRegistry.

@end

#pragma mark - MochiPBEnvResourceRoot_FileDescriptor

static GPBFileDescriptor *MochiPBEnvResourceRoot_FileDescriptor(void) {
  // This is called by +initialize so there is no need to worry
  // about thread safety of the singleton.
  static GPBFileDescriptor *descriptor = NULL;
  if (!descriptor) {
    GPB_DEBUG_CHECK_RUNTIME_VERSIONS();
    descriptor = [[GPBFileDescriptor alloc] initWithPackage:@"env"
                                                 objcPrefix:@"MochiPBEnv"
                                                     syntax:GPBFileSyntaxProto3];
  }
  return descriptor;
}

#pragma mark - MochiPBEnvResource

@implementation MochiPBEnvResource

@dynamic path;

typedef struct MochiPBEnvResource__storage_ {
  uint32_t _has_storage_[1];
  NSString *path;
} MochiPBEnvResource__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "path",
        .dataTypeSpecific.className = NULL,
        .number = MochiPBEnvResource_FieldNumber_Path,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MochiPBEnvResource__storage_, path),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MochiPBEnvResource class]
                                     rootClass:[MochiPBEnvResourceRoot class]
                                          file:MochiPBEnvResourceRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MochiPBEnvResource__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end


#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)