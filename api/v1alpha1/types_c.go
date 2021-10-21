package v1alpha1

type ValidatorResponse struct {
	Status      bool              `json:"status"`
	Data        DataStruct        `json:"data"`
	PluginError map[string]string `json:"plugin_error"`
}

type DataStruct struct {
	// MismatchedFields is a list of fields which have a type mismatch.
	MismatchedFields []MismatchedField

	// MissingFields is a list of JSON field names which were not in src.
	MissingFields []MissingField
}

type MismatchedField struct {
	// Field is the JSON name of the field.
	Field string

	// Expected is the expected type (the type of the dst field).
	Expected string

	// Actual is the actual type (type of the src field).
	Actual string

	// Path is the full path to the field.
	Path []string
}

type MissingField struct {
	// Field is the JSON name of the field.
	Field string

	// Path is the full path to the field.
	Path []string
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

func (in *DataStruct) DeepCopyInto(out *DataStruct) {
	*out = *in

	if in.MismatchedFields != nil {
		in, out := &in.MismatchedFields, &out.MismatchedFields
		*out = make([]MismatchedField, len(*in))
		copy(*out, *in)
	}

	if in.MissingFields != nil {
		in, out := &in.MissingFields, &out.MissingFields
		*out = make([]MissingField, len(*in))
		copy(*out, *in)
	}
}
