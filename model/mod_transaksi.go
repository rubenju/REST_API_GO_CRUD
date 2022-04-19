package model

import (
	"database/sql"
	"standarisasigoapi/entity"

	_ "github.com/lib/pq"
)

func GetTransaksi(db *sql.DB) ([]entity.TransaksiJoin, error) {
	readsql := "SELECT * from read_transaksi()"
	rows, error := db.Query(readsql)

	if error != nil {
		return []entity.TransaksiJoin{}, error
	}
	defer rows.Close()

	hasil := []entity.TransaksiJoin{}
	for rows.Next() {
		data := entity.TransaksiJoin{}
		error := rows.Scan(&data.Id_transaksi, &data.Nama_barang, &data.Harga, &data.Nama_pembeli, &data.Tanggal)
		if error != nil {
			return []entity.TransaksiJoin{}, error
		}
		hasil = append(hasil, data)
	}

	return hasil, error
}

func AddTransaksi(db *sql.DB, data entity.Transaksi) (bool, error) {
	insertsql := "SELECT create_transaksi($1, $2, $3, $4, $5)"

	_, error := db.Query(insertsql, data.Id_transaksi, data.Id_barang, data.Id_pembeli, data.Tanggal, data.Keterangan)

	if error != nil {
		return false, error
	}

	return true, error
}

func UpdateTransaksi(db *sql.DB, data entity.Transaksi) (bool, error) {
	updatesql := "SELECT update_transaksi($1, $2, $3, $4, $5)"

	_, error := db.Query(updatesql, data.Id_transaksi, data.Id_barang, data.Id_pembeli, data.Tanggal, data.Keterangan)

	if error != nil {
		return false, error
	}

	return true, error

}

func DeleteTransaksi(db *sql.DB, data entity.Transaksi) (bool, error) {
	deletesql := "SELECT delete_transaksi($1)"

	_, error := db.Query(deletesql, data.Id_transaksi)

	if error != nil {
		return false, error
	}

	return true, error
}
