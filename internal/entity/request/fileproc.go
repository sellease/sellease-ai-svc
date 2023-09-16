package request

import "mime/multipart"

type FileUploadRequest struct {
	CSVFile *multipart.FileHeader `form:"csv_file" binding:"required"`
}
