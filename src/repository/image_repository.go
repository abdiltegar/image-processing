package repository

import (
	"errors"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
)

type ImageRepository interface {
	SaveTempFile(file *multipart.FileHeader) (string, error)
}

type ImageRepositoryImpl struct {
}

func NewImageRepository() *ImageRepositoryImpl {
	return &ImageRepositoryImpl{}
}

func (repository *ImageRepositoryImpl) SaveTempFile(file *multipart.FileHeader) (string, error) {
	// Save the uploaded file to a temporary location
	tempFile, err := file.Open()
	if err != nil {
		return "", errors.New("error reading file")
	}
	defer tempFile.Close()

	tempFileBytes, err := ioutil.ReadAll(tempFile)
	if err != nil {
		return "", errors.New("error reading file")
	}

	// Write the file to a temporary directory
	tempDir := os.TempDir()
	tempFilePath := filepath.Join(tempDir, file.Filename)
	err = ioutil.WriteFile(tempFilePath, tempFileBytes, 0644)
	if err != nil {
		return "", errors.New("error saving file")
	}

	return tempFilePath, nil
}
