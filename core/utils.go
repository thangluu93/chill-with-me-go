package core

import (
	"main/data"
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
		noRecord = data.DEFAULT_PAGE_SIZE
	}
	limit = noRecord
	offset = (page - 1) * noRecord
	return limit, offset
}
