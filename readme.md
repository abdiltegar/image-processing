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
The GoCV package provides Go language bindings for the OpenCV 4 computer vision library. To install GoCV, you can follow <a href="https://gocv.io/getting-started/" target="_blank">this instructions</a>.


## **Install Other Dependencies**

```
go mod tidy
```

# How to Use

## Configuration
Copy `.env.example` file and rename it to `.env`. Open it and set the `SERVER_HOST` and `SERVER_PORT`.

## Run
To run the program, go to project directory and you can simply run this command<br>
```
go run main.go
```

## Documentation
Documentation about how to use this service can be seen at <br>
```
{{your_host}}/swagger/index.html
```

after running the program. Or if you want to use Postman documentation instead of Swagger, you can go to this <a href="https://documenter.getpostman.com/view/29782284/2sA2xny9wg" target="_blank">link</a>.

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