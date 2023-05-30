package repository

import (
	"os"
	"product-packaging/internal/core"
)

//go:generate mockgen -source=repository.go -destination=\mock\mock-repository.go

type DBHandler interface {
	SaveBox(box core.Box, dataPacks string) *core.ErrorResponse
	GetBoxesBySgtin(serialNumbers string) (map[string]*string, *core.ErrorResponse)
	GetBoxesAndPacksByGtin(gtin string, csvFile *os.File) *core.ErrorResponse
	GetAmountPacksAndVerifyGtin(gtin string) (int, *core.ErrorResponse)
}
