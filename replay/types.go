package replay

type GetServiceCapabilities struct {
	XMLName string `xml:"trp:GetRecordingConfiguration"`
}

type GetServiceCapabilitiesResponse struct {
}

type GetReplayUri struct {
	XMLName string `xml:"trp:GetReplayUri"`
}

type GetReplayUriResponse struct {
}

type GetReplayConfiguration struct {
	XMLName string `xml:"trp:GetReplayConfiguration"`
}

type GetReplayConfigurationResponse struct {
}
