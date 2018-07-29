# serverlessui [![CircleCI](https://circleci.com/gh/larse514/serverlessui.svg?style=svg)](https://circleci.com/gh/larse514/serverlessui) [![Go Report Card](https://goreportcard.com/badge/github.com/larse514/serverlessui)](https://goreportcard.com/report/github.com/larse514/serverlessui) 

CLI application to deploy serverless ui to Cloud providers like AWS


## Parameters
1. HostedZone DNS name of an existing Amazon Route 53 hosted zone e.g. vssdevelopment.com (Required)
2. DomainName The full domain name e.g. www.vssdevelopment.com (Required)
3. CacheValueTTL CDN cache time to live
4. HostedZoneExists Parameter to determine if HostedZone needs to be created (Optional)
5. Tag of hosted zone, used to tag resources for tracking and billing (Optional)
