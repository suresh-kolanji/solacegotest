package main

import (
	"fmt"

	"solace.dev/go/messaging"
	"solace.dev/go/messaging/pkg/solace/config"
)

func main() {

	brokerConfig := config.ServicePropertyMap{
		config.TransportLayerPropertyHost:                "tcps://********:55443",
		config.ServicePropertyVPNName:                    "********",
		config.AuthenticationPropertySchemeBasicPassword: "******",
		config.AuthenticationPropertySchemeBasicUserName: "********",
	}

	// Skip certificate validation
	messagingService, err := messaging.NewMessagingServiceBuilder().
		FromConfigurationProvider(brokerConfig).
		WithTransportSecurityStrategy(config.NewTransportSecurityStrategy().WithoutCertificateValidation()).
		Build()

	if err != nil {
		panic(err)
	}

	if err = messagingService.Connect(); err != nil {
		panic(err)
	}

	fmt.Print("conneted")

	if err = messagingService.Disconnect(); err != nil {
		panic(err)
	}

	fmt.Print("Dis connected ")
}
