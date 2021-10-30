package utils

import (
	"github.com/labstack/echo/v4"
)

type Error struct {
	Code    int64  `json:"code"`
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
	Code      int64       `json:"code"`
	GetRTByID interface{} `json:"get_rt_by_id,omitempty"`
	GetAllRT  interface{} `json:"get_all_rt,omitempty"`
	CreateRT  interface{} `json:"create_rt,omitempty"`
	Message   string      `json:"message"`
}

type JSONResponseDataKeluarga struct {
	Code            int64       `json:"code"`
	GetKeluargaByID interface{} `json:"get_keluarga_by_id,omitempty"`
	GetAllKeluarga  interface{} `json:"get_all_keluarga,omitempty"`
	CreateKeluarga  interface{} `json:"create_keluarga,omitempty"`
	Message         string      `json:"message"`
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

type JSONResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func Response(c echo.Context, res JSONResponse) error {
	return c.JSON(int(res.Code), res)
}

func ResponseData(c echo.Context, res JSONResponseData) error {
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

func ResponseError(c echo.Context, err Error) error {
	return c.JSON(int(err.Code), err)
}
