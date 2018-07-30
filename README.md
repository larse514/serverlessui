# serverlessui [![CircleCI](https://circleci.com/gh/larse514/serverlessui.svg?style=svg)](https://circleci.com/gh/larse514/serverlessui) [![Go Report Card](https://goreportcard.com/badge/github.com/larse514/serverlessui)](https://goreportcard.com/report/github.com/larse514/serverlessui) 

CLI application to deploy serverless ui to Cloud providers like AWS

## Requirements
The following is required to use serverless-ui cli

1. AWS Registered domain name or hosted zone.  This will be used to create AWS Certificates via the Amazon Certificate Manager
2. UI source code

## Parameters
| Parameter                   | Description                                                                                          | Required | Default |
|-----------------------------|------------------------------------------------------------------------------------------------------|----------|---------|
| -hostedzone, -ho            | HostedZone DNS name of an existing Amazon Route 53 hosted zone e.g. example.com or one to be created | Yes      |         |
| -domainname, -d             | DomainName The full domain name e.g. www.example.com                                                 | Yes      |         |
| -cachettl, -c               | CacheValueTTL CDN cache time to live                                                                 | No       | 60      |
| -hostedzoneexists, -e       | HostedZoneExists Parameter to determine if HostedZone needs to be created                            | No       | false   |
| -tag, -t                    | Tag of hosted zone, used to tag resources for tracking and billing                                   | No       |         |
| -applicationdirectory, -dir | Directory of UI source code to upload                                                                | Yes      |         |
| -environment, -env          | Environment of deployed UI application.  Used to differentiate deployed environments                 | No       | prod    |