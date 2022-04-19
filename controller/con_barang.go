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

func GetBarang(ctx echo.Context) error {
	db := settings.AccessDatabase()

	result, err := model.GetBarang(db)
	if err != nil {
		data := entity.M{"Data": result, "Pesan": err.Error(), "Status": 500}
		return ctx.JSON(http.StatusInternalServerError, data)
	} else {
		data := entity.M{"Data": result, "Pesan": "Data barang berhasil dibaca", "Status": 200}
		return ctx.JSON(http.StatusOK, data)
	}
}

func AddBarang(ctx echo.Context) error {
	db := settings.AccessDatabase()

	var barang entity.Barang

	barang.Id = xid.New().Counter()
	barang.Nama = ctx.FormValue("nama_barang")
	barang.Harga = ctx.FormValue("harga")
	barang.Stok = ctx.FormValue("stok")

	_, err := model.AddBarang(db, barang)
	if err != nil {
		data := entity.M{"Pesan": err.Error(), "Status": 500}
		return ctx.JSON(http.StatusInternalServerError, data)
	} else {
		data := entity.M{"Pesan": "Data barang berhasil ditambahkan", "Status": 200}
		return ctx.JSON(http.StatusOK, data)
	}
}

func UpdateBarang(ctx echo.Context) error {
	db := settings.AccessDatabase()

	var barang entity.Barang

	id_barang := ctx.Param("id_barang")
	intVar, err := strconv.ParseInt(id_barang, 0, 32)
	barang.Id = int32(intVar)
	barang.Nama = ctx.FormValue("nama_barang")
	barang.Harga = ctx.FormValue("harga")
	barang.Stok = ctx.FormValue("stok")

	_, err = model.UpdateBarang(db, barang)
	if err != nil {
		data := entity.M{"Pesan": err.Error(), "Status": 500}
		return ctx.JSON(http.StatusInternalServerError, data)
	} else {
		data := entity.M{"Pesan": "Data barang berhasil diubah", "Status": 200}
		return ctx.JSON(http.StatusOK, data)
	}
}

func DeleteBarang(ctx echo.Context) error {
	db := settings.AccessDatabase()

	var barang entity.Barang

	id_barang := ctx.Param("id_barang")
	intVar, err := strconv.ParseInt(id_barang, 0, 32)
	barang.Id = int32(intVar)

	_, err = model.DeleteBarang(db, barang)
	if err != nil {
		data := entity.M{"Pesan": err.Error(), "Status": 500}
		return ctx.JSON(http.StatusInternalServerError, data)
	} else {
		data := entity.M{"Pesan": "Data barang berhasil dihapus", "Status": 200}
		return ctx.JSON(http.StatusOK, data)
	}
}
