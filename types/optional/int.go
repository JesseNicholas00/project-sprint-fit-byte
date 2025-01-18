package optional

import "encoding/json"

type OptionalInt struct {
	Defined bool
	V       *int
}

func (o *OptionalInt) UnmarshalJSON(data []byte) error {
	o.Defined = true
	return json.Unmarshal(data, &o.V)
}
