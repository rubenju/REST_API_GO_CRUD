package model

import (
	"database/sql"
	"standarisasigoapi/entity"

	_ "github.com/lib/pq"
)

func GetBarang(db *sql.DB) ([]entity.Barang, error) {
	readsql := "SELECT * from read_barang()"
	rows, error := db.Query(readsql)

	if error != nil {
		return []entity.Barang{}, error
	}
	defer rows.Close()

	hasil := []entity.Barang{}
	for rows.Next() {
		data_barang := entity.Barang{}
		error := rows.Scan(&data_barang.Id, &data_barang.Nama, &data_barang.Harga, &data_barang.Stok)
		if error != nil {
			return []entity.Barang{}, error
		}
		hasil = append(hasil, data_barang)
	}

	return hasil, error
}

func AddBarang(db *sql.DB, data entity.Barang) (bool, error) {
	insertsql := "SELECT create_barang($1, $2, $3, $4)"

	_, error := db.Query(insertsql, data.Id, data.Nama, data.Harga, data.Stok)

	if error != nil {
		return false, error
	}

	return true, error
}

func UpdateBarang(db *sql.DB, data entity.Barang) (bool, error) {
	updatesql := "SELECT update_barang($1, $2, $3, $4)"

	_, error := db.Query(updatesql, data.Id, data.Nama, data.Harga, data.Stok)

	if error != nil {
		return false, error
	}

	return true, error

}

func DeleteBarang(db *sql.DB, data entity.Barang) (bool, error) {
	deletesql := "SELECT delete_barang($1)"

	_, error := db.Query(deletesql, data.Id)

	if error != nil {
		return false, error
	}

	return true, error
}
