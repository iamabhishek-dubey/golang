package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/awserr"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	// Arguments Required by This Program
	var (
        actionPtr = flag.String("action", "", "Action to perform on Instance. Values could be, start or stop")
        instancePtr = flag.String("instance-id", "", "Id of the instance to perform action")
	)
	flag.Parse()
	// Load session from shared config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create new EC2 client
	svc := ec2.New(sess)

		if *actionPtr == "start" {
			input := &ec2.StartInstancesInput{
				InstanceIds: []*string{
					aws.String(*instancePtr),
				},
				DryRun: aws.Bool(true),
			}
			result, err := svc.StartInstances(input)
			awsErr, ok := err.(awserr.Error)

		if ok && awsErr.Code() == "DryRunOperation" {
			input.DryRun = aws.Bool(false)
			result, err = svc.StartInstances(input)
			if err != nil {
				fmt.Println("Error", err)
			} else {
				fmt.Println("Success", result.StartingInstances)
			}
		} else {
			fmt.Println("Error", err)
		}
	} else if *actionPtr == "stop" {
		input := &ec2.StopInstancesInput {
			InstanceIds: []*string{
				aws.String(*instancePtr),
			},
			DryRun: aws.Bool(true),
		}
		result, err := svc.StopInstances(input)
		awsErr, ok := err.(awserr.Error)
		if ok && awsErr.Code() == "DryRunOperation" {
			input.DryRun = aws.Bool(false)
			result, err = svc.StopInstances(input)
			if err != nil {
				fmt.Println("Error", err)
			} else {
				fmt.Println("Success", result.StoppingInstances)
			}
		} else {
			fmt.Println("Error", err)
		}
	}
}
