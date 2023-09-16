package usecase

import (
	"sellease-ai/internal/repository"
	"sellease-ai/internal/usecase/fileproc"
	"sellease-ai/internal/usecase/product"
)

type Usecase struct {
	// User user.UsecaseInterface
	Product  product.UsecaseInterface
	FileProc fileproc.UsecaseInterface
}

func Init(r *repository.Repository) *Usecase {
	// userUseCase := user.InitUserUsecase(r.User)
	productUseCase := product.InitProductUsecase(r.Product)
	fileProcUseCase := fileproc.InitFileProcUsecase(r.FileProc)
	return &Usecase{
		Product:  productUseCase,
		FileProc: fileProcUseCase,
	}
}
