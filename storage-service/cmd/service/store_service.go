package service

import (
	"context"

	st "github.com/fr13nd230/nebula-fs/storage-service/grpc/storage"
	"github.com/fr13nd230/nebula-fs/storage-service/repository/store"
	"go.uber.org/zap"
)

type StorageService struct {
	q *store.Queries
	sg *zap.SugaredLogger
}

func NewStorageService(q *store.Queries, sg *zap.SugaredLogger) *StorageService {
	return &StorageService{
		q: q,
		sg: sg,
	}
}

func (s *StorageService) Store(ctx context.Context, chunk *st.FileChunk) error {
	t, err := s.q.HealthCheck(ctx)
	s.sg.Infow("Service has been hit", "ready?", t.Valid, "chunk:data", string(chunk.Data))
	return err
}
