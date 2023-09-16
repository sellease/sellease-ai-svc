package fileproc

import (
	"context"
	"sellease-ai/internal/entity/request"
	"sellease-ai/internal/entity/response"
	"sellease-ai/internal/repository/fileproc"
)

type UsecaseInterface interface {
	ProcessFile(ctx context.Context, req request.FileUploadRequest) (resp response.ProcessFileResponse, err error)
}

type fileProcUsecase struct {
	fileProcRepo fileproc.RepositoryInterface
}

func InitFileProcUsecase(fileProc fileproc.RepositoryInterface) UsecaseInterface {
	return &fileProcUsecase{
		fileProcRepo: fileProc,
	}
}
