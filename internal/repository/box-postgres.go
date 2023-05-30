package repository

import (
	"context"
	"encoding/csv"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
	"product-packaging/internal/core"
	"strings"
)

type PostgresHandler struct {
	conn *pgxpool.Pool
}

func NewPostgresHandler(config string) PostgresHandler {
	connString := config
	conn, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalln("Ошибка подключения к БД ", err)
	}
	return PostgresHandler{conn: conn}
}

func (p PostgresHandler) SaveBox(box core.Box, dataPacks string) *core.ErrorResponse {
	_, err := p.conn.Exec(context.Background(),
		"WITH x as(INSERT INTO boxes (pk_sscc, created_at) VALUES ($1,$2)) "+
			"INSERT INTO packs(pk_serial_number, fk_box_sscc, fk_product_gtin) VALUES "+
			dataPacks[:len(dataPacks)-1]+";",
		box.Sscc,
		box.CreatedAt)
	if err != nil {
		if strings.Contains(err.Error(), "packs_pkey") {
			return core.Error2
		}
		if strings.Contains(err.Error(), "boxes_pkey") {
			return core.Error1
		}
		return core.Error9
	}
	return nil
}

func (p PostgresHandler) GetBoxesBySgtin(serialNumbers string) (map[string]*string, *core.ErrorResponse) {
	rows, err := p.conn.Query(context.Background(),
		"SELECT fk_product_gtin || pk_serial_number as sgtin, fk_box_sscc FROM packs "+
			"WHERE pk_serial_number in ("+
			serialNumbers[:len(serialNumbers)-1]+");")
	if err != nil {
		return nil, core.Error9
	}
	defer rows.Close()
	var data = make(map[string]*string)
	for rows.Next() {
		var key string
		var value *string = nil
		err = rows.Scan(&key, &value)
		if err != nil {
			return nil, core.Error9
		}
		data[key] = value
	}
	return data, nil
}

func (p PostgresHandler) GetBoxesAndPacksByGtin(gtin string, csvFile *os.File) *core.ErrorResponse {
	rows, err1 := p.conn.Query(context.Background(),
		"SELECT fk_product_gtin || pk_serial_number as sgtin, fk_box_sscc FROM packs "+
			"WHERE fk_product_gtin = $1 and fk_box_sscc is not null;", gtin)
	if err1 != nil {
		log.Println("Ошибка извлечения данных из базы", err1)
		return core.Error9
	}
	defer rows.Close()

	w := csv.NewWriter(csvFile)
	for rows.Next() {
		var data []string
		var sgtin string
		var sscc string
		err2 := rows.Scan(&sgtin, &sscc)
		if err2 != nil {
			log.Println("Ошибка чтения строки", err2)
			return core.Error9
		}
		err3 := w.Write(append(data, sgtin, sscc))
		if err3 != nil {
			log.Println("Ошибка записи в csv файл", err2)
			return core.Error9
		}
	}
	w.Flush()
	return nil
}

func (p PostgresHandler) GetAmountPacksAndVerifyGtin(gtin string) (int, *core.ErrorResponse) {
	var amountPacks = 0
	rows, err := p.conn.Query(context.Background(),
		"SELECT amount_packs FROM products WHERE pk_gtin = $1", gtin)
	if err != nil {
		log.Println("Ошибка извлечения данных из базы", err)
		return amountPacks, core.Error9
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&amountPacks)
		if err != nil {
			log.Println("Ошибка чтения строки", err)
			return amountPacks, core.Error9
		}
	}
	if amountPacks == 0 {
		return amountPacks, core.Error5
	}
	return amountPacks, nil
}
