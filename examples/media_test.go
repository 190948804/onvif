package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"

	goonvif "github.com/190948804/onvif"
	"github.com/190948804/onvif/media"
	sdk_media "github.com/190948804/onvif/sdk/media"
	"github.com/190948804/onvif/xsd/onvif"
)

func TestMediaService(t *testing.T) {
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

	getProfiles := media.GetProfiles{}
	getProfilesResponse, err := sdk_media.Call_GetProfiles(ctx, dev, getProfiles)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getProfilesResponse)
	}

	for x := 0; x < len(getProfilesResponse.Profiles); x++ {
		osds := media.GetOSDs{ConfigurationToken: onvif.ReferenceToken(getProfilesResponse.Profiles[x].VideoSourceConfiguration.ConfigurationEntity.Token)}
		getOSDsResponse, err := sdk_media.Call_GetOSDs(ctx, dev, osds)
		if err != nil {
			continue
		}
		if len(getOSDsResponse.OSDs) == 0 {
			continue
		}
		for _, osd := range getOSDsResponse.OSDs {
			if osd.TextString.Type == "Plain" && osd.TextString.PlainText != "" {
				fmt.Println(osd.TextString.PlainText)
			}
		}
	}
}
