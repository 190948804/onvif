package record

import (
	"github.com/190948804/onvif/xsd/onvif"
)

type GetRecordingConfiguration struct {
	XMLName        string               `xml:"trc:GetRecordingConfiguration"`
	RecordingToken onvif.ReferenceToken `xml:"RecordingToken"`
}

type GetRecordingConfigurationResponse struct {
	RecordingConfiguration RecordingConfiguration `xml:"RecordingConfiguration"`
}

type RecordingConfiguration struct {
	Source               Source `xml:"Source"`
	Content              string `xml:"Content"`
	MaximumRetentionTime string `xml:"MaximumRetentionTime"`
}

type GetRecordings struct {
	XMLName string `xml:"trc:GetRecordings"`
}

type GetRecordingsResponse struct {
	RecordingItem []RecordingItem `xml:"RecordingItem"`
}

type RecordingItem struct {
	RecordingToken string                 `xml:"RecordingToken"`
	Configuration  RecordingConfiguration `xml:"Configuration"`
}

type Source struct {
	SourceId    string `xml:"SourceId"`
	Name        string `xml:"Name"`
	Location    string `xml:"Location"`
	Description string `xml:"Description"`
	Address     string `xml:"Address"`
}

type Tracks struct {
	Track []Track `xml:"Track"`
}

type Track struct {
	TrackToken    string             `xml:"RecordingToken"`
	Configuration TrackConfiguration `xml:"Configuration"`
}

type TrackConfiguration struct {
	TrackType   string `xml:"TrackType"`
	Description string `xml:"Description"`

	TrackToken *string `xml:"TrackToken"`
	DataFrom   *string `xml:"DataFrom"`
	DataTo     *string `xml:"DataTo"`
}

type GetRecordingJobs struct {
	XMLName string `xml:"trc:GetRecordingJobs"`
}

type GetRecordingJobsResponse struct {
	JobItem []JobItem `xml:"JobItem"`
}

type JobItem struct {
	JobToken         string           `xml:"JobToken"`
	JobConfiguration JobConfiguration `xml:"JobConfiguration"`
}

type JobConfiguration struct {
	RecordingToken string                 `xml:"RecordingToken"`
	Mode           string                 `xml:"Mode"`
	Priority       int                    `xml:"Priority"`
	Source         JobConfigurationSource `xml:"Source"`
}

type JobConfigurationSource struct {
	SourceToken SourceToken `xml:"SourceToken"`
	Tracks      []JobTrack  `xml:"Tracks"`
}

type SourceToken struct {
	Token string `xml:"Token"`
}

type JobTrack struct {
	SourceTag   string `xml:"SourceTag"`
	Destination string `xml:"Destination"`
}

type GetRecordingJobState struct {
	XMLName string `xml:"trc:GetRecordingJobState"`
}

type GetRecordingJobStateResponse struct {
	XMLName string `xml:"trc:GetRecordingJobStateResponse"`
}
