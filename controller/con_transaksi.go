package controller

import (
	"net/http"
	"standarisasigoapi/entity"
	"standarisasigoapi/model"
	"standarisasigoapi/settings"
	"strconv"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/rs/xid"
)

func GetTransaksi(ctx echo.Context) error {
	db := settings.AccessDatabase()

	result, err := model.GetTransaksi(db)
	if err != nil {
		data := entity.M{"Data": result, "Pesan": err.Error(), "Status": 500}
		return ctx.JSON(http.StatusInternalServerError, data)
	} else {
		data := entity.M{"Data": result, "Pesan": "Data transaksi berhasil dibaca", "Status": 200}
		return ctx.JSON(http.StatusOK, data)
	}
}

func AddTransaksi(ctx echo.Context) error {
	db := settings.AccessDatabase()

	var transaksi entity.Transaksi

	transaksi.Id_transaksi = xid.New().Counter()
	transaksi.Id_barang = ctx.FormValue("id_barang")
	transaksi.Id_pembeli = ctx.FormValue("id_pembeli")
	transaksi.Tanggal = ctx.FormValue("tanggal")
	transaksi.Keterangan = ctx.FormValue("keterangan")

	_, err := model.AddTransaksi(db, transaksi)
	if err != nil {
		data := entity.M{"Pesan": err.Error(), "Status": 500}
		return ctx.JSON(http.StatusInternalServerError, data)
	} else {
		data := entity.M{"Pesan": "Data transaksi berhasil ditambahkan", "Status": 200}
		return ctx.JSON(http.StatusOK, data)
	}
}

func UpdateTransaksi(ctx echo.Context) error {
	db := settings.AccessDatabase()

	var transaksi entity.Transaksi

	id_transaksi := ctx.Param("id_transaksi")
	intVar, err := strconv.ParseInt(id_transaksi, 0, 32)
	transaksi.Id_transaksi = int32(intVar)
	transaksi.Id_barang = ctx.FormValue("id_barang")
	transaksi.Id_pembeli = ctx.FormValue("id_pembeli")
	transaksi.Tanggal = ctx.FormValue("tanggal")
	transaksi.Keterangan = ctx.FormValue("keterangan")

	_, err = model.UpdateTransaksi(db, transaksi)
	if err != nil {
		data := entity.M{"Pesan": err.Error(), "Status": 500}
		return ctx.JSON(http.StatusInternalServerError, data)
	} else {
		data := entity.M{"Pesan": "Data transaksi berhasil diubah", "Status": 200}
		return ctx.JSON(http.StatusOK, data)
	}
}

func DeleteTransaksi(ctx echo.Context) error {
	db := settings.AccessDatabase()

	var transaksi entity.Transaksi

	id_transaksi := ctx.Param("id_transaksi")
	intVar, err := strconv.ParseInt(id_transaksi, 0, 32)
	transaksi.Id_transaksi = int32(intVar)

	_, err = model.DeleteTransaksi(db, transaksi)
	if err != nil {
		data := entity.M{"Pesan": err.Error(), "Status": 500}
		return ctx.JSON(http.StatusInternalServerError, data)
	} else {
		data := entity.M{"Pesan": "Data transaksi berhasil dihapus", "Status": 200}
		return ctx.JSON(http.StatusOK, data)
	}
}
