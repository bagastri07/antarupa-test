package customerror

import (
	"net/http"

	"github.com/bagastri07/antarupa-test/model"
	"github.com/labstack/echo/v4"
)

func NewHttpCustomErr(status int, message string) *echo.HTTPError {
	return echo.NewHTTPError(status, model.MessageResponse{
		Message: message,
	})
}

const (
	internalServerErr     = "InternalServerErr"
	notEnoughGameCurrency = "not enough game currency"
	notFound              = "not found"
	maxItemReached        = "max item reached"
)

var (
	ErrInternalServerErr     = NewHttpCustomErr(http.StatusInternalServerError, internalServerErr)
	ErrNotEnoughGameCurrency = NewHttpCustomErr(http.StatusBadRequest, notEnoughGameCurrency)
	ErrNotFound              = NewHttpCustomErr(http.StatusNotFound, notFound)
	ErrMaxItemReached        = NewHttpCustomErr(http.StatusBadRequest, maxItemReached)
)
