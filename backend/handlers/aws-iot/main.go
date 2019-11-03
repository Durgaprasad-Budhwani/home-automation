package main

import (
	"fmt"

	"github.com/durgaprasad-budhwani/home-automation/backend/utils"

	"github.com/durgaprasad-budhwani/aws-iot-device-sdk-go/device"
)

func main() {
	viper := utils.NewViper("./config.yaml")
	fmt.Print(viper.GetString("Database"))

	thing, err := device.NewThing(
		device.KeyPair{
			PrivateKeyPath:    viper.GetString("PrivateKeyPath"),
			CertificatePath:   viper.GetString("CertificatePath"),
			CACertificatePath: viper.GetString("CACertificatePath"),
		},
		viper.GetString("AwsIotEndpoint"), // AWS IoT endpoint
		device.ThingName("AwsIotThinkName"),
	)
	if err != nil {
		panic(err)
	}

	shadowChan, err := thing.SubscribeForThingShadowChanges(viper.GetString("AwsIotTopicName"))
	if err != nil {
		panic(err)
	}

	for {
		select {
		case s, ok := <-shadowChan:
			if !ok {
				panic("failed to read from shadow channel")
			}
			fmt.Println(string(s))
		}
	}
}
