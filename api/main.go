package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	AWS_S3_REGION = "ap-east-1" // Region
	AWS_S3_BUCKET = "cybinon"   // Bucket
)

// We will be using this client everywhere in our code
var awsS3Client *s3.Client

func main() {
	configS3()

	os.Mkdir("../tmp", os.FileMode(0777))

	http.HandleFunc("/upload", handlerUpload) // Upload: /upload (upload file named "file")
	http.HandleFunc("/read/", handlerDownload)
	http.HandleFunc("/list", handlerList) // List: /list?prefix={prefix}&delimeter={delimeter}

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// configS3 creates the S3 client
func configS3() {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(AWS_S3_REGION))
	if err != nil {
		log.Fatal(err)
	}

	awsS3Client = s3.NewFromConfig(cfg)
}

func showError(w http.ResponseWriter, r *http.Request, status int, message string) {
	http.Error(w, message, status)
}
