package main

import (
	"flag"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	confFileName = flag.String("conf", "", "config file")
	config       *Config
	client       *dynamodb.DynamoDB
)

func main() {
	flag.Parse()

	var err error
	config, err = LoadConfig(*confFileName)
	if err != nil {
		panic(err)
	}

	client = NewClient(config)

	routes()

	serve(config)
}
