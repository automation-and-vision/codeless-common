package db_error_helper

import (
	"errors"
	"gorm.io/gorm"
)

func GetDbErrorCode(e error) int {
	switch true {

	case errors.Is(e, gorm.ErrRecordNotFound):
		{
			return 404
		}

	case errors.Is(e, gorm.ErrInvalidData):
		{
			return 400
		}

	case errors.Is(e, gorm.ErrInvalidTransaction):
		{
			return 400
		}

	case errors.Is(e, gorm.ErrInvalidDB):
		{
			return 500
		}

	default:
		{
			return 500
		}
	}
}
