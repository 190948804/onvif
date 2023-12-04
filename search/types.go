package search

import (
	"github.com/190948804/onvif/record"
	"github.com/190948804/onvif/xsd"
	"github.com/190948804/onvif/xsd/onvif"
)

type GetRecordingSummary struct {
	XMLName string `xml:"tse:GetRecordingSummary"`
}

type GetRecordingSummaryResponse struct {
	Summary Summary `xml:"Summary"`
}

type Summary struct {
	DataFrom         string `xml:"DataFrom"`
	DataUntil        string `xml:"DataUntil"`
	NumberRecordings string `xml:"NumberRecordings"`
}

type GetRecordingInformation struct {
	XMLName        string               `xml:"tse:GetRecordingInformation"`
	RecordingToken onvif.ReferenceToken `xml:"RecordingToken"`
}

type GetRecordingInformationResponse struct {
	RecordingInformation RecordingInformation `xml:"RecordingInformation"`
}

type RecordingInformation struct {
	RecordingToken    string                      `xml:"RecordingToken"`
	Source            record.Source               `xml:"Source"`
	EarliestRecording string                      `xml:"EarliestRecording"`
	LatestRecording   string                      `xml:"LatestRecording"`
	Content           string                      `xml:"Content"`
	Track             []record.TrackConfiguration `xml:"Track"`
	RecordingStatus   string                      `xml:"RecordingStatus"`
}

type FindRecordings struct {
	XMLName        string         `xml:"tse:FindRecordings"`
	Scope          SearchScope    `xml:"Scope"`
	MetadataFilter MetadataFilter `xml:"MetadataFilter"`
	MaxMatches     int            `xml:"MaxMatches"`
	KeepAliveTime  xsd.Duration   `xml:"KeepAliveTime"`
}

type SearchScope struct {
	IncludedSources            []IncludedSource `xml:"IncludedSources"`
	IncludedRecordings         []string         `xml:"IncludedRecordings"`
	RecordingInformationFilter string           `xml:"RecordingInformationFilter"`
	Extension                  string           `xml:"Extension"`
}

type IncludedSource struct {
	Type  string `xml:"Type"`
	Token string `xml:"Token"`
}

type MetadataFilter struct {
	MetadataStreamFilter string `xml:"MetadataStreamFilter"`
}

type FindRecordingsResponse struct {
}
