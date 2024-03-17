package model

type ResizeRequest struct {
	Height int `form:"height" validate:"required"`
	Width  int `form:"width" validate:"required"`
}
