package service

import (
	"io"

	up "github.com/fr13nd230/nebula-fs/uploader-service/grpc/uploader"
	"go.uber.org/zap"
)

type UploaderService struct {
    up.UnimplementedFileUploaderServer
    sugar *zap.SugaredLogger
    
    // TODO: maybe add a mutex, and other fields later.
}

func NewUploader() *UploaderService{
    logger, _ := zap.NewDevelopment() 
    return &UploaderService{
        sugar: logger.Sugar(),
    }
}

func (u UploaderService) Upload(stream up.FileUploader_UploadServer) error {
    for {
        chunk, err := stream.Recv()
        if err == io.EOF {
            return stream.SendAndClose(&up.UploadStatus{
                Status: true,
                FileId: "chunk:" + chunk.GetFileName() + "_" + string(chunk.GetNumber()),
                Message: "File has been uploaded successfully.",
            })
        }
        if err != nil {
            u.sugar.Errorw("[Uploader]: Uploading file has been interrupted.", "error", err)
            return err
        }
        
        // TODO: In here we need to include new CID for each chunk and send both new CID and chunk metadata
        // and chunk it self to both indexer and storage, I am thinking about go routines in here.
        // CID: Content Identifier is a hash of the chunk SHA-256, and also metadata shall include which node is going
        // to store this chunk, both services needs a database respectively.
        return stream.SendMsg(&up.UploadStatus{
            Status: true,
            FileId: "chunk:" + chunk.GetFileName() + "_" + string(chunk.GetNumber()),
            Message: "All files has been accapted.",
        })
    }
}