# Image-Processing

A Golang backend service with HTTP routes for image processing. The service provide the following functionalities:

1. Convert image files from PNG to JPEG.
2. Resize images according to specified dimensions.
3. Compress images to reduce file size while maintaining reasonable quality.

# How to Install

## **Clone Repository**
```
git clone https://github.com/abdiltegar/image-processing
```

## **Install GoCV**
The GoCV package provides Go language bindings for the OpenCV 4 computer vision library. To install GoCV, you can follow <a href="https://gocv.io/getting-started/" targer="_blank">this instructions</a>.


## **Install Other Dependencies**

```
go mod tidy
```

# How to Use

## Configuration
Copy `.env.example` file and rename it to `.env`. Open it and set the `SERVER_HOST` and `SERVER_PORT`.

## Run
To run the program, go to project directiry and you can simply run this command<br>
```
go run main.go
```

## Swagger Documentation
Swagger documentation about how to use this service can be seen at <br>
```
{{your_host}}/swagger/index.html
```

# How to Execute Test

To test and see the coverage :

1. Execute test and save the coverage to `c.out`
    ```
    go test -v ./... -coverprofile="c.out"
    ```
2. Show coverage
    ```
    go tool cover -html="c.out"
    ```