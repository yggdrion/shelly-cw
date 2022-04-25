package clients

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go/aws"
)

type CWPutMetricDataAPI interface {
	PutMetricData(ctx context.Context,
		params *cloudwatch.PutMetricDataInput,
		optFns ...func(*cloudwatch.Options)) (*cloudwatch.PutMetricDataOutput, error)
}

func CreateCustomMetric(c context.Context, api CWPutMetricDataAPI, input *cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error) {
	return api.PutMetricData(c, input)
}

func PutCloudwatchMetric(name string, value float64) {

	// write metrics to cloudwatch

	// TODO Read env only once
	awsAccessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	// TODO error handling
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-central-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAccessKeyId, awsSecretAccessKey, "")))
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := cloudwatch.NewFromConfig(cfg)

	input := &cloudwatch.PutMetricDataInput{
		Namespace: aws.String("custom"),
		MetricData: []types.MetricDatum{
			{
				MetricName: aws.String("watt"),
				Unit:       types.StandardUnitSeconds,
				Value:      aws.Float64(value),
				Dimensions: []types.Dimension{
					{
						Name:  aws.String("name"),
						Value: aws.String(name),
					},
				},
			},
		},
	}

	_, err = CreateCustomMetric(context.TODO(), client, input)
	if err != nil {
		log.Println(err)
		// TODO exit or return here
		//return
		os.Exit(1)
	}

}
