// Code generated by protoapi:go; DO NOT EDIT.

package apisvr

// KeyValueResponse
type KeyValueResponse struct {
	Key_values []*KeyValue `json:"key_values"`
}

func (r *KeyValueResponse) GetKey_values() []*KeyValue {
	if r == nil {
		var zeroVal []*KeyValue
		return zeroVal
	}
	return r.Key_values
}
