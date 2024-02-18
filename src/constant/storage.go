package constant

import "time"

const (
	StorageApiBaseUri     = "https://firebasestorage.googleapis.com/v0"
	PublicBucketPath      = "public/"
	PrivateBucketPath     = "private/"
	StorageDefaultTimeout = 5 * time.Second
)
