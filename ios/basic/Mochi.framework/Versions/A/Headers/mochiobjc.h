#ifndef MOCHIOBJC_H
#define MOCHIOBJC_H

#import <Foundation/Foundation.h>

int MochiTest();

typedef int64_t ObjcRef;
typedef int64_t GoRef;

typedef struct CGoBuffer {
    void *ptr; // UTF8 encoded string
    int64_t len; // length in bytes
} CGoBuffer;

@interface MochiObjcBridge : NSObject
+ (MochiObjcBridge *)sharedBridge;
@end

ObjcRef MochiObjcBridge_();

ObjcRef MochiObjcBool(bool v);
bool MochiObjcToBool(ObjcRef v);
ObjcRef MochiObjcInt64(int64_t v);
int64_t MochiObjcToInt64(ObjcRef v);
ObjcRef MochiObjcFloat64(double v);
double MochiObjcToFloat64(ObjcRef v);
ObjcRef MochiObjcGoRef(GoRef v);
GoRef MochiObjcToGoRef(ObjcRef v);
ObjcRef MochiObjcString(CGoBuffer str); // Frees the buffer
CGoBuffer MochiObjcToString(ObjcRef v);
ObjcRef MochiObjcBytes(CGoBuffer bytes); // Frees the buffer
CGoBuffer MochiObjcToBytes(ObjcRef v);

ObjcRef MochiObjcArray();
int64_t MochiObjcArrayLen(ObjcRef v);
void MochiObjcArrayAppend(ObjcRef v, ObjcRef a);
ObjcRef MochiObjcArrayAt(ObjcRef v, int64_t index);

// ObjcRef MochiObjcDict();
// ObjcRef MochiObjcDictKeys(ObjcRef v);
// ObjcRef MochiObjcDictGet(ObjcRef v, ObjcRef key);
// ObjcRef MochiObjcDictSet(ObjcRef v, ObjcRef key, ObjCRef value);

// Call
ObjcRef MochiObjcCallSentinel();
ObjcRef MochiObjcCall(ObjcRef v, CGoBuffer str, ObjcRef args);

// Tracker
ObjcRef MochiTrackObjc(id value);
id MochiGetObjc(ObjcRef key);
void MochiUntrackObjc(ObjcRef key);

// Utilities
NSString *MochiCGoBufferToNSString(CGoBuffer buf); // Frees the buffer.
CGoBuffer MochiNSStringToCGoBuffer(NSString *str); // Allocates a buffer.
NSData *MochiCGoBufferToNSData(CGoBuffer buf); // Frees the buffer.
CGoBuffer MochiNSDataToCGoBuffer(NSData *data); // Allocates a buffer.


// ObjcRef MochiObjcWithGo(GoRef v);
// GoRef MochiObjcToGo(ObjcRef v);

#endif //MOCHIOBJC_H