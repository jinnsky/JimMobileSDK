// Objective-C API for talking to JimMobileSDK/jimsdk Go package.
//   gobind -lang=objc -prefix=Go JimMobileSDK/jimsdk
//
// File is generated by gobind. Do not edit.

#ifndef __GoJimsdk_H__
#define __GoJimsdk_H__

#include <Foundation/Foundation.h>
#include "GoUniverse.h"

@class GoJimsdkAppVersionInfoResponse;
@class GoJimsdkBindEmailParams;
@class GoJimsdkBindEmailResponse;
@class GoJimsdkBindPhoneParams;
@class GoJimsdkBindPhoneResponse;
@class GoJimsdkBloodPressureCommitListParams;
@class GoJimsdkBloodPressureCommitListResponse;
@class GoJimsdkBloodPressureCommitParams;
@class GoJimsdkBloodPressureCommitResponse;
@class GoJimsdkBloodPressureCommitResponseCollection;
@class GoJimsdkBloodPressureDeleteParams;
@class GoJimsdkBloodPressureDeviceTroubleResponse;
@class GoJimsdkBloodPressureTotalCountResponse;
@class GoJimsdkChangePasswordParams;
@class GoJimsdkChangePasswordResponse;
@class GoJimsdkClient;
@class GoJimsdkFacebookUserParams;
@class GoJimsdkFacebookUserResponse;
@class GoJimsdkFaqListParams;
@class GoJimsdkFaqListResponse;
@class GoJimsdkFeedbackSubmitParams;
@class GoJimsdkFeedbackSubmitResponse;
@class GoJimsdkFirmwareOtaInfoResponse;
@class GoJimsdkHasPhoneParams;
@class GoJimsdkHasPhoneResponse;
@class GoJimsdkLinkedInUserParams;
@class GoJimsdkLinkedInUserResponse;
@class GoJimsdkLoginParams;
@class GoJimsdkNewsDigest;
@class GoJimsdkNewsDigestCollection;
@class GoJimsdkNewsDigestParams;
@class GoJimsdkNewsDigestResponse;
@class GoJimsdkQqUserParams;
@class GoJimsdkQqUserResponse;
@class GoJimsdkRegisterInfoParams;
@class GoJimsdkRegisterParams;
@class GoJimsdkResetPasswordEmailParams;
@class GoJimsdkResetPasswordEmailResponse;
@class GoJimsdkResetPasswordParams;
@class GoJimsdkResetPasswordResponse;
@class GoJimsdkResetPasswordSmsParams;
@class GoJimsdkResetPasswordSmsResponse;
@class GoJimsdkResponseError;
@class GoJimsdkSystemInfoResponse;
@class GoJimsdkTwitterUserParams;
@class GoJimsdkTwitterUserResponse;
@class GoJimsdkUpdateBindEmailParams;
@class GoJimsdkUpdateBindEmailResponse;
@class GoJimsdkUpdateBindPhoneParams;
@class GoJimsdkUpdateBindPhoneResponse;
@class GoJimsdkUpdateUserParams;
@class GoJimsdkUpdateUserResponse;
@class GoJimsdkUploadAvatarResponse;
@class GoJimsdkUserInfoParams;
@class GoJimsdkUserInfoResponse;
@class GoJimsdkVerifyEmailParams;
@class GoJimsdkVerifyEmailResponse;
@class GoJimsdkVerifySmsParams;
@class GoJimsdkVerifySmsResponse;
@class GoJimsdkWeiboUserParams;
@class GoJimsdkWeiboUserResponse;
@class GoJimsdkWeixinUserResponse;
@class GoJimsdkWexinUserParams;
@protocol GoJimsdkLoginResponseListener;
@class GoJimsdkLoginResponseListener;
@protocol GoJimsdkNewsDigestResponseListener;
@class GoJimsdkNewsDigestResponseListener;
@protocol GoJimsdkRegisterResponseListener;
@class GoJimsdkRegisterResponseListener;
@protocol GoJimsdkUserInfoResponseListener;
@class GoJimsdkUserInfoResponseListener;
@protocol GoJimsdkVerifyEmailResponseListener;
@class GoJimsdkVerifyEmailResponseListener;

@interface GoJimsdkAppVersionInfoResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (NSString*)type;
- (void)setType:(NSString*)v;
- (NSString*)downloadURL;
- (void)setDownloadURL:(NSString*)v;
- (NSString*)limit;
- (void)setLimit:(NSString*)v;
- (NSString*)current;
- (void)setCurrent:(NSString*)v;
- (NSString*)description;
- (void)setDescription:(NSString*)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkBindEmailParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)userID;
- (void)setUserID:(long)v;
- (NSString*)email;
- (void)setEmail:(NSString*)v;
- (NSString*)verificationCode;
- (void)setVerificationCode:(NSString*)v;
@end

@interface GoJimsdkBindEmailResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (BOOL)result;
- (void)setResult:(BOOL)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkBindPhoneParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)userID;
- (void)setUserID:(long)v;
- (NSString*)phone;
- (void)setPhone:(NSString*)v;
- (NSString*)verificationCode;
- (void)setVerificationCode:(NSString*)v;
@end

@interface GoJimsdkBindPhoneResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (BOOL)result;
- (void)setResult:(BOOL)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkBloodPressureCommitListParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)lastSyncID;
- (void)setLastSyncID:(long)v;
// skipped field BloodPressureCommitListParams.Items with unsupported type: *types.Slice

- (void)addParams:(GoJimsdkBloodPressureCommitParams*)params;
@end

@interface GoJimsdkBloodPressureCommitListResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (GoJimsdkBloodPressureCommitResponseCollection*)collection;
- (void)setCollection:(GoJimsdkBloodPressureCommitResponseCollection*)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkBloodPressureCommitParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (double)systolicBP;
- (void)setSystolicBP:(double)v;
- (double)diastolicBP;
- (void)setDiastolicBP:(double)v;
- (long)pulse;
- (void)setPulse:(long)v;
- (NSString*)deviceID;
- (void)setDeviceID:(NSString*)v;
- (NSString*)deviceType;
- (void)setDeviceType:(NSString*)v;
- (int64_t)time;
- (void)setTime:(int64_t)v;
@end

@interface GoJimsdkBloodPressureCommitResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (int64_t)id_;
- (void)setID:(int64_t)v;
- (int64_t)userID;
- (void)setUserID:(int64_t)v;
- (double)systolicBP;
- (void)setSystolicBP:(double)v;
- (double)diastolicBP;
- (void)setDiastolicBP:(double)v;
- (long)pulse;
- (void)setPulse:(long)v;
- (NSString*)deviceID;
- (void)setDeviceID:(NSString*)v;
- (NSString*)deviceType;
- (void)setDeviceType:(NSString*)v;
- (int64_t)time;
- (void)setTime:(int64_t)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkBloodPressureCommitResponseCollection : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
// skipped field BloodPressureCommitResponseCollection.Items with unsupported type: *types.Slice

- (GoJimsdkBloodPressureCommitResponse*)getItemAt:(long)index;
- (long)getSize;
@end

@interface GoJimsdkBloodPressureDeleteParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
// skipped field BloodPressureDeleteParams.Items with unsupported type: *types.Slice

- (void)addDeleteID:(int64_t)id;
@end

@interface GoJimsdkBloodPressureDeviceTroubleResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (int64_t)id_;
- (void)setID:(int64_t)v;
- (int64_t)appID;
- (void)setAppID:(int64_t)v;
- (int64_t)userID;
- (void)setUserID:(int64_t)v;
- (NSString*)deviceType;
- (void)setDeviceType:(NSString*)v;
- (NSString*)deviceID;
- (void)setDeviceID:(NSString*)v;
- (NSString*)troubleType;
- (void)setTroubleType:(NSString*)v;
- (int64_t)time;
- (void)setTime:(int64_t)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkBloodPressureTotalCountResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (int64_t)count;
- (void)setCount:(int64_t)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkChangePasswordParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (NSString*)oldPassword;
- (void)setOldPassword:(NSString*)v;
- (NSString*)newPassword;
- (void)setNewPassword:(NSString*)v;
@end

@interface GoJimsdkChangePasswordResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (BOOL)result;
- (void)setResult:(BOOL)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

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
- (long)requestTimeout;
- (void)setRequestTimeout:(long)v;
- (BOOL)hasValidSession;
- (GoJimsdkAppVersionInfoResponse*)sendAppVersionInfo:(NSString*)appType;
- (GoJimsdkBindEmailResponse*)sendBindEmail:(long)userID email:(NSString*)email verificationCode:(NSString*)verificationCode;
- (GoJimsdkBindPhoneResponse*)sendBindPhone:(long)userID phone:(NSString*)phone verificationCode:(NSString*)verificationCode;
- (GoJimsdkBloodPressureCommitResponse*)sendBloodPressureCommit:(GoJimsdkBloodPressureCommitParams*)params;
- (GoJimsdkBloodPressureCommitListResponse*)sendBloodPressureCommitList:(GoJimsdkBloodPressureCommitListParams*)paramsList;
- (GoJimsdkBloodPressureCommitListResponse*)sendBloodPressureDelete:(GoJimsdkBloodPressureDeleteParams*)params;
- (GoJimsdkBloodPressureDeviceTroubleResponse*)sendBloodPressureDeviceTrouble:(NSString*)deviceType deviceID:(NSString*)deviceID trouble:(NSString*)trouble time:(int64_t)time;
- (GoJimsdkBloodPressureCommitListResponse*)sendBloodPressureList:(long)index size:(long)size;
- (GoJimsdkBloodPressureCommitListResponse*)sendBloodPressureSync:(long)lastSyncId size:(long)size;
- (GoJimsdkBloodPressureTotalCountResponse*)sendBloodPressureTotalCount;
- (GoJimsdkChangePasswordResponse*)sendChangePassword:(NSString*)oldPwd newPwd:(NSString*)newPwd;
- (GoJimsdkFacebookUserResponse*)sendFacebookUser:(NSString*)accessToken;
- (GoJimsdkFaqListResponse*)sendFaqList:(GoJimsdkFaqListParams*)params;
- (GoJimsdkFeedbackSubmitResponse*)sendFeedback:(NSString*)contactInfo content:(NSString*)content;
- (GoJimsdkFirmwareOtaInfoResponse*)sendFirmwareLastOta:(NSString*)curHwVer;
- (GoJimsdkFirmwareOtaInfoResponse*)sendFirmwareOtaInfoByVersion:(NSString*)curHwVer curFwVer:(NSString*)curFwVer;
- (GoJimsdkHasPhoneResponse*)sendHasPhone:(NSString*)phone;
- (GoJimsdkLinkedInUserResponse*)sendLinkedInUser:(NSString*)accessToken;
- (GoJimsdkUserInfoResponse*)sendLogin:(GoJimsdkLoginParams*)params;
- (void)sendLoginAsync:(GoJimsdkLoginParams*)params listener:(id<GoJimsdkLoginResponseListener>)listener;
- (void)sendLogout;
- (GoJimsdkNewsDigestResponse*)sendNewsDigest:(GoJimsdkNewsDigestParams*)params;
- (void)sendNewsDigestAsync:(GoJimsdkNewsDigestParams*)params listener:(id<GoJimsdkNewsDigestResponseListener>)listener;
- (GoJimsdkQqUserResponse*)sendQqUser:(NSString*)openID;
- (GoJimsdkUserInfoResponse*)sendRegister:(GoJimsdkRegisterParams*)params;
- (void)sendRegisterAsync:(GoJimsdkRegisterParams*)params listener:(id<GoJimsdkRegisterResponseListener>)listener;
- (GoJimsdkUserInfoResponse*)sendRegisterInfo:(GoJimsdkRegisterInfoParams*)params;
- (GoJimsdkResetPasswordResponse*)sendResetPassword:(GoJimsdkResetPasswordParams*)params;
- (GoJimsdkResetPasswordEmailResponse*)sendResetPasswordEmail:(NSString*)email;
- (GoJimsdkResetPasswordSmsResponse*)sendResetPasswordSms:(NSString*)phone;
- (GoJimsdkSystemInfoResponse*)sendSystemInfo:(NSString*)model os:(NSString*)os resolution:(NSString*)resolution network:(NSString*)network;
- (GoJimsdkTwitterUserResponse*)sendTwitterUser:(NSString*)userToken tokenSecret:(NSString*)tokenSecret;
- (GoJimsdkUpdateBindEmailResponse*)sendUpdateBindEmail:(long)userID email:(NSString*)email;
- (GoJimsdkUpdateBindPhoneResponse*)sendUpdateBindPhone:(long)userID phone:(NSString*)phone;
- (GoJimsdkUpdateUserResponse*)sendUpdateUser:(GoJimsdkUpdateUserParams*)params;
- (GoJimsdkUploadAvatarResponse*)sendUploadAvatar:(NSString*)file;
- (GoJimsdkUploadAvatarResponse*)sendUploadAvatarBase64:(NSString*)encodedStr;
- (GoJimsdkUserInfoResponse*)sendUserInfo:(long)userID subUserID:(long)subUserID;
- (void)sendUserInfoAsync:(long)userID subUserID:(long)subUserID listener:(id<GoJimsdkUserInfoResponseListener>)listener;
- (GoJimsdkVerifyEmailResponse*)sendVerifyEmail:(NSString*)email;
- (void)sendVerifyEmailAsync:(NSString*)email listener:(id<GoJimsdkVerifyEmailResponseListener>)listener;
- (GoJimsdkVerifySmsResponse*)sendVerifySms:(NSString*)phone;
- (GoJimsdkWeiboUserResponse*)sendWeiboUser:(NSString*)sinaUID;
- (GoJimsdkWeixinUserResponse*)sendWeixinUser:(NSString*)openID;
@end

@interface GoJimsdkFacebookUserParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)appID;
- (void)setAppID:(long)v;
- (NSString*)accessToken;
- (void)setAccessToken:(NSString*)v;
@end

@interface GoJimsdkFacebookUserResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (BOOL)result;
- (void)setResult:(BOOL)v;
- (NSString*)infoID;
- (void)setInfoID:(NSString*)v;
- (NSString*)infoName;
- (void)setInfoName:(NSString*)v;
- (NSString*)infoEmail;
- (void)setInfoEmail:(NSString*)v;
- (NSString*)infoGender;
- (void)setInfoGender:(NSString*)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkFaqListParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)appID;
- (void)setAppID:(long)v;
- (NSString*)language;
- (void)setLanguage:(NSString*)v;
- (long)from;
- (void)setFrom:(long)v;
- (long)size;
- (void)setSize:(long)v;
@end

@interface GoJimsdkFaqListResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (NSString*)text;
- (void)setText:(NSString*)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkFeedbackSubmitParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (NSString*)contactInfo;
- (void)setContactInfo:(NSString*)v;
- (NSString*)content;
- (void)setContent:(NSString*)v;
@end

@interface GoJimsdkFeedbackSubmitResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (NSString*)contactInfo;
- (void)setContactInfo:(NSString*)v;
- (NSString*)content;
- (void)setContent:(NSString*)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkFirmwareOtaInfoResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (NSString*)limit;
- (void)setLimit:(NSString*)v;
- (NSString*)downloadURL;
- (void)setDownloadURL:(NSString*)v;
- (NSString*)version;
- (void)setVersion:(NSString*)v;
- (NSString*)hardwareVersion;
- (void)setHardwareVersion:(NSString*)v;
- (NSString*)description;
- (void)setDescription:(NSString*)v;
- (NSString*)date;
- (void)setDate:(NSString*)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkHasPhoneParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)appID;
- (void)setAppID:(long)v;
- (NSString*)phone;
- (void)setPhone:(NSString*)v;
@end

@interface GoJimsdkHasPhoneResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (BOOL)result;
- (void)setResult:(BOOL)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkLinkedInUserParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)appID;
- (void)setAppID:(long)v;
- (NSString*)accessToken;
- (void)setAccessToken:(NSString*)v;
@end

@interface GoJimsdkLinkedInUserResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (BOOL)result;
- (void)setResult:(BOOL)v;
- (NSString*)infoID;
- (void)setInfoID:(NSString*)v;
- (NSString*)infoName;
- (void)setInfoName:(NSString*)v;
- (NSString*)infoEmail;
- (void)setInfoEmail:(NSString*)v;
- (NSString*)infoGender;
- (void)setInfoGender:(NSString*)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkLoginParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)appID;
- (void)setAppID:(long)v;
- (NSString*)username;
- (void)setUsername:(NSString*)v;
- (NSString*)password;
- (void)setPassword:(NSString*)v;
- (NSString*)email;
- (void)setEmail:(NSString*)v;
- (NSString*)phone;
- (void)setPhone:(NSString*)v;
- (NSString*)weixinOpenID;
- (void)setWeixinOpenID:(NSString*)v;
- (NSString*)qqOpenID;
- (void)setQqOpenID:(NSString*)v;
- (NSString*)sinaUID;
- (void)setSinaUID:(NSString*)v;
- (NSString*)facebookID;
- (void)setFacebookID:(NSString*)v;
- (NSString*)twitterID;
- (void)setTwitterID:(NSString*)v;
- (NSString*)linkedInID;
- (void)setLinkedInID:(NSString*)v;
@end

