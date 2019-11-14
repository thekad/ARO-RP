package v20191231preview

import (
	"encoding/json"
)

// UnmarshalJSON unmarshals tags.  We override this to ensure that PATCH
// behaviour overwrites an existing tags map rather than endlessly adding to it
func (t *Tags) UnmarshalJSON(b []byte) error {
	var m map[string]string
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	*t = m
	return nil
}