package controllers

import (
	"encoding/json"
	"net/http"
	"product-packaging/internal/core"
	"product-packaging/internal/usecases"
)

type gtinRequest struct {
	Gtin string `json:"gtin"`
}

type okResponse struct {
	IsOk bool `json:"ok"`
}

var ok = okResponse{IsOk: true}

type BoxController struct {
	box *usecases.BoxUsecases
}

func NewBoxController(b *usecases.BoxUsecases) *BoxController {
	return &BoxController{box: b}
}

func (b *BoxController) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var box core.Box
	err1 := json.NewDecoder(req.Body).Decode(&box)
	if err1 != nil {
		json.NewEncoder(res).Encode(err1.Error())
		return
	}
	err2 := b.box.CreateBox(box)
	if err2 != nil {
		json.NewEncoder(res).Encode(err2)
	} else {
		json.NewEncoder(res).Encode(ok)
	}
}

func (b *BoxController) GetBoxesBySgtin(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	box := core.Box{}
	err1 := json.NewDecoder(req.Body).Decode(&box)
	if err1 != nil {
		json.NewEncoder(res).Encode(err1.Error())
		return
	}
	data, err2 := b.box.GetBoxesBySgtin(box.Packs)
	if err2 != nil {
		json.NewEncoder(res).Encode(err2)
	} else {
		json.NewEncoder(res).Encode(data)
	}
}

func (b *BoxController) GetBoxesAndPacksByGtin(res http.ResponseWriter, req *http.Request) {

	data := gtinRequest{}
	err1 := json.NewDecoder(req.Body).Decode(&data)
	if err1 != nil {
		json.NewEncoder(res).Encode(err1.Error())
		return
	}
	csvFileByte, err2 := b.box.GetBoxesAndPacksByGtin(data.Gtin)
	res.Header().Set("Content-Disposition", "attachment; filename=\"test.csv\"")
	res.Header().Set("Content-Type", "text/csv")
	if err2 != nil {
		json.NewEncoder(res).Encode(err2)
	} else {
		res.Write(*csvFileByte)
	}
}
