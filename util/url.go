package util

import (
	"net/url"

	"github.com/youkale/go-querystruct/params"
)

// URLValuesToStruct 把url.Values转成struct
func URLValuesToStruct(value url.Values, out interface{}) error {
	err := params.Unmarshal(value, out)
	return err
}
