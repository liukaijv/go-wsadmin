package common

type Pager struct {
	Size    int `json:"size"`
	Total   int `json:"total"`
	Current int `json:"current"`
}

func New(current, size, total int) *Pager {
	return &Pager{
		Current: current,
		Size:    size,
		Total:   total,
	}
}
