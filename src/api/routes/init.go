package routes

import (
	"net/http"
	"src/config"
	"src/utils"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	// ErrRateLimitExceeded denotes an error raised when rate limit is exceeded
	ErrRateLimitExceeded = echo.NewHTTPError(http.StatusTooManyRequests, "rate limit exceeded")
	// ErrExtractorError denotes an error raised when extractor function is unsuccessful
	ErrExtractorError = echo.NewHTTPError(http.StatusForbidden, "error while extracting identifier")
)

func Init(e *echo.Echo) *echo.Echo {
	JWTconfig := middleware.JWTConfig{
		BeforeFunc: func(c echo.Context) {
			c.Logger().Info("Token: ", c.Request().Header.Get("Authorization"))
		},
		TokenLookup: "header:Authorization",
		Claims:      &utils.JWTCustomClaims{},
		SigningKey:  []byte(config.GetConfig(e).Secret),
		ErrorHandlerWithContext: func(err error, c echo.Context) error {
			c.Logger().Error("Error JWT Context: ", err)
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  "Token Invalid",
				Internal: err,
			}
		},
	}
	middleware.ErrJWTMissing.Code = http.StatusUnauthorized

	middleware.ErrJWTMissing.Message = utils.Error{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
	}
	middleware.ErrJWTInvalid.Message = utils.Error{
		Code:    http.StatusUnauthorized,
		Message: "Token Invalid",
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// DefaultRateLimiterConfig defines default values for RateLimiterConfig
	var DefaultRateLimiterConfig = middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 30, ExpiresIn: 5 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return &echo.HTTPError{
				Code:     ErrExtractorError.Code,
				Message:  ErrExtractorError.Message,
				Internal: err,
			}
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return &echo.HTTPError{
				Code:     ErrRateLimitExceeded.Code,
				Message:  ErrRateLimitExceeded.Message,
				Internal: err,
			}
		},
	}
	e.Use(middleware.RateLimiterWithConfig(DefaultRateLimiterConfig))

	e.Logger.Info("menginisialisasikan routes")
	e = Keluarga(e, JWTconfig)
	e = RT(e, JWTconfig)
	e = PengurusRT(e, JWTconfig)
	e = Warga(e, JWTconfig)
	e = Produk(e, JWTconfig)
	e = DompetRT(e, JWTconfig)
	e = Order(e, JWTconfig)
	e = Persuratan(e, JWTconfig)
	e = DompetKeluarga(e, JWTconfig)
	e = Informasi(e, JWTconfig)
	e = Aduan(e, JWTconfig)
	e = Tagihan(e, JWTconfig)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Worlds!!!")
	})

	e.Logger.Info("routes terinisialisasi")

	return e
}
