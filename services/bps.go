package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func FetchWilayah(level, periode, parent string) (interface{}, error) {
	var url string
	if level == "provinsi" || level == "kecamatan" {
		url = fmt.Sprintf("https://sig.bps.go.id/rest-drop-down/getwilayah?level=%s&parent=%s&periode_merge=%s", level, parent, periode)
	} else {
		url = fmt.Sprintf("https://sig.bps.go.id/rest-bridging/getwilayah?level=%s&parent=%s&periode_merge=%s", level, parent, periode)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func SaveJSON(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func LoadJSON(filename string) (interface{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data interface{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func LoadOrFetchWilayah(level string, parent string, periode string) (interface{}, error) {
	var filename string
	switch level {
	case "provinsi":
		filename = "data/provinsi.json"
	case "kabupaten":
		filename = "data/kabupaten_" + parent + ".json"
	case "kecamatan":
		filename = "data/kecamatan_" + parent + ".json"
	case "desa":
		filename = "data/desa_" + parent + ".json"
	default:
		return nil, fmt.Errorf("level tidak dikenal")
	}

	// Coba load file lokal dulu
	data, err := LoadJSON(filename)
	if err == nil {
		return data, nil
	}

	// Kalau gagal, fetch dari BPS
	data, err = FetchWilayah(level, periode, parent)
	if err != nil {
		return nil, err
	}

	// Simpan ke file
	err = SaveJSON(filename, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
