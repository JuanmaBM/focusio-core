package entity

type FocusApp struct {
	Name             string `json:"name" bson:"name,omitempty"`
	Type             string `json:"type" bson:"type,omitempty"`
	SourceRepository string `json:"source_repository" bson:"source_repository,omitempty"`
	ConfigRepository string `json:"config_repository" bson:"config_repository,omitempty"`
	ImageRegistry    string `json:"image_registry" bson:"image_registry,omitempty"`
}

func NewFocusApp(n string, t string, sr string, cr string, ir string) *FocusApp {
	return &FocusApp{
		Name:             n,
		Type:             t,
		SourceRepository: sr,
		ConfigRepository: cr,
		ImageRegistry:    ir,
	}
}
