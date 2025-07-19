package handler

import "github.com/fr13nd230/nebula-fs/storage-service/grpc/storage"

type StoreAbstract interface {
	Store(storage.StorageService_StoreServer) error
}