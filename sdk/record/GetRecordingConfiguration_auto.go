// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package record

import (
	"context"
	"github.com/juju/errors"
	"github.com/190948804/onvif"
	"github.com/190948804/onvif/sdk"
	"github.com/190948804/onvif/record"
)

// Call_GetRecordingConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetRecordingConfigurationResponse.
func Call_GetRecordingConfiguration(ctx context.Context, dev *onvif.Device, request record.GetRecordingConfiguration) (record.GetRecordingConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRecordingConfigurationResponse record.GetRecordingConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetRecordingConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetRecordingConfiguration")
		return reply.Body.GetRecordingConfigurationResponse, errors.Annotate(err, "reply")
	}
}
