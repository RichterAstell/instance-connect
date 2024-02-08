package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func main() {
	fmt.Println("list ec2 instance")
	list, err := listInstance()
	if err != nil {
		panic("instance list is nil")
	}

	// print header
	fmt.Println(" Name | Public IP | Instance ID ")
	fmt.Println("------+-----------+-------------")
	for _, info := range list {
		fmt.Println(info.name + " | " + info.publicIp + " | " + info.id)
	}
}

func listInstance() ([]InstanceInfo, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := ec2.NewFromConfig(cfg)

	output, err := client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("=== Describe instances: ")
	infoList := make([]InstanceInfo, 0, len(output.Reservations))

	for _, obj := range output.Reservations {
		// 停止中のインスタンスは無視する
		if obj.Instances[0].State.Name != "running" {
			continue
		}

		ins := InstanceInfo{
			id:       string(*obj.Instances[0].InstanceId),
			publicIp: string(*obj.Instances[0].PublicIpAddress),
		}

		for _, tag := range obj.Instances[0].Tags {
			if *tag.Key == "Name" {
				ins.name = *tag.Value
			}
		}
		log.Printf("%v", ins)
		infoList = append(infoList, ins)
	}
	return infoList, nil
}
