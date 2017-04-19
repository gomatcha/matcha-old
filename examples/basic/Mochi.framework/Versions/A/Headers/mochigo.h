#ifndef MOCHIGO_H
#define MOCHIGO_H

#import <Foundation/Foundation.h>
#include "mochiobjc.h"
@class MochiGoValue;

typedef int64_t GoRef;

GoRef mochiGoRoot();
GoRef mochiGoBool(bool);
bool mochiGoToBool(GoRef);
GoRef mochiGoInt64(int64_t);
int64_t mochiGoToInt64(GoRef);
GoRef mochiGoUint64(uint64_t);
uint64_t mochiGoToUint64(GoRef);
GoRef mochiGoFloat64(double);
double mochiGoToFloat64(GoRef);
GoRef mochiGoString(CGoBuffer); // Frees the buffer
CGoBuffer mochiGoToString(GoRef);
GoRef mochiGoBytes(CGoBuffer); // Frees the buffer
CGoBuffer mochiGoToBytes(GoRef);

GoRef mochiGoArray();
int64_t mochiGoArrayLen(GoRef);
GoRef mochiGoArrayAppend(GoRef, GoRef);
GoRef mochiGoArrayAt(GoRef, int64_t);

GoRef mochiGoCall(GoRef, CGoBuffer, GoRef);
GoRef mochiGoField(GoRef, CGoBuffer);

void mochiGoUntrack(GoRef);

@interface MochiGoBridge : NSObject
+ (MochiGoBridge *)sharedBridge;
@property (nonatomic, readonly) MochiGoValue *root;
@end

@interface MochiGoValue : NSObject
- (id)initWithGoRef:(GoRef)ref;
- (id)initWithBool:(BOOL)v;
- (id)initWithLongLong:(long long)v;
- (id)initWithUnsignedLongLong:(unsigned long long)v;
- (id)initWithDouble:(double)v;
- (id)initWithString:(NSString *)v;
- (id)initWithData:(NSData *)v;
- (id)initWithArray:(NSArray<MochiGoValue *> *)v;
- (BOOL)toBool;
- (long long)toLongLong;
- (unsigned long long)toUnsignedLongLong;
- (double)toDouble;
- (NSString *)toString;
- (NSData *)toData;
- (NSArray *)toArray;
// - (NSDictionary *)toDictionary;
// - (BOOL)isNil;
- (NSArray<MochiGoValue *> *)call:(NSString *)method args:(NSArray<MochiGoValue *> *)args;
- (MochiGoValue *)field:(NSString *)name;
- (MochiGoValue *)objectForKeyedSubscript:(NSString *)key;
@end

#endif // MOCHIGO_H