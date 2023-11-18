package proxy

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var listableBucketOps = map[string]string{
	"analytics":           "BucketAnalyticsConfiguration",
	"intelligent-tiering": "BucketIntelligentTieringConfiguration",
	"inventory":           "BucketInventoryConfiguration",
	"metrics":             "BucketMetricsConfiguration",
	"uploads":             "MultipartUpload",
}

var bucketOps = map[string]string{
	"accelerate":        "BucketAccelerateConfiguration",
	"acl":               "BucketAcl",
	"analytics":         "BucketAnalyticsConfiguration",
	"cors":              "BucketCors",
	"encryption":        "BucketEncryption",
	"intelligent":       "BucketIntelligentTieringConfiguration",
	"inventory":         "BucketInventoryConfiguration",
	"lifecycle":         "BucketLifecycleConfiguration",
	"location":          "BucketLocation",
	"logging":           "BucketLogging",
	"metrics":           "BucketMetricsConfiguration",
	"notification":      "BucketNotificationConfiguration",
	"ownershipControls": "BucketOwnershipControls",
	"policy":            "BucketPolicy",
	"policyStatus":      "BucketPolicyStatus",
	"publicAccessBlock": "PublicAccessBlock",
	"replication":       "BucketReplication",
	"requestPayment":    "BucketRequestPayment",
	"tagging":           "BucketTagging",
	"website":           "BucketWebsite",
}

var objectOps = map[string]string{
	"tagging": "ObjectTagging",
}

func testQueryAction(q *url.Values, actions map[string]string) (string, bool) {
	for param, action := range actions {
		if q.Has(param) {
			return action, true
		}
	}
	return "", false
}

func currentAction(req *http.Request) (string, string, error) {
	bucket, key, _ := strings.Cut(strings.TrimPrefix(req.URL.Path, "/"), "/")
	q := req.URL.Query()

	// global
	if bucket == "" {
		switch req.Method {
		case http.MethodGet:
			return "ListBuckets", "s3:ListAllMyBuckets", nil
		default:
			return "", "", fmt.Errorf("unexpected method on the global level: %s", req.Method)
		}
	}

	// bucket-level
	if key == "" {
		switch req.Method {
		case http.MethodGet:
			if action, ok := testQueryAction(&q, listableBucketOps); ok {
				if q.Get("id") != "" {
					return "Get" + action, "UNSUPPORTED", nil
				}
				return "List" + action + "s", "UNSUPPORTED", nil
			}
			if action, ok := testQueryAction(&q, bucketOps); ok {
				return "Get" + action, "UNSUPPORTED", nil
			}
			if q.Has("versions") {
				return "ListObjectVersions", "s3:ListBucketVersions", nil
			}
			if q.Get("list-type") == "2" {
				return "ListObjectsV2", "s3:ListBucket", nil
			}
			return "ListObjects", "s3:ListBucket", nil
		case http.MethodPost:
			if q.Has("delete") {
				return "DeleteObjects", "s3:DeleteObject", nil
			}
		case http.MethodPut:
			return "CreateBucket", "s3:CreateBucket", nil
		case http.MethodDelete:
			if q.Has("lifecycle") {
				// special case where `listableBucketOps` doesn't work
				return "DeleteBucketLifecycle", "UNSUPPORTED", nil
			}
			if action, ok := testQueryAction(&q, listableBucketOps); ok {
				if q.Get("id") != "" {
					return "Delete" + action, "UNSUPPORTED", nil
				}
			}
			if action, ok := testQueryAction(&q, bucketOps); ok {
				return "Delete" + action, "UNSUPPORTED", nil
			}
			return "DeleteBucket", "s3:DeleteBucket", nil
		}
		return "", "", errors.New("unknown bucket-level action")
	}

	// object level
	switch req.Method {
	case http.MethodGet:

	case http.MethodPost:
		if q.Get("uploadId") != "" {
			return "CompleteMultipartUpload", "UNSUPPORTED", nil
		}
		if q.Has("uploads") {
			return "CreateMultipartUpload", "s3:PutObject", nil
		}

	case http.MethodPut:
	case http.MethodDelete:
		if action, ok := testQueryAction(&q, objectOps); ok {
			return "Delete" + action, "UNSUPPORTED", nil
		}
		if q.Get("uploadId") != "" {
			return "AbortMultipartUpload", "UNSUPPORTED", nil
		}
		return "DeleteObject", "s3:DeleteObject", nil
	}

	return "", "", errors.New("unknown object-level action")
}
