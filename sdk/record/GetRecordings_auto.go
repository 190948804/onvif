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

// Call_GetRecordings forwards the call to dev.CallMethod() then parses the payload of the reply as a GetRecordingsResponse.
func Call_GetRecordings(ctx context.Context, dev *onvif.Device, request record.GetRecordings) (record.GetRecordingsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRecordingsResponse record.GetRecordingsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetRecordingsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetRecordings")
		return reply.Body.GetRecordingsResponse, errors.Annotate(err, "reply")
	}
}
