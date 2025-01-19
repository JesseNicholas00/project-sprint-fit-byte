package optional

import (
	"github.com/bytedance/sonic"
)

type OptionalStr struct {
	Defined bool
	V       *string
}

func (o *OptionalStr) UnmarshalJSON(data []byte) error {
	o.Defined = true
	return sonic.ConfigFastest.Unmarshal(data, &o.V)
}
