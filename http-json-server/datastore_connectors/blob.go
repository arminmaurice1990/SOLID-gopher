package datastore_connectors

import (
	"context"
	"gocloud.dev/blob"
)

const Blob_Url = "BLOB_URL"

func GetBucket(ctx context.Context) (*blob.Bucket, error) {
	buck, err := blob.OpenBucket(ctx, Blob_Url)
	if err != nil {
		return nil, err
	}
	return buck, nil
}
