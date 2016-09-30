// Objective-C API for talking to JimMobileSDK/jimsdk Go package.
//   gobind -lang=objc -prefix=Go JimMobileSDK/jimsdk
//
// File is generated by gobind. Do not edit.

#ifndef __GoJimsdk_H__
#define __GoJimsdk_H__

#include <Foundation/Foundation.h>
#include "GoUniverse.h"

@class GoJimsdkClient;
@class GoJimsdkResponseData;
@protocol GoJimsdkResponseListener;
@class GoJimsdkResponseListener;

@interface GoJimsdkClient : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (NSString*)clusterURL;
- (void)setClusterURL:(NSString*)v;
- (long)appID;
- (void)setAppID:(long)v;
- (NSString*)jimAppID;
- (void)setJimAppID:(NSString*)v;
- (NSString*)jimAppSecret;
- (void)setJimAppSecret:(NSString*)v;
- (int64_t)serverTimestampDiff;
- (void)setServerTimestampDiff:(int64_t)v;
- (GoJimsdkResponseData*)sendVerifyEmail:(NSString*)email;
- (void)sendVerifyEmailAsync:(NSString*)email listener:(id<GoJimsdkResponseListener>)listener;
@end

@interface GoJimsdkResponseData : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (BOOL)result;
- (void)setResult:(BOOL)v;
@end

@protocol GoJimsdkResponseListener
- (void)onFailure:(NSString*)err;
- (void)onSuccess:(GoJimsdkResponseData*)respData;
@end

FOUNDATION_EXPORT BOOL GoJimsdkNewClient(NSString* clusterURL, long appID, NSString* jimAppID, NSString* jimAppSecret, GoJimsdkClient** ret0_, NSError** error);

@class GoJimsdkResponseListener;

@interface GoJimsdkResponseListener : NSObject <GoJimsdkResponseListener> {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (void)onFailure:(NSString*)err;
- (void)onSuccess:(GoJimsdkResponseData*)respData;
@end

#endif
