package handler

import (
	"io"

	"github.com/fr13nd230/nebula-fs/storage-service/cmd/service"
	st "github.com/fr13nd230/nebula-fs/storage-service/grpc/storage"
	"go.uber.org/zap"
)

type StorageHandler struct {
	st.UnimplementedStorageServiceServer
	sg *zap.SugaredLogger
	sv *service.StorageService
}

func NewStorageHandler(sg *zap.SugaredLogger, sv *service.StorageService) *StorageHandler {
	return &StorageHandler{
		sg: sg,
		sv: sv,
	}
}

func (s *StorageHandler) Store(stream st.StorageService_StoreServer) error {
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&st.StorageResponse{
				Status:  true,
				Message: "Streaming has ended.",
				FileId:  chunk.GetId(),
				Node:    1,
			})
		}
		if err != nil {
			s.sg.Errorw("[StorageService]: Storing handler faced an issue.", "error", err)
			return err
		}

		if err := s.sv.Store(stream.Context(), chunk); err != nil {
			s.sg.Error(err)
			return err
		}
	}
}
