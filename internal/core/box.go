package core

import (
	"os"
	"time"
)

const (
	Sscclength         = 18
	Gtinlength         = 14
	SerialNumberlength = 13
	Strlength          = Sscclength + Gtinlength + SerialNumberlength + 11
)

type Box struct {
	Sscc      string    `json:"sscc"`
	Packs     []string  `json:"sgtins"`
	CreatedAt time.Time `json:"created"`
}

type BoxRepository interface {
	SaveBox(box Box, dataPacks string) *ErrorResponse
	GetBoxesBySgtin(serialNumbers string) (map[string]*string, *ErrorResponse)
	GetBoxesAndPacksByGtin(gtin string, csvFile *os.File) *ErrorResponse
	GetAmountPacksAndVerifyGtin(gtin string) (int, *ErrorResponse)
}
