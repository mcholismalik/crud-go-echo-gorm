package model

type PaginationDTO struct {
	Page     int `query:"page" validate:"required"`
	PageSize int `query:"page_size" validate:"required"`
}

type PaginationInfoDTO struct {
	Page        int  `json:"page"`
	PageSize    int  `json:"page_size"`
	Count       int  `json:"count"`
	MoreRecords bool `json:"more_records"`
}
