package utils

import (
	"io"
	"mime/multipart"
	"os"
)

func SaveUploadedFile(file *multipart.FileHeader, fileName, dst string) error {
	err := os.MkdirAll(dst, os.ModePerm)
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst + "/" + fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)

	return err
}
