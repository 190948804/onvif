package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"

	goonvif "github.com/190948804/onvif"
	"github.com/190948804/onvif/device"
	"github.com/190948804/onvif/record"
	sdk_device "github.com/190948804/onvif/sdk/device"
	sdk_record "github.com/190948804/onvif/sdk/record"
	"github.com/190948804/onvif/xsd/onvif"
)

func TestRecordingConfiguration(t *testing.T) {
	ctx := context.Background()

	//Getting an camera instance
	dev, err := goonvif.NewDevice(goonvif.DeviceParams{
		Xaddr:      "192.168.15.137:80",
		Username:   login,
		Password:   password,
		HttpClient: new(http.Client),
	})
	if err != nil {
		panic(err)
	}

	getCapabilities := device.GetCapabilities{Category: "All"}
	getCapabilitiesResponse, err := sdk_device.Call_GetCapabilities(ctx, dev, getCapabilities)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getCapabilitiesResponse)
	}

	getRecordings := record.GetRecordings{}
	getRecordingsResponse, err := sdk_record.Call_GetRecordings(ctx, dev, getRecordings)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getRecordingsResponse)
	}

	getRecordingConfiguration := record.GetRecordingConfiguration{RecordingToken: onvif.ReferenceToken(getRecordingsResponse.RecordingItem[0].RecordingToken)}
	getRecordingConfigurationResponse, err := sdk_record.Call_GetRecordingConfiguration(ctx, dev, getRecordingConfiguration)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getRecordingConfigurationResponse)
	}

	getRecordingJobs := record.GetRecordingJobs{}
	getRecordingJobsResponse, err := sdk_record.Call_GetRecordingJobs(ctx, dev, getRecordingJobs)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getRecordingJobsResponse)
	}
}
