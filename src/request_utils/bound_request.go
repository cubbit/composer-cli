package request_utils

type BoundRequest struct {
	boundModifiers []RequestModifier
}

func NewBoundRequest(mods ...RequestModifier) *BoundRequest {
	return &BoundRequest{boundModifiers: mods}
}

func (br *BoundRequest) Do(url string, opts ...RequestModifier) error {
	all := append(br.boundModifiers, opts...)
	return DoRequest(url, all...)
}