@interface GoJimsdkNewsDigest : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (NSString*)title;
- (void)setTitle:(NSString*)v;
- (NSString*)articleURL;
- (void)setArticleURL:(NSString*)v;
- (NSString*)thumbURL;
- (void)setThumbURL:(NSString*)v;
@end

@interface GoJimsdkNewsDigestCollection : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
// skipped field NewsDigestCollection.Items with unsupported type: *types.Slice

- (GoJimsdkNewsDigest*)getItemAt:(long)index;
- (long)getSize;
@end

@interface GoJimsdkNewsDigestParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)appID;
- (void)setAppID:(long)v;
- (long)fromPage;
- (void)setFromPage:(long)v;
- (long)pageSize;
- (void)setPageSize:(long)v;
- (long)thumbWidth;
- (void)setThumbWidth:(long)v;
- (long)thumbHeight;
- (void)setThumbHeight:(long)v;
- (NSString*)tags;
- (void)setTags:(NSString*)v;
- (NSString*)language;
- (void)setLanguage:(NSString*)v;
@end

@interface GoJimsdkNewsDigestResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (GoJimsdkNewsDigestCollection*)collection;
- (void)setCollection:(GoJimsdkNewsDigestCollection*)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkQqUserParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)appID;
- (void)setAppID:(long)v;
- (NSString*)openID;
- (void)setOpenID:(NSString*)v;
@end

@interface GoJimsdkQqUserResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (BOOL)result;
- (void)setResult:(BOOL)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkRegisterInfoParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)appID;
- (void)setAppID:(long)v;
- (NSString*)username;
- (void)setUsername:(NSString*)v;
- (NSString*)password;
- (void)setPassword:(NSString*)v;
- (NSString*)phone;
- (void)setPhone:(NSString*)v;
- (NSString*)email;
- (void)setEmail:(NSString*)v;
- (NSString*)weixinOpenID;
- (void)setWeixinOpenID:(NSString*)v;
- (NSString*)qqOpenID;
- (void)setQqOpenID:(NSString*)v;
- (NSString*)sinaUID;
- (void)setSinaUID:(NSString*)v;
- (NSString*)facebookID;
- (void)setFacebookID:(NSString*)v;
- (NSString*)twitterID;
- (void)setTwitterID:(NSString*)v;
- (NSString*)linkedInID;
- (void)setLinkedInID:(NSString*)v;
- (NSString*)verificationCode;
- (void)setVerificationCode:(NSString*)v;
- (NSString*)nickname;
- (void)setNickname:(NSString*)v;
- (float)height;
- (void)setHeight:(float)v;
- (float)weight;
- (void)setWeight:(float)v;
- (long)gender;
- (void)setGender:(long)v;
- (NSString*)caseHistory;
- (void)setCaseHistory:(NSString*)v;
- (NSString*)birthday;
- (void)setBirthday:(NSString*)v;
- (NSString*)avatarURL;
- (void)setAvatarURL:(NSString*)v;
@end

@interface GoJimsdkRegisterParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)appID;
- (void)setAppID:(long)v;
- (NSString*)username;
- (void)setUsername:(NSString*)v;
- (NSString*)password;
- (void)setPassword:(NSString*)v;
- (NSString*)phone;
- (void)setPhone:(NSString*)v;
- (NSString*)email;
- (void)setEmail:(NSString*)v;
- (NSString*)weixinOpenID;
- (void)setWeixinOpenID:(NSString*)v;
- (NSString*)qqOpenID;
- (void)setQqOpenID:(NSString*)v;
- (NSString*)sinaUID;
- (void)setSinaUID:(NSString*)v;
- (NSString*)facebookID;
- (void)setFacebookID:(NSString*)v;
- (NSString*)twitterID;
- (void)setTwitterID:(NSString*)v;
- (NSString*)linkedInID;
- (void)setLinkedInID:(NSString*)v;
- (NSString*)verificationCode;
- (void)setVerificationCode:(NSString*)v;
@end

@interface GoJimsdkResetPasswordEmailParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)appID;
- (void)setAppID:(long)v;
- (NSString*)email;
- (void)setEmail:(NSString*)v;
@end

@interface GoJimsdkResetPasswordEmailResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (BOOL)result;
- (void)setResult:(BOOL)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkResetPasswordParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)appID;
- (void)setAppID:(long)v;
- (NSString*)phone;
- (void)setPhone:(NSString*)v;
- (NSString*)email;
- (void)setEmail:(NSString*)v;
- (NSString*)verificationCode;
- (void)setVerificationCode:(NSString*)v;
- (NSString*)password;
- (void)setPassword:(NSString*)v;
@end

@interface GoJimsdkResetPasswordResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (BOOL)result;
- (void)setResult:(BOOL)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkResetPasswordSmsParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)appID;
- (void)setAppID:(long)v;
- (NSString*)phone;
- (void)setPhone:(NSString*)v;
@end

@interface GoJimsdkResetPasswordSmsResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (BOOL)result;
- (void)setResult:(BOOL)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkResponseError : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (NSString*)key;
- (void)setKey:(NSString*)v;
- (NSString*)message;
- (void)setMessage:(NSString*)v;
@end

@interface GoJimsdkSystemInfoResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (BOOL)result;
- (void)setResult:(BOOL)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkTwitterUserParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)appID;
- (void)setAppID:(long)v;
- (NSString*)userToken;
- (void)setUserToken:(NSString*)v;
- (NSString*)tokenSecret;
- (void)setTokenSecret:(NSString*)v;
@end

@interface GoJimsdkTwitterUserResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (BOOL)result;
- (void)setResult:(BOOL)v;
- (NSString*)infoID;
- (void)setInfoID:(NSString*)v;
- (NSString*)infoName;
- (void)setInfoName:(NSString*)v;
- (NSString*)infoEmail;
- (void)setInfoEmail:(NSString*)v;
- (NSString*)infoGender;
- (void)setInfoGender:(NSString*)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkUpdateBindEmailParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)userID;
- (void)setUserID:(long)v;
- (NSString*)email;
- (void)setEmail:(NSString*)v;
@end

@interface GoJimsdkUpdateBindEmailResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (BOOL)result;
- (void)setResult:(BOOL)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkUpdateBindPhoneParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)userID;
- (void)setUserID:(long)v;
- (NSString*)phone;
- (void)setPhone:(NSString*)v;
@end

@interface GoJimsdkUpdateBindPhoneResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (BOOL)result;
- (void)setResult:(BOOL)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkUpdateUserParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)subUserID;
- (void)setSubUserID:(long)v;
- (NSString*)nickname;
- (void)setNickname:(NSString*)v;
- (long)height;
- (void)setHeight:(long)v;
- (long)weight;
- (void)setWeight:(long)v;
- (long)gender;
- (void)setGender:(long)v;
- (NSString*)birthday;
- (void)setBirthday:(NSString*)v;
- (NSString*)caseHistory;
- (void)setCaseHistory:(NSString*)v;
@end

@interface GoJimsdkUpdateUserResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (int64_t)id_;
- (void)setID:(int64_t)v;
- (NSString*)username;
- (void)setUsername:(NSString*)v;
- (int64_t)registerTime;
- (void)setRegisterTime:(int64_t)v;
- (NSString*)email;
- (void)setEmail:(NSString*)v;
- (BOOL)emailChecked;
- (void)setEmailChecked:(BOOL)v;
- (NSString*)phone;
- (void)setPhone:(NSString*)v;
- (BOOL)phoneChecked;
- (void)setPhoneChecked:(BOOL)v;
- (NSString*)infoBirthday;
- (void)setInfoBirthday:(NSString*)v;
- (NSString*)infoCaseHistory;
- (void)setInfoCaseHistory:(NSString*)v;
- (NSString*)infoNickname;
- (void)setInfoNickname:(NSString*)v;
- (long)infoHeight;
- (void)setInfoHeight:(long)v;
- (long)infoWeight;
- (void)setInfoWeight:(long)v;
- (long)infoGender;
- (void)setInfoGender:(long)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkUploadAvatarResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (NSString*)url;
- (void)setURL:(NSString*)v;
- (NSString*)message;
- (void)setMessage:(NSString*)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkUserInfoParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)userID;
- (void)setUserID:(long)v;
- (long)subUserID;
- (void)setSubUserID:(long)v;
@end

@interface GoJimsdkUserInfoResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (int64_t)id_;
- (void)setID:(int64_t)v;
- (NSString*)username;
- (void)setUsername:(NSString*)v;
- (int64_t)registerTime;
- (void)setRegisterTime:(int64_t)v;
- (NSString*)avatarURL;
- (void)setAvatarURL:(NSString*)v;
- (NSString*)email;
- (void)setEmail:(NSString*)v;
- (BOOL)emailChecked;
- (void)setEmailChecked:(BOOL)v;
- (NSString*)phone;
- (void)setPhone:(NSString*)v;
- (BOOL)phoneChecked;
- (void)setPhoneChecked:(BOOL)v;
- (NSString*)infoBirthday;
- (void)setInfoBirthday:(NSString*)v;
- (NSString*)infoCaseHistory;
- (void)setInfoCaseHistory:(NSString*)v;
- (NSString*)infoNickname;
- (void)setInfoNickname:(NSString*)v;
- (long)infoHeight;
- (void)setInfoHeight:(long)v;
- (long)infoWeight;
- (void)setInfoWeight:(long)v;
- (long)infoGender;
- (void)setInfoGender:(long)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkVerifyEmailParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)appID;
- (void)setAppID:(long)v;
- (NSString*)email;
- (void)setEmail:(NSString*)v;
@end

@interface GoJimsdkVerifyEmailResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (BOOL)result;
- (void)setResult:(BOOL)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkVerifySmsParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)appID;
- (void)setAppID:(long)v;
- (NSString*)phone;
- (void)setPhone:(NSString*)v;
@end

@interface GoJimsdkVerifySmsResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (BOOL)result;
- (void)setResult:(BOOL)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkWeiboUserParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)appID;
- (void)setAppID:(long)v;
- (NSString*)sinaUID;
- (void)setSinaUID:(NSString*)v;
@end

@interface GoJimsdkWeiboUserResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (BOOL)result;
- (void)setResult:(BOOL)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkWeixinUserResponse : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (BOOL)result;
- (void)setResult:(BOOL)v;
- (GoJimsdkResponseError*)error;
- (void)setError:(GoJimsdkResponseError*)v;
@end

