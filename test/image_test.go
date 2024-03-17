package test

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/abdiltegar/image-processing/src/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type imageTestSuite struct {
	suite.Suite
}

func (s *imageTestSuite) SetupSuite() {
}

func TestImageSuite(t *testing.T) {
	suite.Run(t, new(imageTestSuite))
}

func (s *imageTestSuite) TestConvertImage() {

	testCases := []struct {
		Name               string
		Method             string
		FileName           string
		ExpectedStatusCode int
		ExpectedBody       model.WebResponse
	}{
		{
			"Success",
			http.MethodPost,
			"sample.png",
			http.StatusOK,
			model.WebResponse{
				Message: "OK",
				Data:    "/result/sample_converted.jpg",
			},
		},
		{
			"Failed",
			http.MethodPost,
			"",
			http.StatusBadRequest,
			model.WebResponse{
				Message: "Missing or invalid file",
				Data:    nil,
			},
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {

			imageFilePath := filepath.Join(".", "file", testCase.FileName)

			imageFile, err := os.Open(imageFilePath)
			assert.NoError(t, err)
			defer imageFile.Close()

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)

			if testCase.FileName != "" {
				part, err := writer.CreateFormFile("image_file", filepath.Base(imageFilePath))
				assert.NoError(t, err)
				_, err = io.Copy(part, imageFile)
				assert.NoError(t, err)
			}

			writer.Close()

			req := httptest.NewRequest(http.MethodPost, "/api/v1/convert", body)
			req.Header.Set("Content-Type", writer.FormDataContentType())

			response, err := app.Test(req)
			assert.Nil(t, err)

			bytes, err := io.ReadAll(response.Body)
			assert.Nil(t, err)

			responseBody := new(model.WebResponse)
			err = json.Unmarshal(bytes, responseBody)
			assert.Nil(t, err)

			assert.Equal(t, testCase.ExpectedStatusCode, response.StatusCode)
			assert.Equal(t, testCase.ExpectedBody.Data, responseBody.Data)
			assert.Equal(t, testCase.ExpectedBody.Message, responseBody.Message)
		})
	}
}

func (s *imageTestSuite) TestResizeImage() {

	testCases := []struct {
		Name               string
		Method             string
		FileName           string
		BodyParam          model.ResizeRequest
		ExpectedStatusCode int
		ExpectedBody       model.WebResponse
	}{
		{
			"Success",
			http.MethodPost,
			"sample.png",
			model.ResizeRequest{
				Height: 200,
				Width:  200,
			},
			http.StatusOK,
			model.WebResponse{
				Message: "OK",
				Data:    "/result/sample_resized.png",
			},
		},
		{
			"Failed",
			http.MethodPost,
			"",
			model.ResizeRequest{
				Height: 200,
				Width:  200,
			},
			http.StatusBadRequest,
			model.WebResponse{
				Message: "Missing or invalid file",
				Data:    nil,
			},
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {

			imageFilePath := filepath.Join(".", "file", testCase.FileName)

			imageFile, err := os.Open(imageFilePath)
			assert.NoError(t, err)
			defer imageFile.Close()

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)

			writer.WriteField("height", strconv.Itoa(testCase.BodyParam.Height))
			writer.WriteField("width", strconv.Itoa(testCase.BodyParam.Width))

			if testCase.FileName != "" {
				part, err := writer.CreateFormFile("image_file", filepath.Base(imageFilePath))
				assert.NoError(t, err)
				_, err = io.Copy(part, imageFile)
				assert.NoError(t, err)
			}

			writer.Close()

			req := httptest.NewRequest(http.MethodPost, "/api/v1/resize", body)
			req.Header.Set("Content-Type", writer.FormDataContentType())

			response, err := app.Test(req)
			assert.Nil(t, err)

			bytes, err := io.ReadAll(response.Body)
			assert.Nil(t, err)

			responseBody := new(model.WebResponse)
			err = json.Unmarshal(bytes, responseBody)
			assert.Nil(t, err)

			assert.Equal(t, testCase.ExpectedStatusCode, response.StatusCode)
			assert.Equal(t, testCase.ExpectedBody.Data, responseBody.Data)
			assert.Equal(t, testCase.ExpectedBody.Message, responseBody.Message)
		})
	}
}

func (s *imageTestSuite) TestCompressImage() {

	testCases := []struct {
		Name               string
		Method             string
		FileName           string
		ExpectedStatusCode int
		ExpectedBody       model.WebResponse
	}{
		{
			"Success",
			http.MethodPost,
			"sample.png",
			http.StatusOK,
			model.WebResponse{
				Message: "OK",
				Data:    "/result/sample_compressed.png",
			},
		},
		{
			"Failed",
			http.MethodPost,
			"",
			http.StatusBadRequest,
			model.WebResponse{
				Message: "Missing or invalid file",
				Data:    nil,
			},
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {

			imageFilePath := filepath.Join(".", "file", testCase.FileName)

			imageFile, err := os.Open(imageFilePath)
			assert.NoError(t, err)
			defer imageFile.Close()

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)

			if testCase.FileName != "" {
				part, err := writer.CreateFormFile("image_file", filepath.Base(imageFilePath))
				assert.NoError(t, err)
				_, err = io.Copy(part, imageFile)
				assert.NoError(t, err)
			}

			writer.Close()

			req := httptest.NewRequest(http.MethodPost, "/api/v1/compress", body)
			req.Header.Set("Content-Type", writer.FormDataContentType())

			response, err := app.Test(req)
			assert.Nil(t, err)

			bytes, err := io.ReadAll(response.Body)
			assert.Nil(t, err)

			responseBody := new(model.WebResponse)
			err = json.Unmarshal(bytes, responseBody)
			assert.Nil(t, err)

			assert.Equal(t, testCase.ExpectedStatusCode, response.StatusCode)
			assert.Equal(t, testCase.ExpectedBody.Data, responseBody.Data)
			assert.Equal(t, testCase.ExpectedBody.Message, responseBody.Message)
		})
	}
}
