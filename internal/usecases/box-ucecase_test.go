package usecases

import (
	"github.com/golang/mock/gomock"
	"product-packaging/internal/core"
	"product-packaging/internal/repository"
	mock_repository "product-packaging/internal/repository/mock"
	"testing"
	"time"
)

func Test_BoxUsecase_CreateBox(t *testing.T) {
	//Тестовые данные
	type mockBehavior func(u *mock_repository.MockDBHandler,
		box core.Box, gtin string, dataPacks string, amountPacks int)

	testTable := []struct {
		name                  string
		box                   core.Box
		gtin                  string
		dataPacks             string
		amountPacks           int
		mockBehavior          mockBehavior
		expectedErrorResponse *core.ErrorResponse
	}{
		{
			name: "Правильные данные",
			box: core.Box{Sscc: "000000000000000001",
				Packs: []string{"04603988000001IE9HALA4IBIH1",
					"04603988000001JE91ALA4H5J18",
					"04603988000001KE91ALA517K1J",
					"046039880000015E9FALA4L95F8"},
				CreatedAt: time.Now()},
			gtin: "04603988000001",
			dataPacks: "('IE9HALA4IBIH1','000000000000000001','04603988000001')," +
				"('JE91ALA4H5J18','000000000000000001','04603988000001')," +
				"('KE91ALA517K1J','000000000000000001','04603988000001')," +
				"('5E9FALA4L95F8','000000000000000001','04603988000001'),",
			amountPacks: 4,
			mockBehavior: func(u *mock_repository.MockDBHandler, box core.Box, gtin string, dataPacks string, amountPacks int) {
				u.EXPECT().GetAmountPacksAndVerifyGtin(gtin).Return(amountPacks, nil)
				u.EXPECT().SaveBox(box, dataPacks).Return(nil)
			},
			expectedErrorResponse: nil,
		},
		{
			name: "Ошибка в записи sscc - буква вместо цифры",
			box: core.Box{Sscc: "0000000000000000G1",
				Packs: []string{"04603988000001IE9HALA4IBIH1",
					"04603988000001JE91ALA4H5J18",
					"04603988000001KE91ALA517K1J",
					"046039880000015E9FALA4L95F8"},
				CreatedAt: time.Now()},
			gtin: "04603988000001",
			dataPacks: "('IE9HALA4IBIH1','0000000000000000G1','04603988000001')," +
				"('JE91ALA4H5J18','0000000000000000G1','04603988000001')," +
				"('KE91ALA517K1J','0000000000000000G1','04603988000001')," +
				"('5E9FALA4L95F8','0000000000000000G1','04603988000001'),",
			amountPacks: 4,
			mockBehavior: func(u *mock_repository.MockDBHandler, box core.Box, gtin string, dataPacks string, amountPacks int) {
				u.EXPECT().GetAmountPacksAndVerifyGtin(gtin).Return(amountPacks, nil).AnyTimes()
				u.EXPECT().SaveBox(box, dataPacks).Return(nil).AnyTimes()
			},
			expectedErrorResponse: core.Error6,
		},
		{
			name: "Ошибка в записи sscc - меньше цифр чем должно быть",
			box: core.Box{Sscc: "00000000000000001",
				Packs: []string{"04603988000001IE9HALA4IBIH1",
					"04603988000001JE91ALA4H5J18",
					"04603988000001KE91ALA517K1J",
					"046039880000015E9FALA4L95F8"},
				CreatedAt: time.Now()},
			gtin: "04603988000001",
			dataPacks: "('IE9HALA4IBIH1','00000000000000001','04603988000001')," +
				"('JE91ALA4H5J18','00000000000000001','04603988000001')," +
				"('KE91ALA517K1J','00000000000000001','04603988000001')," +
				"('5E9FALA4L95F8','00000000000000001','04603988000001'),",
			amountPacks: 4,
			mockBehavior: func(u *mock_repository.MockDBHandler, box core.Box, gtin string, dataPacks string, amountPacks int) {
				u.EXPECT().GetAmountPacksAndVerifyGtin(gtin).Return(amountPacks, nil).AnyTimes()
				u.EXPECT().SaveBox(box, dataPacks).Return(nil).AnyTimes()
			},
			expectedErrorResponse: core.Error6,
		},
		{
			name: "Ошибка в записи sscc - больше цифр чем должно быть",
			box: core.Box{Sscc: "0000000000000000001",
				Packs: []string{"04603988000001IE9HALA4IBIH1",
					"04603988000001JE91ALA4H5J18",
					"04603988000001KE91ALA517K1J",
					"046039880000015E9FALA4L95F8"},
				CreatedAt: time.Now()},
			gtin: "04603988000001",
			dataPacks: "('IE9HALA4IBIH1','0000000000000000001','04603988000001')," +
				"('JE91ALA4H5J18','0000000000000000001','04603988000001')," +
				"('KE91ALA517K1J','0000000000000000001','04603988000001')," +
				"('5E9FALA4L95F8','0000000000000000001','04603988000001'),",
			amountPacks: 4,
			mockBehavior: func(u *mock_repository.MockDBHandler, box core.Box, gtin string, dataPacks string, amountPacks int) {
				u.EXPECT().GetAmountPacksAndVerifyGtin(gtin).Return(amountPacks, nil).AnyTimes()
				u.EXPECT().SaveBox(box, dataPacks).Return(nil).AnyTimes()
			},
			expectedErrorResponse: core.Error6,
		},
		{
			name: "SSCC уже использован - записан в БД",
			box: core.Box{Sscc: "000000000000000002",
				Packs: []string{"04603988000001IE9HALA4IBIH1",
					"04603988000001JE91ALA4H5J18",
					"04603988000001KE91ALA517K1J",
					"046039880000015E9FALA4L95F8"},
				CreatedAt: time.Now()},
			gtin: "04603988000001",
			dataPacks: "('IE9HALA4IBIH1','000000000000000002','04603988000001')," +
				"('JE91ALA4H5J18','000000000000000002','04603988000001')," +
				"('KE91ALA517K1J','000000000000000002','04603988000001')," +
				"('5E9FALA4L95F8','000000000000000002','04603988000001'),",
			amountPacks: 4,
			mockBehavior: func(u *mock_repository.MockDBHandler, box core.Box, gtin string, dataPacks string, amountPacks int) {
				u.EXPECT().GetAmountPacksAndVerifyGtin(gtin).Return(amountPacks, nil).AnyTimes()
				u.EXPECT().SaveBox(box, dataPacks).Return(core.Error1).AnyTimes()
			},
			expectedErrorResponse: core.Error1,
		},
		{
			name: "Ошибка в записи первого gtin - буква вместо цифры",
			box: core.Box{Sscc: "000000000000000002",
				Packs: []string{"R4603988000001IE9HALA4IBIH1",
					"04603988000001JE91ALA4H5J18",
					"04603988000001KE91ALA517K1J",
					"046039880000015E9FALA4L95F8"},
				CreatedAt: time.Now()},
			gtin: "R4603988000001",
			dataPacks: "('IE9HALA4IBIH1','000000000000000002','R4603988000001')," +
				"('JE91ALA4H5J18','000000000000000002','04603988000001')," +
				"('KE91ALA517K1J','000000000000000002','04603988000001')," +
				"('5E9FALA4L95F8','000000000000000002','04603988000001'),",
			amountPacks: 4,
			mockBehavior: func(u *mock_repository.MockDBHandler, box core.Box, gtin string, dataPacks string, amountPacks int) {
				u.EXPECT().GetAmountPacksAndVerifyGtin(gtin).Return(amountPacks, nil).AnyTimes()
				u.EXPECT().SaveBox(box, dataPacks).Return(nil).AnyTimes()
			},
			expectedErrorResponse: core.Error7,
		},
		{
			name: "В пачке есть разные gtin",
			box: core.Box{Sscc: "000000000000000002",
				Packs: []string{"04603988000001IE9HALA4IBIH1",
					"F4603988000001JE91ALA4H5J18",
					"04603988000001KE91ALA517K1J",
					"046039880000015E9FALA4L95F8"},
				CreatedAt: time.Now()},
			gtin: "04603988000001",
			dataPacks: "('IE9HALA4IBIH1','000000000000000002','04603988000001')," +
				"('JE91ALA4H5J18','000000000000000002','F4603988000001')," +
				"('KE91ALA517K1J','000000000000000002','04603988000001')," +
				"('5E9FALA4L95F8','000000000000000002','04603988000001'),",
			amountPacks: 4,
			mockBehavior: func(u *mock_repository.MockDBHandler, box core.Box, gtin string, dataPacks string, amountPacks int) {
				u.EXPECT().GetAmountPacksAndVerifyGtin(gtin).Return(amountPacks, nil).AnyTimes()
				u.EXPECT().SaveBox(box, dataPacks).Return(nil).AnyTimes()
			},
			expectedErrorResponse: core.Error3,
		},
		{
			name: "Не соответствует количество пачек",
			box: core.Box{Sscc: "000000000000000002",
				Packs: []string{"04603988000001IE9HALA4IBIH1",
					"04603988000001JE91ALA4H5J18",
					"04603988000001KE91ALA517K1J"},
				CreatedAt: time.Now()},
			gtin: "04603988000001",
			dataPacks: "('IE9HALA4IBIH1','000000000000000002','04603988000001')," +
				"('JE91ALA4H5J18','000000000000000002','04603988000001')," +
				"('KE91ALA517K1J','000000000000000002','04603988000001'),",
			amountPacks: 4,
			mockBehavior: func(u *mock_repository.MockDBHandler, box core.Box, gtin string, dataPacks string, amountPacks int) {
				u.EXPECT().GetAmountPacksAndVerifyGtin(gtin).Return(amountPacks, nil).AnyTimes()
				u.EXPECT().SaveBox(box, dataPacks).Return(nil).AnyTimes()
			},
			expectedErrorResponse: core.Error4,
		},
		{
			name: "Ошибка в записи серийного номера пачки - не правильный символ",
			box: core.Box{Sscc: "000000000000000001",
				Packs: []string{"04603988000001IE9HALA4IBIH1",
					"04603988000001JE91ALA4H5J18",
					"04603988000001KE91ALA517K1J",
					"046039880000015E9FALA4L95Fл"},
				CreatedAt: time.Now()},
			gtin: "04603988000001",
			dataPacks: "('IE9HALA4IBIH1','000000000000000001','04603988000001')," +
				"('JE91ALA4H5J18','000000000000000001','04603988000001')," +
				"('KE91ALA517K1J','000000000000000001','04603988000001')," +
				"('5E9FALA4L95Fл','000000000000000001','04603988000001'),",
			amountPacks: 4,
			mockBehavior: func(u *mock_repository.MockDBHandler, box core.Box, gtin string, dataPacks string, amountPacks int) {
				u.EXPECT().GetAmountPacksAndVerifyGtin(gtin).Return(amountPacks, nil).AnyTimes()
				u.EXPECT().SaveBox(box, dataPacks).Return(nil).AnyTimes()
			},
			expectedErrorResponse: core.Error8,
		},
		{
			name: "Ошибка в записи серийного номера пачки - меньше символов чем должно быть",
			box: core.Box{Sscc: "000000000000000001",
				Packs: []string{"04603988000001IE9HALA4IBIH1",
					"04603988000001JE91ALA4H5J18",
					"04603988000001KE91ALA517K1J",
					"046039880000015E9FALA4L95"},
				CreatedAt: time.Now()},
			gtin: "04603988000001",
			dataPacks: "('IE9HALA4IBIH1','000000000000000001','04603988000001')," +
				"('JE91ALA4H5J18','000000000000000001','04603988000001')," +
				"('KE91ALA517K1J','000000000000000001','04603988000001')," +
				"('5E9FALA4L95','000000000000000001','04603988000001'),",
			amountPacks: 4,
			mockBehavior: func(u *mock_repository.MockDBHandler, box core.Box, gtin string, dataPacks string, amountPacks int) {
				u.EXPECT().GetAmountPacksAndVerifyGtin(gtin).Return(amountPacks, nil).AnyTimes()
				u.EXPECT().SaveBox(box, dataPacks).Return(nil).AnyTimes()
			},
			expectedErrorResponse: core.Error8,
		},
		{
			name: "Ошибка в записи серийного номера пачки - больше символов чем должно быть",
			box: core.Box{Sscc: "000000000000000001",
				Packs: []string{"04603988000001IE9HALA4IBIH1",
					"04603988000001JE91ALA4H5J18",
					"04603988000001KE91ALA517K1J",
					"046039880000015E9FALA4L95F8G"},
				CreatedAt: time.Now()},
			gtin: "04603988000001",
			dataPacks: "('IE9HALA4IBIH1','000000000000000001','04603988000001')," +
				"('JE91ALA4H5J18','000000000000000001','04603988000001')," +
				"('KE91ALA517K1J','000000000000000001','04603988000001')," +
				"('5E9FALA4L95F8G','000000000000000001','04603988000001'),",
			amountPacks: 4,
			mockBehavior: func(u *mock_repository.MockDBHandler, box core.Box, gtin string, dataPacks string, amountPacks int) {
				u.EXPECT().GetAmountPacksAndVerifyGtin(gtin).Return(amountPacks, nil).AnyTimes()
				u.EXPECT().SaveBox(box, dataPacks).Return(nil).AnyTimes()
			},
			expectedErrorResponse: core.Error8,
		},
		{
			name: "Одна и та же пачка в одной упаковке",
			box: core.Box{Sscc: "000000000000000001",
				Packs: []string{"04603988000001IE9HALA4IBIH1",
					"04603988000001JE91ALA4H5J18",
					"046039880000015E9FALA4L95F8",
					"046039880000015E9FALA4L95F8"},
				CreatedAt: time.Now()},
			gtin: "04603988000001",
			dataPacks: "('IE9HALA4IBIH1','000000000000000001','04603988000001')," +
				"('JE91ALA4H5J18','000000000000000001','04603988000001')," +
				"('5E9FALA4L95F8','000000000000000001','04603988000001')," +
				"('5E9FALA4L95F8','000000000000000001','04603988000001'),",
			amountPacks: 4,
			mockBehavior: func(u *mock_repository.MockDBHandler, box core.Box, gtin string, dataPacks string, amountPacks int) {
				u.EXPECT().GetAmountPacksAndVerifyGtin(gtin).Return(amountPacks, nil).AnyTimes()
				u.EXPECT().SaveBox(box, dataPacks).Return(nil).AnyTimes()
			},
			expectedErrorResponse: core.Error2,
		},
		{
			name: "Пачка уже использована в другой коробке",
			box: core.Box{Sscc: "000000000000000001",
				Packs: []string{"04603988000001IE9HALA4IBIH1",
					"04603988000001JE91ALA4H5J18",
					"046039880000015E9FALA4L95F8",
					"046039880000015E9FALA4L95F8"},
				CreatedAt: time.Now()},
			gtin: "04603988000001",
			dataPacks: "('IE9HALA4IBIH1','000000000000000001','04603988000001')," +
				"('JE91ALA4H5J18','000000000000000001','04603988000001')," +
				"('5E9FALA4L95F8','000000000000000001','04603988000001')," +
				"('5E9FALA4L95F8','000000000000000001','04603988000001'),",
			amountPacks: 4,
			mockBehavior: func(u *mock_repository.MockDBHandler, box core.Box, gtin string, dataPacks string, amountPacks int) {
				u.EXPECT().GetAmountPacksAndVerifyGtin(gtin).Return(amountPacks, nil).AnyTimes()
				u.EXPECT().SaveBox(box, dataPacks).Return(core.Error2).AnyTimes()
			},
			expectedErrorResponse: core.Error2,
		},
		{
			name: "Пачка уже использована в другой коробке",
			box: core.Box{Sscc: "000000000000000001",
				Packs: []string{"04603988000001IE9HALA4IBIH1",
					"04603988000001JE91ALA4H5J18",
					"046039880000015E9FALA4L95F8",
					"046039880000015E9FALA4L95F8"},
				CreatedAt: time.Now()},
			gtin: "04603988000001",
			dataPacks: "('IE9HALA4IBIH1','000000000000000001','04603988000001')," +
				"('JE91ALA4H5J18','000000000000000001','04603988000001')," +
				"('5E9FALA4L95F8','000000000000000001','04603988000001')," +
				"('5E9FALA4L95F8','000000000000000001','04603988000001'),",
			amountPacks: 4,
			mockBehavior: func(u *mock_repository.MockDBHandler, box core.Box, gtin string, dataPacks string, amountPacks int) {
				u.EXPECT().GetAmountPacksAndVerifyGtin(gtin).Return(amountPacks, nil).AnyTimes()
				u.EXPECT().SaveBox(box, dataPacks).Return(core.Error2).AnyTimes()
			},
			expectedErrorResponse: core.Error2,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Внедрение зависимостей

			ctr := gomock.NewController(t)
			dbhandler := mock_repository.NewMockDBHandler(ctr)
			testCase.mockBehavior(dbhandler, testCase.box, testCase.gtin, testCase.dataPacks, testCase.amountPacks)
			boxRepo := repository.NewBoxRepo(dbhandler)
			boxUsecase := NewBox(boxRepo)

			// Тестирование функции
			result := boxUsecase.CreateBox(testCase.box)

			// Сравнение результатов
			t.Log("Ожидалось  - ", testCase.expectedErrorResponse)
			t.Log("Получилось - ", result)
			if testCase.expectedErrorResponse != nil {
				if testCase.expectedErrorResponse != result {
					t.Errorf("Ошибка теста")
				}
			}
		})

	}
}
