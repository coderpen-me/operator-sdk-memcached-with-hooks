package v1alpha1

type Plugin1 struct {
	Field1 struct {
		Field4 int `json:"field4"`
		Field5 struct {
			Field6 int `json:"field6"`
			Field7 int `json:"field7"`
		} `json:"field5"`
	} `json:"field1"`
	Field2 map[string]int `json:"field2"`
	Field3 map[string]int `json:"field3"`
}
