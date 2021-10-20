package v1alpha1

type Plugin1 struct {
	Field1 Field1         `json:"field1"`
	Field2 map[string]int `json:"field2"`
	Field3 map[string]int `json:"field3"`
}

type Field1 struct {
	Field4 map[string]int `json:"field4"`
	Field5 map[string]int `json:"field5"`
}
