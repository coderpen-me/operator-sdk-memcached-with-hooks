package v1alpha1

import schema "github.com/Kangaroux/go-map-schema"

type ValidatorResponse struct {
	status      bool                  `json:"status"`
	Data        schema.CompareResults `json:"data"`
	PluginError map[string]string     `json:"plugin_error"`
}

type Plugin1 struct {
	Plugin1Name struct {
		Field1 int `json:"field1"`
		Field2 struct {
			Field3 int `json:"field3"`
			Field4 int `json:"field4"`
		} `json:"field2"`
	} `json:"plugin_1_name"`
}

type Plugin2 struct {
	Plugin2Name struct {
		Field1 int `json:"field1"`
		Field2 struct {
			Field3 int `json:"field3"`
			Field4 int `json:"field4"`
		} `json:"field2"`
	} `json:"plugin_2_name"`
}
