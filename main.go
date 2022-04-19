package main

import (
	"os"
	"standarisasigoapi/controller"

	"github.com/labstack/echo"
)

func main() {
	r := echo.New()

	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "Penjualan")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "password")

	r.GET("/barang", controller.GetBarang)
	r.POST("/barang", controller.AddBarang)
	r.PUT("/barang/:id_barang", controller.UpdateBarang)
	r.DELETE("/barang/:id_barang", controller.DeleteBarang)
	r.GET("/pembeli", controller.GetPembeli)
	r.POST("/pembeli", controller.AddPembeli)
	r.PUT("/pembeli/:id_pembeli", controller.UpdatePembeli)
	r.DELETE("/pembeli/:id_pembeli", controller.DeletePembeli)
	r.GET("/transaksi_join", controller.GetTransaksi)
	r.POST("/transaksi", controller.AddTransaksi)
	r.PUT("/transaksi/:id_transaksi", controller.UpdateTransaksi)
	r.DELETE("/transaksi/:id_transaksi", controller.DeleteTransaksi)

	r.Start(":9000")
}
