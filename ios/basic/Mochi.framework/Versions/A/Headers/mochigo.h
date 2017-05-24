#ifndef MOCHIGO_H
#define MOCHIGO_H

#import <Foundation/Foundation.h>
#include "mochiobjc.h"
@class MochiGoValue;

GoRef mochiGoRoot();

GoRef mochiGoBool(bool);
bool mochiGoToBool(GoRef);
GoRef mochiGoInt(int);
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

GoRef mochiGoMap();
GoRef mochiGoMapKeys(GoRef);
GoRef mochiGoMapGet(GoRef map, GoRef key);
void mochiGoMapSet(GoRef map, GoRef key, GoRef value);

GoRef mochiGoType(CGoBuffer); // Frees the buffer
GoRef mochiGoFunc(CGoBuffer); // Frees the buffer

bool mochiGoIsNil(GoRef);
bool mochiGoEqual(GoRef, GoRef);
GoRef mochiGoElem(GoRef);
GoRef mochiGoCall(GoRef, CGoBuffer, GoRef);
GoRef mochiGoField(GoRef, CGoBuffer);
void mochiGoFieldSet(GoRef, CGoBuffer, GoRef);

void mochiGoUntrack(GoRef);

@interface MochiGoBridge : NSObject
+ (MochiGoBridge *)sharedBridge;
@property (nonatomic, readonly) MochiGoValue *root;
@end

@interface MochiGoValue : NSObject
- (id)initWithGoRef:(GoRef)ref;
- (id)initWithBool:(BOOL)v;
- (id)initWithInt:(int)v;
- (id)initWithLongLong:(long long)v;
- (id)initWithUnsignedLongLong:(unsigned long long)v;
- (id)initWithDouble:(double)v;
- (id)initWithString:(NSString *)v;
- (id)initWithData:(NSData *)v;
- (id)initWithArray:(NSArray<MochiGoValue *> *)v;
- (id)initWithType:(NSString *)typeName;
- (id)initWithFunc:(NSString *)funcName;
@property (nonatomic, readonly) GoRef ref;
- (BOOL)toBool;
- (long long)toLongLong;
- (unsigned long long)toUnsignedLongLong;
- (double)toDouble;
- (NSString *)toString;
- (NSData *)toData;
- (NSArray *)toArray;
- (NSMapTable *)toMapTable;
// - (NSDictionary *)toDictionary;
- (BOOL)isNil;
- (BOOL)isEqual:(MochiGoValue *)value;
- (MochiGoValue *)elem;
- (NSArray<MochiGoValue *> *)call:(NSString *)method args:(NSArray<MochiGoValue *> *)args; // pass in nil for the method to call a closure.
- (MochiGoValue *)field:(NSString *)name;
- (void)setField:(NSString *)name value:(MochiGoValue *)value;
- (MochiGoValue *)objectForKeyedSubscript:(NSString *)key;
- (void)setObject:(MochiGoValue *)object forKeyedSubscript:(NSString *)key;
@end

#endif // MOCHIGO_H