package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yhartanto178dev/api-bps-public/handlers"
)

func main() {
	e := echo.New()

	e.GET("/sync/provinsi", handlers.SyncProvinsi)
	e.GET("/sync/kabupaten/:provID", handlers.SyncKabupaten)
	e.GET("/sync/kecamatan/:kabID", handlers.SyncKecamatan)
	e.GET("/sync/desa/:kecID", handlers.SyncDesa)

	//get wilayah data dari Json files

	e.GET("/data/provinsi", handlers.GetProvinsiData)
	e.GET("/data/kabupaten/:provID", handlers.GetKabupatenData)
	e.GET("/data/kecamatan/:kabID", handlers.GetKecamatanData)
	e.GET("/data/desa/:kecID", handlers.GetDesaData)

	e.Logger.Fatal(e.Start(":8080"))
}
