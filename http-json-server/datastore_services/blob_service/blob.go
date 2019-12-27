package blob_service

import (
	"context"
	"gocloud.dev/blob"
)

type BlobService interface {
	ListBlobs(ctx context.Context) (map[string][]byte, error)
}

type blobservice struct {
	bloburl string
}

func NewBlobConnection(bloburl string) *blobservice {
	return &blobservice{bloburl: bloburl}
}

func (b *blobservice) ListBlobs(ctx context.Context) (map[string][]byte, error) {
	bucket, err := blob.OpenBucket(ctx, b.bloburl)
	if err != nil {
		return nil, err
	}

	keys := []string{}
	opts := &blob.ListOptions{Prefix: "ITEM", Delimiter: ","}
	messageIter := bucket.List(opts)
	for {
		obj, err := messageIter.Next(ctx)
		if err != nil {
			return nil, err
		}
		keys = append(keys, obj.Key)
	}

	messagemap := map[string][]byte{}
	for _, key := range keys {
		readBytes, err := bucket.ReadAll(ctx, key)
		if err != nil {
			return nil, err
		}
		messagemap[key] = readBytes
	}

	return messagemap, nil
}