@interface GoJimsdkWexinUserParams : NSObject {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (long)appID;
- (void)setAppID:(long)v;
- (NSString*)openID;
- (void)setOpenID:(NSString*)v;
@end

@protocol GoJimsdkLoginResponseListener
- (void)onFailure:(GoJimsdkResponseError*)respErr;
- (void)onSuccess:(GoJimsdkUserInfoResponse*)respData;
@end

@protocol GoJimsdkNewsDigestResponseListener
- (void)onFailure:(GoJimsdkResponseError*)respErr;
- (void)onSuccess:(GoJimsdkNewsDigestResponse*)respData;
@end

@protocol GoJimsdkRegisterResponseListener
- (void)onFailure:(GoJimsdkResponseError*)respErr;
- (void)onSuccess:(GoJimsdkUserInfoResponse*)respData;
@end

@protocol GoJimsdkUserInfoResponseListener
- (void)onFailure:(GoJimsdkResponseError*)respErr;
- (void)onSuccess:(GoJimsdkUserInfoResponse*)respData;
@end

@protocol GoJimsdkVerifyEmailResponseListener
- (void)onFailure:(GoJimsdkResponseError*)respErr;
- (void)onSuccess:(GoJimsdkVerifyEmailResponse*)respData;
@end

FOUNDATION_EXPORT NSString* const GoJimsdkAppVersionInfoRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkBindEmailRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkBindPhoneRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkBloodPressureCommitListRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkBloodPressureCommitRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkBloodPressureDeleteRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkBloodPressureDeviceTroubleRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkBloodPressureListRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkBloodPressureSyncRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkBloodPressureTotalCountRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkChangePasswordRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkFacebookUserRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkFaqListRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkFeedbackSubmitRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkFirmwareLastOtaRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkFirmwareOtaInfoRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkHasPhoneRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkLinkedInUserRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkLoginRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkNewsDigestRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkQqUserRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkRegisterInfoRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkRegisterRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkResetPasswordEmailRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkResetPasswordRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkResetPasswordSmsRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkSystemInfoRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkTwitterUserRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkUpdateBindEmailRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkUpdateBindPhoneRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkUpdateUserRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkUploadAvatarBase64Router;
FOUNDATION_EXPORT NSString* const GoJimsdkUploadAvatarRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkUserInfoRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkVerifyEmailRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkVerifySmsRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkWeiboUserRouter;
FOUNDATION_EXPORT NSString* const GoJimsdkWeixinUserRouter;

FOUNDATION_EXPORT BOOL GoJimsdkCatchResponseError(GoJimsdkResponseError* respError);

FOUNDATION_EXPORT GoJimsdkBloodPressureCommitListParams* GoJimsdkNewBloodPressureCommitListParams();

FOUNDATION_EXPORT GoJimsdkBloodPressureCommitParams* GoJimsdkNewBloodPressureCommitParams();

FOUNDATION_EXPORT GoJimsdkBloodPressureDeleteParams* GoJimsdkNewBloodPressureDeleteParams();

FOUNDATION_EXPORT BOOL GoJimsdkNewClient(NSString* clusterURL, long appID, NSString* jimAppID, NSString* jimAppSecret, NSString* cookieFilePath, GoJimsdkClient** ret0_, NSError** error);

FOUNDATION_EXPORT GoJimsdkFaqListParams* GoJimsdkNewFaqListParams();

FOUNDATION_EXPORT GoJimsdkLoginParams* GoJimsdkNewLoginParams();

FOUNDATION_EXPORT GoJimsdkNewsDigestParams* GoJimsdkNewNewsDigestParams();

FOUNDATION_EXPORT GoJimsdkRegisterInfoParams* GoJimsdkNewRegisterInfoParams();

FOUNDATION_EXPORT GoJimsdkRegisterParams* GoJimsdkNewRegisterParams();

FOUNDATION_EXPORT GoJimsdkResetPasswordParams* GoJimsdkNewResetPasswordParams();

FOUNDATION_EXPORT GoJimsdkUpdateUserParams* GoJimsdkNewUpdateUserParams();

@class GoJimsdkLoginResponseListener;

@class GoJimsdkNewsDigestResponseListener;

@class GoJimsdkRegisterResponseListener;

@class GoJimsdkUserInfoResponseListener;

@class GoJimsdkVerifyEmailResponseListener;

@interface GoJimsdkLoginResponseListener : NSObject <GoJimsdkLoginResponseListener> {
}
@property(strong, readonly) id _ref;

- (instancetype)initWithRef:(id)ref;
- (void)onFailure:(GoJimsdkResponseError*)respErr;
- (void)onSuccess:(GoJimsdkUserInfoResponse*)respData;
@end

@interface GoJimsdkNewsDigestResponseListener : NSObject <GoJimsdkNewsDigestResponseListener> {
}
@property(strong, readonly) id _ref;

- (instancetype)initWithRef:(id)ref;
- (void)onFailure:(GoJimsdkResponseError*)respErr;
- (void)onSuccess:(GoJimsdkNewsDigestResponse*)respData;
@end

@interface GoJimsdkRegisterResponseListener : NSObject <GoJimsdkRegisterResponseListener> {
}
@property(strong, readonly) id _ref;

- (instancetype)initWithRef:(id)ref;
- (void)onFailure:(GoJimsdkResponseError*)respErr;
- (void)onSuccess:(GoJimsdkUserInfoResponse*)respData;
@end

@interface GoJimsdkUserInfoResponseListener : NSObject <GoJimsdkUserInfoResponseListener> {
}
@property(strong, readonly) id _ref;

- (instancetype)initWithRef:(id)ref;
- (void)onFailure:(GoJimsdkResponseError*)respErr;
- (void)onSuccess:(GoJimsdkUserInfoResponse*)respData;
@end

@interface GoJimsdkVerifyEmailResponseListener : NSObject <GoJimsdkVerifyEmailResponseListener> {
}
@property(strong, readonly) id _ref;

- (instancetype)initWithRef:(id)ref;
- (void)onFailure:(GoJimsdkResponseError*)respErr;
- (void)onSuccess:(GoJimsdkVerifyEmailResponse*)respData;
@end

#endif
