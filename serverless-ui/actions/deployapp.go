package actions

import (
	"log"
	"os"

	"github.com/larse514/serverlessui/serverless-ui/commands"
)

//Bucket is an interface to define creation of Bucket based sites
type Bucket interface {
	DeploySite(input *commands.BucketInput) error
}

//DNS is an interface to represent Cloud DNS Services
type DNS interface {
	DeployHostedZone(input *commands.DNSInput) error
}

//Uploader is an interface defined to upload an application
type Uploader interface {
	UploadApplication(bucketName string, bucketPrefix string, dirPath string) error
}

//ServerlessUI struct to implement DeployAction
type ServerlessUI struct {
	DNS      DNS
	Bucket   Bucket
	Uploader Uploader
}

//Deploy method to deploy serverless UI
func (serverless ServerlessUI) Deploy(dnsInput *commands.DNSInput, bucketInput *commands.BucketInput, appDir string) error {
	err := serverless.DNS.DeployHostedZone(dnsInput)
	if err != nil {
		log.Println("error creating hosted zone ", err)
		os.Exit(1)
	}
	err = serverless.Bucket.DeploySite(bucketInput)
	if err != nil {
		log.Println("error creating hosted zone ", err)
		os.Exit(1)
	}

	return serverless.Uploader.UploadApplication(bucketInput.FullDomainName, "", appDir)
}
