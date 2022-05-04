package model

import (
	"database/sql"
	"standarisasigoapi/entity"

	_ "github.com/lib/pq"
)

func GetPembeli(db *sql.DB) ([]entity.Pembeli, error) {
	readsql := "SELECT * from read_pembeli()"
	rows, error := db.Query(readsql)

	if error != nil {
		return []entity.Pembeli{}, error
	}
	defer rows.Close()

	hasil := []entity.Pembeli{}
	for rows.Next() {
		data := entity.Pembeli{}
		error := rows.Scan(&data.Id, &data.Nama, &data.Jk, &data.No_telp, &data.Alamat, &data.Foto)
		if error != nil {
			return []entity.Pembeli{}, error
		}
		hasil = append(hasil, data)
	}

	return hasil, error
}

func AddPembeli(db *sql.DB, data entity.Pembeli) (bool, error) {
	insertsql := "SELECT create_pembeli($1, $2, $3, $4, $5, $6)"

	_, error := db.Query(insertsql, data.Id, data.Nama, data.Jk, data.No_telp, data.Alamat, data.Foto)

	if error != nil {
		return false, error
	}

	return true, error
}

func UpdatePembeli(db *sql.DB, data entity.Pembeli) (bool, error) {
	updatesql := "SELECT update_pembeli($1, $2, $3, $4, $5, $6)"

	_, error := db.Query(updatesql, data.Id, data.Nama, data.Jk, data.No_telp, data.Alamat, data.Foto)

	if error != nil {
		return false, error
	}

	return true, error

}

func DeletePembeli(db *sql.DB, data entity.Pembeli) (bool, error) {
	deletesql := "SELECT delete_pembeli($1)"

	_, error := db.Query(deletesql, data.Id)

	if error != nil {
		return false, error
	}

	return true, error
}
