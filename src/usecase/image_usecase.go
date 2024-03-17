package usecase

import (
	"errors"
	"image"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/abdiltegar/image-processing/src/helper"
	"github.com/abdiltegar/image-processing/src/repository"
	"gocv.io/x/gocv"
)

type ImageUseCase interface {
	ConvertPNGToJPG(file *multipart.FileHeader) (string, error)
	Resize(file *multipart.FileHeader, width, height int) (string, error)
	Compress(file *multipart.FileHeader) (string, error)
}

type ImageUseCaseImpl struct {
	ImageRepository repository.ImageRepository
}

func NewImageUseCase(imgRepo repository.ImageRepository) *ImageUseCaseImpl {
	return &ImageUseCaseImpl{
		ImageRepository: imgRepo,
	}
}

func (useCase *ImageUseCaseImpl) ConvertPNGToJPG(file *multipart.FileHeader) (string, error) {
	tempFilePath, err := useCase.ImageRepository.SaveTempFile(file)
	if err != nil {
		return "", errors.New("error saving image")
	}

	// Load PNG image
	img := gocv.IMRead(tempFilePath, gocv.IMReadColor)
	if img.Empty() {
		return "", errors.New("error reading image")
	}
	defer img.Close()

	// Save image as JPG
	outputFilePath := filepath.Join(os.TempDir(), helper.RemoveExtension(file.Filename)+"_converted.jpg")
	ok := gocv.IMWrite(outputFilePath, img)
	if !ok {
		return "", errors.New("error writing image")
	}

	// Remove temporary files
	os.Remove(tempFilePath)

	return "/result/" + filepath.Base(outputFilePath), nil
}

func (useCase *ImageUseCaseImpl) Resize(file *multipart.FileHeader, width, height int) (string, error) {
	tempFilePath, err := useCase.ImageRepository.SaveTempFile(file)
	if err != nil {
		return "", errors.New("error saving image")
	}

	// Load image
	img := gocv.IMRead(tempFilePath, gocv.IMReadColor)
	if img.Empty() {
		return "", errors.New("error reading image")
	}
	defer img.Close()

	// Resize the image
	resized := gocv.NewMat()
	gocv.Resize(img, &resized, image.Point{X: width, Y: height}, 0, 0, gocv.InterpolationDefault)

	// Save result image as temp
	outputFilePath := filepath.Join(os.TempDir(), helper.RemoveExtension(file.Filename)+"_resized"+filepath.Ext(file.Filename))
	ok := gocv.IMWrite(outputFilePath, resized)
	if !ok {
		return "", errors.New("error writing image")
	}

	// Remove temporary files
	os.Remove(tempFilePath)

	return "/result/" + filepath.Base(outputFilePath), nil
}

func (useCase *ImageUseCaseImpl) Compress(file *multipart.FileHeader) (string, error) {
	tempFilePath, err := useCase.ImageRepository.SaveTempFile(file)
	if err != nil {
		return "", errors.New("error saving image")
	}

	// Load image
	img := gocv.IMRead(tempFilePath, gocv.IMReadColor)
	if img.Empty() {
		return "", errors.New("error reading image")
	}
	defer img.Close()

	// Define the desired JPEG quality (0-100)
	jpegQuality := 50

	// Convert the image to JPEG with the specified quality
	jpegBytes, err := gocv.IMEncodeWithParams(".jpg", img, []int{int(gocv.IMWriteJpegQuality), jpegQuality})
	if err != nil {
		return "", errors.New("error encoding image")
	}

	compressed, err := gocv.IMDecode(jpegBytes.GetBytes(), gocv.IMReadColor)
	if err != nil {
		return "", errors.New("error decoding image")
	}

	// Save result image as temp
	outputFilePath := filepath.Join(os.TempDir(), helper.RemoveExtension(file.Filename)+"_compressed"+filepath.Ext(file.Filename))
	ok := gocv.IMWrite(outputFilePath, compressed)
	if !ok {
		return "", errors.New("error writing image")
	}

	// Remove temporary files
	os.Remove(tempFilePath)

	return "/result/" + filepath.Base(outputFilePath), nil
}
