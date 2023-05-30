package usecases

import (
	"io/ioutil"
	"os"
	"product-packaging/internal/core"
	"regexp"
	"strings"
	"time"
)

type BoxUsecases struct {
	repository core.BoxRepository
}

func NewBox(repo core.BoxRepository) *BoxUsecases {
	return &BoxUsecases{repository: repo}
}

func (b *BoxUsecases) CreateBox(box core.Box) *core.ErrorResponse {
	err1 := b.verifySscc(box.Sscc)
	if err1 != nil {
		return err1
	}
	gtin := box.Packs[0][:14]
	amountPacks, err2 := b.verifyGtin(gtin)
	if err2 != nil {
		return err2
	}
	dataPacks, err3 := b.verifyPacks(gtin, amountPacks, box)
	if err3 != nil {
		return err3
	}
	err4 := b.repository.SaveBox(box, dataPacks)
	if err4 != nil {
		return err4
	}
	return nil
}

func (b *BoxUsecases) GetBoxesBySgtin(stgins []string) (map[string]*string, *core.ErrorResponse) {
	var serialNumbers strings.Builder
	serialNumbers.Grow((core.SerialNumberlength + 3) * len(stgins))
	for _, sgtin := range stgins {
		serialNumbers.WriteRune('\'')
		serialNumbers.WriteString(sgtin[14:])
		serialNumbers.WriteRune('\'')
		serialNumbers.WriteRune(',')
	}
	data, err := b.repository.GetBoxesBySgtin(serialNumbers.String())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (b *BoxUsecases) GetBoxesAndPacksByGtin(gtin string) (*[]byte, *core.ErrorResponse) {
	fileName := strings.ReplaceAll(time.Now().Format(time.RFC3339), ":", "-")
	csvFile, _ := os.Create("./data/csv/" + fileName + ".csv")
	defer csvFile.Close()
	err := b.repository.GetBoxesAndPacksByGtin(gtin, csvFile)
	if err != nil {
		return nil, err
	}
	csvFileByte, _ := ioutil.ReadFile(csvFile.Name())

	return &csvFileByte, nil
}

func (b *BoxUsecases) verifySscc(sscc string) *core.ErrorResponse {
	matched, err := regexp.MatchString(`\d{18}`, sscc)
	if err != nil {
		return core.Error6
	}
	if !(matched && len(sscc) == 18) {
		return core.Error6
	}
	return nil
}

func (b *BoxUsecases) verifyGtin(gtin string) (int, *core.ErrorResponse) {
	matched, err1 := regexp.MatchString(`\d{14}`, gtin)
	if err1 != nil {
		return 0, core.Error7
	}
	if !(matched && len(gtin) == 14) {
		return 0, core.Error7
	}
	amountPacks, err2 := b.repository.GetAmountPacksAndVerifyGtin(gtin)
	if err2 != nil {
		return amountPacks, err2
	}
	return amountPacks, nil
}

func (b *BoxUsecases) verifyPacks(trueGtin string, trueAmountPack int, box core.Box) (string, *core.ErrorResponse) {
	if !(trueAmountPack == len(box.Packs)) {
		return "", core.Error4
	}

	var dataPacks strings.Builder
	AmountBytes := trueAmountPack * core.Strlength
	dataPacks.Grow(AmountBytes)

	serialNumbers := make(map[string]string)

	for _, sgtin := range box.Packs {
		if !(trueGtin == sgtin[:14]) {
			return "", core.Error3
		}
		matched, err := regexp.MatchString(`[\da-zA-Z]+`, sgtin[14:])
		if err != nil {
			return "", core.Error8
		}
		if !(matched && len(sgtin[14:]) == 13) {
			return "", core.Error8
		}
		if _, ok := serialNumbers[sgtin[14:]]; ok {
			return "", core.Error2
		} else {
			serialNumbers[sgtin[14:]] = trueGtin
		}
		addRowInData(box.Sscc, sgtin, &dataPacks)
	}
	return dataPacks.String(), nil
}

func addRowInData(sscc string, sgtin string, data *strings.Builder) {
	data.WriteRune('(')
	data.WriteRune('\'')
	data.WriteString(sgtin[14:])
	data.WriteRune('\'')
	data.WriteRune(',')
	data.WriteRune('\'')
	data.WriteString(sscc)
	data.WriteRune('\'')
	data.WriteRune(',')
	data.WriteRune('\'')
	data.WriteString(sgtin[:14])
	data.WriteRune('\'')
	data.WriteRune(')')
	data.WriteRune(',')
}
