package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"

	goonvif "github.com/190948804/onvif"
	"github.com/190948804/onvif/record"
	sdk_record "github.com/190948804/onvif/sdk/record"
	sdk_search "github.com/190948804/onvif/sdk/search"
	"github.com/190948804/onvif/search"
	"github.com/190948804/onvif/xsd/onvif"
)

func TestSearchConfiguration(t *testing.T) {
	ctx := context.Background()

	//Getting an camera instance
	dev, err := goonvif.NewDevice(goonvif.DeviceParams{
		Xaddr:      "192.168.58.2:80",
		Username:   "admin",
		Password:   "admin123",
		HttpClient: new(http.Client),
	})
	if err != nil {
		panic(err)
	}

	getRecordingSummary := search.GetRecordingSummary{}
	getRecordingSummaryResponse, err := sdk_search.Call_GetRecordingSummary(ctx, dev, getRecordingSummary)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getRecordingSummaryResponse)
	}

	getRecordings := record.GetRecordings{}
	getRecordingsResponse, err := sdk_record.Call_GetRecordings(ctx, dev, getRecordings)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getRecordingsResponse)
	}

	if len(getRecordingsResponse.RecordingItem) == 0 {
		fmt.Println("无录像数据")
		return
	}

	getRecordingInformation := search.GetRecordingInformation{RecordingToken: onvif.ReferenceToken(getRecordingsResponse.RecordingItem[0].RecordingToken)}
	getRecordingInformationResponse, err := sdk_search.Call_GetRecordingInformation(ctx, dev, getRecordingInformation)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getRecordingInformationResponse)
	}

	findRecordings := search.FindRecordings{}
	findRecordings.KeepAliveTime = "PT60S"
	findRecordingsResponse, err := sdk_search.Call_FindRecordings(ctx, dev, findRecordings)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(findRecordingsResponse)
	}

}
