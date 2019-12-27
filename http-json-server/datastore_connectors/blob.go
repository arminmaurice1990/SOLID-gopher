package datastore_connectors

import (
	"context"
	"gocloud.dev/blob"
)

func GetBucket(ctx context.Context) (*blob.Bucket , error){
	buck, err := blob.OpenBucket(ctx, "<BLOB_URL>")
	if err != nil {
		return nil,  err
	}
	return buck, nil
}
