package core

import (
	"errors"
	"main/data"
	"strconv"
)

type Utility struct {
}

func UseUtil() *Utility {
	return &Utility{}
}

func (u *Utility) GetLimitOffset(page int, noRecord int) (limit int, offset int) {
	if page == 0 {
		page = 1
	}
	if noRecord == 0 {
		noRecord = data.DefaultPageSize
	}
	limit = noRecord
	offset = (page - 1) * noRecord
	return limit, offset
}

func (u *Utility) StringToInt(number string) (intNumber int, err error) {
	var numberInt int
	if number == "" {
		return numberInt, errors.New("number is empty")
	}
	numberInt, errorParse := strconv.Atoi(number)
	if errorParse != nil {
		return numberInt, errorParse
	}
	return numberInt, nil
}
