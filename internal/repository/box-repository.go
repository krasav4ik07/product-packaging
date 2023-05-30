package repository

import (
	"os"
	"product-packaging/internal/core"
)

type BoxRepo struct {
	handler DBHandler
}

func NewBoxRepo(handler DBHandler) BoxRepo {
	return BoxRepo{handler: handler}
}

func (repo BoxRepo) SaveBox(box core.Box, dataPacks string) *core.ErrorResponse {
	err := repo.handler.SaveBox(box, dataPacks)
	if err != nil {
		return err
	}
	return nil
}

func (repo BoxRepo) GetBoxesBySgtin(serialNumbers string) (map[string]*string, *core.ErrorResponse) {
	boxes, err := repo.handler.GetBoxesBySgtin(serialNumbers)
	if err != nil {
		return nil, err
	}
	return boxes, err
}

func (repo BoxRepo) GetBoxesAndPacksByGtin(gtin string, csvFile *os.File) *core.ErrorResponse {
	err := repo.handler.GetBoxesAndPacksByGtin(gtin, csvFile)
	if err != nil {
		return err
	}
	return nil
}

func (repo BoxRepo) GetAmountPacksAndVerifyGtin(gtin string) (int, *core.ErrorResponse) {
	amountPacks, err := repo.handler.GetAmountPacksAndVerifyGtin(gtin)
	if err != nil {
		return amountPacks, err
	}
	return amountPacks, nil
}
