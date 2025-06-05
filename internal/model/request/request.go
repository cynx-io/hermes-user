package request

type PaginateRequest struct {
	Sort    *string `json:"sort" validate:"omitempty,oneof=ASC DESC"`
	Keyword *string `json:"keyword"`
	Page    int     `json:"page" validate:"required"`
	Size    int     `json:"size" validate:"required"`
}
