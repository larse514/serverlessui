# serverlessui
Repository containing CLI application to deploy serverless ui to Cloud providers like AWS


## Parameters
1. HostedZone DNS name of an existing Amazon Route 53 hosted zone e.g. vssdevelopment.com (Required)
2. DomainName The full domain name e.g. www.vssdevelopment.com (Required)
3. CacheValueTTL the Amazon Resource Name (ARN) of an AWS Certificate Manager (ACM) certificate. (Optional)
4. HostedZoneExists Parameter to determine if HostedZone needs to be created (Optional)
5. Tag of hosted zone, used to tag resources for tracking and billing (Optional)


```go
//create flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  hostedZoneArg,
			Usage: "Route 53 `hostedzone`",
		},
		cli.StringFlag{
			Name:  domainNameArg,
			Usage: "`dns name` for serverless ui application",
		},
		cli.StringFlag{
			Name:  cacheTTLArg,
			Usage: "`cache ttl` for Cloudfront cache",
		},
		cli.StringFlag{
			Name:  hostedZoneExistsArg,
			Usage: "Route 53 `hostedzone exists`",
		},
		cli.StringFlag{
			Name:  tagArg,
			Usage: "`tag` used to tag resources for tracking and billing ",
		},
	}
	//create an action
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend! ", c.String(hostedZone), c.String(domainName))
		return nil
    }
```