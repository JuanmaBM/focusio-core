package domain

type FocusApp struct {
	Name string `json:"name" bson:"name,omitempty"`
}

func NewFocusApp(name string) *FocusApp {
	return &FocusApp{Name: name}
}
