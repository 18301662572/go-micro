package prodmodel

import "strconv"

type ProModel struct {
	ProdID   int    `json:"pid"`
	ProdName string `json:"pname"`
}

func NewProd(id int, pname string) *ProModel {
	return &ProModel{
		ProdID:   id,
		ProdName: pname,
	}
}

func NewProdList(n int) []*ProModel {
	ret := make([]*ProModel, 0)
	for i := 0; i < n; i++ {
		ret = append(ret, NewProd(100+i, "pname"+strconv.Itoa(i)))
	}
	return ret
}
