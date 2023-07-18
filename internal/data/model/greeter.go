package model

type WideCut struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (WideCut) TableName() string {
	return "wide_cut"
}
