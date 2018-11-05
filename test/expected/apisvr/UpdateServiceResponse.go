// Code generated by protoapi:go; DO NOT EDIT.

package apisvr

// UpdateServiceResponse
type UpdateServiceResponse struct {
	Service_id int    `json:"service_id"`
	Tags       []*Tag `json:"tags"`
	Desc       string `json:"desc"`
}

func (r *UpdateServiceResponse) GetService_id() int {
	if r == nil {
		var zeroVal int
		return zeroVal
	}
	return r.Service_id
}

func (r *UpdateServiceResponse) GetTags() []*Tag {
	if r == nil {
		var zeroVal []*Tag
		return zeroVal
	}
	return r.Tags
}

func (r *UpdateServiceResponse) GetDesc() string {
	if r == nil {
		var zeroVal string
		return zeroVal
	}
	return r.Desc
}
