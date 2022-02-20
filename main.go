package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/giantswarm/apiextensions/v3/pkg/apis/provider/v1alpha1"
)

func main() {
	_ = v1alpha1.AzureConfig{}
	_ = aws.Config{Credentials: nil, Region: nil}
}
