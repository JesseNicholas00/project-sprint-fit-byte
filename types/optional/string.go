package optional

import (
	"github.com/bytedance/sonic"
)

var parser = sonic.ConfigFastest

type OptionalStr struct {
	Defined bool
	V       *string
}

func (o *OptionalStr) UnmarshalJSON(data []byte) error {
	o.Defined = true
	return parser.Unmarshal(data, &o.V)
}
