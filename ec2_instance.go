package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"time"
)

func createSession() *session.Session {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{Region: aws.String("us-east-1")},
	}))
	return sess
}

func serviceClient(sess *session.Session) *ec2.EC2 {
	svc := ec2.New(sess)
	return svc
}

func createInstance() *ec2.Reservation {
	svc := serviceClient(createSession())
	parameters := &ec2.RunInstancesInput{
		ImageId:      aws.String("ami-42a2532b"),
		InstanceType: aws.String("t2.micro"),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	}

	runResult, err := svc.RunInstances(parameters)

	if err != nil {
		fmt.Printf("Could not create instance %s", err)
		return nil
	}

	fmt.Println("Instance Successfully Created.")

	return runResult
}


func listInstance(svc *ec2.EC2) []string {
	lst := []string{}
	region := createSession()
	result, _ := svc.DescribeInstances(nil)
	for i := 0; i < len(result.Reservations); i++ {
		if *result.Reservations[i].Instances[0].State.Name == "running" || *result.Reservations[i].Instances[0].State.Name == "pending" {
			fmt.Println("Instance ID", *result.Reservations[i].Instances[0].InstanceId, "," ,
				"Instance Region", *region.Config.Region, ",",
				"Instance IP", *result.Reservations[i].Instances[0].PublicIpAddress)
			lst = append(lst, *result.Reservations[i].Instances[0].InstanceId)
		}
	}
	return lst
}

func terminateInstance(svc *ec2.EC2) {
	lst := listInstance((serviceClient(createSession())))
	result, err := svc.TerminateInstances(&ec2.TerminateInstancesInput{
		InstanceIds: aws.StringSlice(lst),
	})
	if err != nil {
		fmt.Printf("Error terminating instance %s", err)
		return
	}

	fmt.Println("Instance Terminated.", result.String())
}

func createKeyValue(svc *ec2.EC2) *ec2.CreateKeyPairOutput {
	keyPair, err := svc.CreateKeyPair(&ec2.CreateKeyPairInput{
		KeyName: aws.String("Demo"),
	})
	if err != nil {
		fmt.Printf("Error creating keyPair %s", err)
		return nil
	}
	return keyPair
}

func createSecurityGroup(svc *ec2.EC2) *ec2.CreateSecurityGroupOutput {
	secGrp, err := svc.CreateSecurityGroup(&ec2.CreateSecurityGroupInput{
		Description: aws.String("Demo Security Group"),
		GroupName: aws.String("Demo"),
	})
	if err != nil {
		fmt.Printf("Error creating security group %s", err)
		return nil
	}
	return secGrp
}


func createInstanceUsingSecurityGroup() *ec2.Reservation {
	svc := serviceClient(createSession())
	parameters := &ec2.RunInstancesInput{
		ImageId:        aws.String("ami-42a2532b"),
		InstanceType:   aws.String("t2.micro"),
		MinCount:       aws.Int64(1),
		MaxCount:       aws.Int64(1),
		KeyName:        aws.String("Demo"),
		SecurityGroups: aws.StringSlice([]string{"Demo"}),
	}

	runResult, err := svc.RunInstances(parameters)

	if err != nil {
		fmt.Printf("Error creating instance %s", err)
		return nil
	}

	fmt.Println("Instance Successfully Created.")

	return runResult
}

func main() {
	svc := serviceClient(createSession())
	fmt.Println(createInstance())
	time.Sleep(10 * time.Second)
	listInstance(svc)
	terminateInstance(svc)
	fmt.Println(*createKeyValue(svc))
	fmt.Println(*createSecurityGroup(svc))
	fmt.Println(createInstanceUsingSecurityGroup())
}
