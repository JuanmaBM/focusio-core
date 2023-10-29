package domain

type FocusCatalogItem struct {
	Name        string `json:"name" bson:"name,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
	Repository  string `json:"repository" bson:"repository,omitempty"`
}

func NewFocusCatalogItem(n string, d string, r string) *FocusCatalogItem {
	return &FocusCatalogItem{
		Name:        n,
		Description: d,
		Repository:  r,
	}
}
