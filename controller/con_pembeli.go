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

func GetPembeli(ctx echo.Context) error {
	db := settings.AccessDatabase()

	result, err := model.GetPembeli(db)
	if err != nil {
		data := entity.M{"Data": result, "Pesan": err.Error(), "Status": 500}
		return ctx.JSON(http.StatusInternalServerError, data)
	} else {
		data := entity.M{"Data": result, "Pesan": "Data pembeli berhasil dibaca", "Status": 200}
		return ctx.JSON(http.StatusOK, data)
	}
}

func AddPembeli(ctx echo.Context) error {
	db := settings.AccessDatabase()

	var pembeli entity.Pembeli

	pembeli.Id = xid.New().Counter()
	pembeli.Nama = ctx.FormValue("nama_pembeli")
	pembeli.Jk = ctx.FormValue("jk")
	pembeli.No_telp = ctx.FormValue("no_telp")
	pembeli.Alamat = ctx.FormValue("alamat")

	_, err := model.AddPembeli(db, pembeli)
	if err != nil {
		data := entity.M{"Pesan": err.Error(), "Status": 500}
		return ctx.JSON(http.StatusInternalServerError, data)
	} else {
		data := entity.M{"Pesan": "Data pembeli berhasil ditambahkan", "Status": 200}
		return ctx.JSON(http.StatusOK, data)
	}
}

func UpdatePembeli(ctx echo.Context) error {
	db := settings.AccessDatabase()

	var pembeli entity.Pembeli

	id_pembeli := ctx.Param("id_pembeli")
	intVar, err := strconv.ParseInt(id_pembeli, 0, 32)
	pembeli.Id = int32(intVar)
	pembeli.Nama = ctx.FormValue("nama_pembeli")
	pembeli.Jk = ctx.FormValue("jk")
	pembeli.No_telp = ctx.FormValue("no_telp")
	pembeli.Alamat = ctx.FormValue("alamat")

	_, err = model.UpdatePembeli(db, pembeli)
	if err != nil {
		data := entity.M{"Pesan": err.Error(), "Status": 500}
		return ctx.JSON(http.StatusInternalServerError, data)
	} else {
		data := entity.M{"Pesan": "Data pembeli berhasil diubah", "Status": 200}
		return ctx.JSON(http.StatusOK, data)
	}
}

func DeletePembeli(ctx echo.Context) error {
	db := settings.AccessDatabase()

	var pembeli entity.Pembeli

	id_pembeli := ctx.Param("id_pembeli")
	intVar, err := strconv.ParseInt(id_pembeli, 0, 32)
	pembeli.Id = int32(intVar)

	_, err = model.DeletePembeli(db, pembeli)
	if err != nil {
		data := entity.M{"Pesan": err.Error(), "Status": 500}
		return ctx.JSON(http.StatusInternalServerError, data)
	} else {
		data := entity.M{"Pesan": "Data pembeli berhasil dihapus", "Status": 200}
		return ctx.JSON(http.StatusOK, data)
	}
}
