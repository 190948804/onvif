package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"

	goonvif "github.com/190948804/onvif"
	"github.com/190948804/onvif/replay"
	sdk_replay "github.com/190948804/onvif/sdk/replay"
)

func TestReplayConfiguration(t *testing.T) {
	ctx := context.Background()

	//Getting an camera instance
	dev, err := goonvif.NewDevice(goonvif.DeviceParams{
		Xaddr:      "192.168.15.117:80",
		Username:   "admin",
		Password:   "admin123",
		HttpClient: new(http.Client),
	})
	if err != nil {
		panic(err)
	}

	getServiceCapabilities := replay.GetServiceCapabilities{}
	getServiceCapabilitiesResponse, err := sdk_replay.Call_GetServiceCapabilities(ctx, dev, getServiceCapabilities)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getServiceCapabilitiesResponse)
	}
}
