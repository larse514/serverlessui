package bucket

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

//Uploader is an interface defined to upload an application
type Uploader interface {
	UploadApplication(bucketName string, bucketPrefix string, dirPath string)
}

//S3Uploader is a struct to upload an application to S3
type S3Uploader struct {
	Client s3iface.S3API
}

//UploadApplication is a method to upload an application to s3 bucket
func (uploader S3Uploader) UploadApplication(bucketName string, bucketPrefix string, dirPath string) error {
	fileList := getFilePath(dirPath)
	for _, file := range fileList {
		uploader.uploadFileToS3(bucketName, bucketPrefix, file)
	}
	return nil
}

func (uploader S3Uploader) uploadFileToS3(bucketName string, bucketPrefix string, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to open file", file, err)
		return errors.New("Error opening file")
	}
	//defer close until after method returns
	defer file.Close()
	var key string
	fileDirectory, _ := filepath.Abs(filePath)
	key = bucketPrefix + fileDirectory

	// Upload the file to the s3 given bucket
	params := &s3.PutObjectInput{
		Bucket: aws.String(bucketName), // Required
		Key:    aws.String(key),        // Required
		Body:   file,
	}
	_, err = uploader.Client.PutObject(params)
	if err != nil {
		fmt.Printf("Failed to upload data to %s/%s, %s\n",
			bucketName, key, err.Error())
		return errors.New("Error uploading object")
	}
	return nil
}

//method to check if path is a directory
func isDirectory(path string) (bool, error) {
	fd, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return false, errors.New("Error describing file")
	}
	switch mode := fd.Mode(); {
	case mode.IsDir():
		return false, nil
	case mode.IsRegular():
		return false, nil
	}
	return false, nil
}

//method to walk directory to get array of files
func getFilePath(dirPath string) []string {
	fileList := []string{}
	filepath.Walk(dirPath, func(path string, f os.FileInfo, err error) error {
		fmt.Println("PATH ==> " + path)
		isDirectory, err := isDirectory(path)
		if err != nil {
			return err
		}
		if isDirectory {
			// Do nothing
			return nil
		}
		fileList = append(fileList, path)
		return nil
	})
	return fileList
}
