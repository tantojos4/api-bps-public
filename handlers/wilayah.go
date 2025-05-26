package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yhartanto178dev/api-bps-public/services"
)

func SyncProvinsi(c echo.Context) error {
	data, err := services.FetchWilayah("provinsi", "2024_1.2022", "")
	if err != nil {
		return err
	}
	err = services.SaveJSON("data/provinsi.json", data)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, data)
}

func SyncKabupaten(c echo.Context) error {
	provID := c.Param("provID")
	data, err := services.FetchWilayah("kabupaten", "2024_1.2022", provID)
	if err != nil {
		return err
	}
	err = services.SaveJSON("data/kabupaten_"+provID+".json", data)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, data)
}

func SyncKecamatan(c echo.Context) error {
	kabID := c.Param("kabID")
	data, err := services.FetchWilayah("kecamatan", "2024_1.2022", kabID)
	if err != nil {
		return err
	}
	err = services.SaveJSON("data/kecamatan_"+kabID+".json", data)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, data)
}

func SyncDesa(c echo.Context) error {
	kecID := c.Param("kecID")
	data, err := services.FetchWilayah("desa", "2024_1.2022", kecID)
	if err != nil {
		return err
	}
	err = services.SaveJSON("data/desa_"+kecID+".json", data)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, data)
}

func GetProvinsiData(c echo.Context) error {
	data, err := services.LoadOrFetchWilayah("provinsi", "2024_1.2022", "")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func GetKabupatenData(c echo.Context) error {
	provID := c.Param("provID")
	data, err := services.LoadOrFetchWilayah("kabupaten", provID, "2024_1.2022")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func GetKecamatanData(c echo.Context) error {
	kabID := c.Param("kabID")
	data, err := services.LoadOrFetchWilayah("kecamatan", kabID, "2024_1.2022")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func GetDesaData(c echo.Context) error {
	kecID := c.Param("kecID")
	data, err := services.LoadOrFetchWilayah("desa", kecID, "2024_1.2022")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}
