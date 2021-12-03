package utils

import (
	"github.com/labstack/echo/v4"
)

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ErrorLogin struct {
	Code    int64  `json:"code"`
	Token   string `json:"token"`
	Message string `json:"message"`
}

type JSONResponseLogin struct {
	Code    int64  `json:"code"`
	Token   string `json:"token"`
	Message string `json:"message"`
}
type JSONResponseData struct {
	Code    int64       `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type JSONResponseDataRT struct {
	Code             int64       `json:"code"`
	GetRTByID        interface{} `json:"get_rt_by_id,omitempty"`
	GetAllRT         interface{} `json:"get_all_rt,omitempty"`
	GetAllRTPengurus interface{} `json:"get_all_rt_pengurus,omitempty"`
	GetAllRTKeluarga interface{} `json:"get_all_rt_keluarga,omitempty"`
	CreateRT         interface{} `json:"create_rt,omitempty"`
	Message          string      `json:"message"`
}

type JSONResponseDataProduk struct {
	Code                 int64       `json:"code"`
	GetProdukByID        interface{} `json:"get_produk_by_id,omitempty"`
	GetAllProduk         interface{} `json:"get_all_produk,omitempty"`
	GetAllProdukKeluarga interface{} `json:"get_all_produk_keluarga,omitempty"`
	CreateProduk         interface{} `json:"create_produk,omitempty"`
	Message              string      `json:"message"`
}
type JSONResponseDataKeluarga struct {
	Code                int64       `json:"code"`
	GetKeluargaByID     interface{} `json:"get_keluarga_by_id,omitempty"`
	GetKeluargaSaya     interface{} `json:"get_keluarga_saya,omitempty"`
	GetAllKeluarga      interface{} `json:"get_all_keluarga,omitempty"`
	GetAllKeluargaWarga interface{} `json:"get_all_keluarga_warga,omitempty"`
	CreateKeluarga      interface{} `json:"create_keluarga,omitempty"`
	Message             string      `json:"message"`
}

type JSONResponseDataPengurusRT struct {
	Code            int64       `json:"code"`
	GetPengurusByID interface{} `json:"get_pengurus_by_id,omitempty"`
	GetAllPengurus  interface{} `json:"get_all_pengurus,omitempty"`
	CreatePengurus  interface{} `json:"create_pengurus,omitempty"`
	Message         string      `json:"message"`
}

type JSONResponseDataWarga struct {
	Code         int64       `json:"code"`
	GetWargaByID interface{} `json:"get_warga_by_id,omitempty"`
	GetAllWarga  interface{} `json:"get_all_warga,omitempty"`
	CreateWarga  interface{} `json:"create_warga,omitempty"`
	Message      string      `json:"message"`
}

type JSONResponseDataDompetRT struct {
	Code          int64       `json:"code"`
	GetDompetByID interface{} `json:"get_dompet_by_id,omitempty"`
	GetAllDompet  interface{} `json:"get_all_dompet,omitempty"`
	CreateDompet  interface{} `json:"create_dompet,omitempty"`
	Message       string      `json:"message"`
}

type JSONResponseDataDompetKeluarga struct {
	Code                  int64       `json:"code"`
	GetDompetKeluargaByID interface{} `json:"get_dompet_keluarga_by_id,omitempty"`
	GetAllDompetKeluarga  interface{} `json:"get_all_dompet_keluarga,omitempty"`
	CreateDompetKeluarga  interface{} `json:"create_dompet_keluarga,omitempty"`
	Message               string      `json:"message"`
}

type JSONResponseDataPersuratan struct {
	Code              int64       `json:"code"`
	GetPersuratanByID interface{} `json:"get_persuratan_by_id,omitempty"`
	GetAllPersuratan  interface{} `json:"get_all_persuratan,omitempty"`
	CreatePersuratan  interface{} `json:"create_persuratan,omitempty"`
	Message           string      `json:"message"`
}

type JSONResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type JSONResponseDataOrder struct {
	Code         int64       `json:"code"`
	GetOrderByID interface{} `json:"get_order_by_id,omitempty"`
	GetAllOrder  interface{} `json:"get_all_order,omitempty"`
	CreateOrder  interface{} `json:"create_order,omitempty"`
	Message      string      `json:"message"`
}

func Response(c echo.Context, res JSONResponse) error {
	return c.JSON(int(res.Code), res)
}

func ResponseData(c echo.Context, res JSONResponseData) error {
	return c.JSON(int(res.Code), res)
}

func ResponseDataProduk(c echo.Context, res JSONResponseDataProduk) error {
	return c.JSON(int(res.Code), res)
}

func ResponseLogin(c echo.Context, res JSONResponseLogin) error {
	return c.JSON(int(res.Code), res)
}

func ResponseDataRT(c echo.Context, res JSONResponseDataRT) error {
	return c.JSON(int(res.Code), res)
}

func ResponseDataKeluarga(c echo.Context, res JSONResponseDataKeluarga) error {
	return c.JSON(int(res.Code), res)
}

func ResponseDataPengurusRT(c echo.Context, res JSONResponseDataPengurusRT) error {
	return c.JSON(int(res.Code), res)
}

func ResponseDataWarga(c echo.Context, res JSONResponseDataWarga) error {
	return c.JSON(int(res.Code), res)
}

func ResponseDataDompet(c echo.Context, res JSONResponseDataDompetRT) error {
	return c.JSON(int(res.Code), res)
}

func ResponseDataDompetKeluarga(c echo.Context, res JSONResponseDataDompetKeluarga) error {
	return c.JSON(int(res.Code), res)
}

func ResponseDataPersuratan(c echo.Context, res JSONResponseDataPersuratan) error {
	return c.JSON(int(res.Code), res)
}

func ResponseDataOrder(c echo.Context, res JSONResponseDataOrder) error {
	return c.JSON(int(res.Code), res)
}

func ResponseError(c echo.Context, err Error) error {
	return c.JSON(int(err.Code), err)
}

func ResponseErrorLogin(c echo.Context, err ErrorLogin) error {
	return c.JSON(int(err.Code), err)
}
