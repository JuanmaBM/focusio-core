package entity

type FocusApp struct {
	Name             string `json:"name" bson:"name,omitempty"`
	Type             string `json:"type" bson:"type,omitempty"`
	SourceRepository string `json:"source_repository" bson:"source_repository,omitempty"`
	ConfigRepository string `json:"config_repository" bson:"config_repository,omitempty"`
	ImageRegistry    string `json:"image_registry" bson:"image_registry,omitempty"`
	ArtifactId       string `json:"artifact_id" bson:"artifact_id"`
	GroupId          string `json:"group_id" bson:"group_id"`
	Package          string `json:"package" bson:"package"`
	FocusCatalogItem string `json:"focus_catalog_item" bson:"focus_catalog_item"`
}

func NewFocusApp(n string, t string, sr string, cr string, ir string, ai string, gi string, pkg string, fci string) *FocusApp {
	return &FocusApp{
		Name:             n,
		Type:             t,
		SourceRepository: sr,
		ConfigRepository: cr,
		ImageRegistry:    ir,
		ArtifactId:       ai,
		GroupId:          gi,
		Package:          pkg,
		FocusCatalogItem: fci,
	}
}
