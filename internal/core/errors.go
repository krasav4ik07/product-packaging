package core

type ErrorResponse struct {
	IsOk      bool   `json:"ok"`
	Err       string `json:"error"`
	ErrorCode int    `json:"error_code"`
}

//Не может быть нескольких коробок с одинаковым SSCC
var Error1 = &ErrorResponse{
	IsOk:      false,
	Err:       "SSCC уже использован",
	ErrorCode: 1}

//Одна и та же пачка может быть агрегирована только в одну коробку
var Error2 = &ErrorResponse{
	IsOk:      false,
	Err:       "пачка уже использована",
	ErrorCode: 2}

//В одну коробку можно упаковывать только пачки с одинаковым GTIN
var Error3 = &ErrorResponse{
	IsOk:      false,
	Err:       "пачки с разным GTIN",
	ErrorCode: 3}

//В одну коробку можно упаковать только N штук пачек, где N задается в справочнике продуктов
var Error4 = &ErrorResponse{
	IsOk:      false,
	Err:       "не соответствует количество пачек",
	ErrorCode: 4}

//Нельзя агрегировать пачки с неизвестным GTIN
var Error5 = &ErrorResponse{
	IsOk:      false,
	Err:       "неизвестный GTIN",
	ErrorCode: 5}

var Error6 = &ErrorResponse{
	IsOk:      false,
	Err:       "ошибка в записи SSCC",
	ErrorCode: 6}

var Error7 = &ErrorResponse{
	IsOk:      false,
	Err:       "ошибка в записи GTIN",
	ErrorCode: 7}

var Error8 = &ErrorResponse{
	IsOk:      false,
	Err:       "ошибка в записи серийного номера пачки",
	ErrorCode: 8}

var Error9 = &ErrorResponse{
	IsOk:      false,
	Err:       "проблема с данными или БД",
	ErrorCode: 9}
