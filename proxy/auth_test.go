package proxy

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestActions(t *testing.T) {
	tests := map[string]struct {
		method             string
		url                string
		expectedAction     string
		expectedPermission string
		expectedErr        error
	}{
		"AbortMultipartUpload": {
			method:         http.MethodDelete,
			url:            "/bucket/my/test/key?uploadId=2",
			expectedAction: "AbortMultipartUpload",
		},
		"CompleteMultipartUpload": {
			method:         http.MethodPost,
			url:            "/bucket/my/test/key?uploadId=2",
			expectedAction: "CompleteMultipartUpload",
		},
		// "CopyObject": {
		// 	method:         http.MethodPut,
		// 	url:            "/bucket/my/test/key",
		// 	expectedAction: "CopyObject",
		// },
		"CreateBucket": {
			method:         http.MethodPut,
			url:            "/new-bucket/",
			expectedAction: "CreateBucket",
		},
		"CreateMultipartUpload": {
			method:         http.MethodPost,
			url:            "/bucket/my/test/key?uploads",
			expectedAction: "CreateMultipartUpload",
		},
		"DeleteBucket": {
			method:         http.MethodDelete,
			url:            "/bucket/",
			expectedAction: "DeleteBucket",
		},
		"DeleteBucketAnalyticsConfiguration": {
			method:         http.MethodDelete,
			url:            "/bucket/?analytics&id=3",
			expectedAction: "DeleteBucketAnalyticsConfiguration",
		},
		"DeleteBucketCors": {
			method:         http.MethodDelete,
			url:            "/bucket/?cors",
			expectedAction: "DeleteBucketCors",
		},
		"DeleteBucketEncryption": {
			method:         http.MethodDelete,
			url:            "/bucket/?encryption",
			expectedAction: "DeleteBucketEncryption",
		},
		"DeleteBucketIntelligentTieringConfiguration": {
			method:         http.MethodDelete,
			url:            "/bucket/?intelligent-tiering&id=4",
			expectedAction: "DeleteBucketIntelligentTieringConfiguration",
		},
		"DeleteBucketInventoryConfiguration": {
			method:         http.MethodDelete,
			url:            "/bucket/?inventory&id=5",
			expectedAction: "DeleteBucketInventoryConfiguration",
		},
		"DeleteBucketLifecycle": {
			method:         http.MethodDelete,
			url:            "/bucket/?lifecycle",
			expectedAction: "DeleteBucketLifecycle",
		},
		"DeleteBucketMetricsConfiguration": {
			method:         http.MethodDelete,
			url:            "/bucket/?metrics&id=6",
			expectedAction: "DeleteBucketMetricsConfiguration",
		},
		"DeleteBucketOwnershipControls": {
			method:         http.MethodDelete,
			url:            "/bucket/?ownershipControls",
			expectedAction: "DeleteBucketOwnershipControls",
		},
		"DeleteBucketPolicy": {
			method:         http.MethodDelete,
			url:            "/bucket/?policy",
			expectedAction: "DeleteBucketPolicy",
		},
		"DeleteBucketReplication": {
			method:         http.MethodDelete,
			url:            "/bucket/?replication",
			expectedAction: "DeleteBucketReplication",
		},
		"DeleteBucketTagging": {
			method:         http.MethodDelete,
			url:            "/bucket/?tagging",
			expectedAction: "DeleteBucketTagging",
		},
		"DeleteBucketWebsite": {
			method:         http.MethodDelete,
			url:            "/bucket/?website",
			expectedAction: "DeleteBucketWebsite",
		},
		"DeleteObject": {
			method:         http.MethodDelete,
			url:            "/bucket/my/test/key",
			expectedAction: "DeleteObject",
		},
		"DeleteObjects": {
			method:         http.MethodPost,
			url:            "/bucket/?delete",
			expectedAction: "DeleteObjects",
		},
		"DeleteObjectTagging": {
			method:         http.MethodDelete,
			url:            "/bucket/my/test/key?tagging",
			expectedAction: "DeleteObjectTagging",
		},
		"DeletePublicAccessBlock": {
			method:         http.MethodDelete,
			url:            "/bucket/?publicAccessBlock",
			expectedAction: "DeletePublicAccessBlock",
		},
		"GetBucketAccelerateConfiguration": {
			method:         http.MethodGet,
			url:            "/bucket/?accelerate",
			expectedAction: "GetBucketAccelerateConfiguration",
		},
		"GetBucketAcl": {
			method:         http.MethodGet,
			url:            "/bucket/?acl",
			expectedAction: "GetBucketAcl",
		},
		"GetBucketAnalyticsConfiguration": {
			method:         http.MethodGet,
			url:            "/bucket/?analytics&id=65",
			expectedAction: "GetBucketAnalyticsConfiguration",
		},
		"GetBucketCors": {
			method:         http.MethodGet,
			url:            "/bucket/?cors",
			expectedAction: "GetBucketCors",
		},
		"GetBucketEncryption": {
			method:         http.MethodGet,
			url:            "/bucket/?encryption",
			expectedAction: "GetBucketEncryption",
		},
		"GetBucketIntelligentTieringConfiguration": {
			method:         http.MethodGet,
			url:            "/bucket/?intelligent-tiering&id=7",
			expectedAction: "GetBucketIntelligentTieringConfiguration",
		},
		"GetBucketInventoryConfiguration": {
			method:         http.MethodGet,
			url:            "/bucket/?inventory&id=7",
			expectedAction: "GetBucketInventoryConfiguration",
		},
		"GetBucketLifecycleConfiguration": {
			method:         http.MethodGet,
			url:            "/bucket/?lifecycle",
			expectedAction: "GetBucketLifecycleConfiguration",
		},
		"GetBucketLocation": {
			method:         http.MethodGet,
			url:            "/bucket/?location",
			expectedAction: "GetBucketLocation",
		},
		"GetBucketLogging": {
			method:         http.MethodGet,
			url:            "/bucket/?logging",
			expectedAction: "GetBucketLogging",
		},
		"GetBucketMetricsConfiguration": {
			method:         http.MethodGet,
			url:            "/bucket/?metrics&id=8",
			expectedAction: "GetBucketMetricsConfiguration",
		},
		"GetBucketNotificationConfiguration": {
			method:         http.MethodGet,
			url:            "/bucket/?notification",
			expectedAction: "GetBucketNotificationConfiguration",
		},
		"GetBucketOwnershipControls": {
			method:         http.MethodGet,
			url:            "/bucket/?ownershipControls",
			expectedAction: "GetBucketOwnershipControls",
		},
		"GetBucketPolicy": {
			method:         http.MethodGet,
			url:            "/bucket/?policy",
			expectedAction: "GetBucketPolicy",
		},
		"GetBucketPolicyStatus": {
			method:         http.MethodGet,
			url:            "/bucket/?policyStatus",
			expectedAction: "GetBucketPolicyStatus",
		},
		"GetBucketReplication": {
			method:         http.MethodGet,
			url:            "/bucket/?replication",
			expectedAction: "GetBucketReplication",
		},
		"GetBucketRequestPayment": {
			method:         http.MethodGet,
			url:            "/bucket/?requestPayment",
			expectedAction: "GetBucketRequestPayment",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest(test.method, test.url, nil)
			require.NoError(t, err)
			action, perm, err := currentAction(req)
			if test.expectedErr != nil {
				assert.Equal(t, test.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedAction, action)
				if test.expectedPermission != "" {
					assert.Equal(t, test.expectedPermission, perm)
				}
			}
		})
	}
}
