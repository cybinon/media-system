package main

import (
	"context"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func handlerDownload(w http.ResponseWriter, r *http.Request) {

	// We get the name of the file on the URL
	key := r.URL.Query().Get("key")
	prevFile, err := os.Open("../tmp/" + key)
	if err == nil {
		defer prevFile.Close()
		http.ServeFile(w, r, key)
		return
	}
	print("gone")
	// Create the file
	newFile, err := os.Create("../tmp/" + key)
	if err != nil {
		print(err.Error())
		showError(w, r, http.StatusBadRequest, "Something went wrong creating the local file")
	}
	defer newFile.Close()

	downloader := manager.NewDownloader(awsS3Client)
	_, err = downloader.Download(context.TODO(), newFile, &s3.GetObjectInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key:    aws.String(key),
	})

	if err != nil {
		showError(w, r, http.StatusBadRequest, "Something went wrong retrieving the file from S3")
		return
	}

	http.ServeFile(w, r, "../tmp/"+key)
}
