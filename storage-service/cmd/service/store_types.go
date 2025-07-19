package service

import (
	"context"

	"github.com/fr13nd230/nebula-fs/storage-service/grpc/storage"
)

type StorageAbstract interface {
	Store(context.Context, *storage.FileChunk) error
}
