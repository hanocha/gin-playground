package main

import (
	"context"
	"log"
	"os"

	// "./lib/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/gin-gonic/gin"
)

func test() string {
	awsProfile := os.Getenv("AWS_PROFILE")

	if awsProfile == "" {
		awsProfile = "default"
	}

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigProfile(awsProfile),
	)
	if err != nil {
		panic(err)
	}

	cfg.Region = "ap-northeast-1"
	client := ec2.NewFromConfig(cfg)

	res, err := client.DescribeInstances(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	// client.StartInstances(context.TODO(), &ec2.StartInstancesInput{ })

	for _, v := range res.Reservations {
		for _, instances := range v.Instances {
			log.Print(*instances.Tags[0].Value)
		}
	}

	return ""
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/aws", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": test(),
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
