#import <Foundation/Foundation.h>

typedef void *NSURLCredentialPtr;

typedef struct {
	NSURLSessionAuthChallengeDisposition disposition;
	NSURLCredentialPtr                   credential;
} AuthenticationChallengeResponse;
