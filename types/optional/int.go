package optional

import (
	"github.com/bytedance/sonic"
)

type OptionalInt struct {
	Defined bool
	V       *int
}

func (o *OptionalInt) UnmarshalJSON(data []byte) error {
	o.Defined = true
	return sonic.ConfigFastest.Unmarshal(data, &o.V)
}
