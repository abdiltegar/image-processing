package controller

import (
	"github.com/abdiltegar/image-processing/src/model"
	"github.com/abdiltegar/image-processing/src/usecase"
	"github.com/gofiber/fiber/v2"
)

type ImageController struct {
	ImageUseCase usecase.ImageUseCase
}

func NewImageController(imageUseCase usecase.ImageUseCase) *ImageController {
	return &ImageController{
		ImageUseCase: imageUseCase,
	}
}

func (controller *ImageController) Convert(ctx *fiber.Ctx) error {
	// get image file
	file, err := ctx.FormFile("image_file")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Message: "Missing or invalid file",
		})
	}

	// process convert
	image, err := controller.ImageUseCase.ConvertPNGToJPG(file)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse{
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
		Data:    image,
		Message: "OK",
	})
}

func (controller *ImageController) Resize(ctx *fiber.Ctx) error {
	var request model.ResizeRequest

	// parse request to model
	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse{
			Message: err.Error(),
		})
	}

	// get image file
	file, err := ctx.FormFile("image_file")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Message: "Missing or invalid file",
		})
	}

	// process resize
	image, err := controller.ImageUseCase.Resize(file, request.Height, request.Width)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse{
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
		Data:    image,
		Message: "OK",
	})
}

func (controller *ImageController) Compress(ctx *fiber.Ctx) error {
	// get image file
	file, err := ctx.FormFile("image_file")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Message: "Missing or invalid file",
		})
	}

	// process compress
	image, err := controller.ImageUseCase.Compress(file)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse{
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
		Data:    image,
		Message: "OK",
	})
}
