package order_dtos

type OrderFilter struct {
	UserID    uint   `form:"user_id" validate:"omitempty,min=1"`
	Status    string `form:"status" validate:"omitempty,oneof=pending preparing ready delivered cancelled"`
	Sort      string `form:"sort" validate:"omitempty,oneof=ASC DESC"`
	StartDate string `form:"start_date" validate:"omitempty,datetime=2006-01-02"`
	EndDate   string `form:"end_date" validate:"omitempty,datetime=2006-01-02"`
	UserName  string `form:"user_name" validate:"omitempty,min=1,max=100"`
	Email     string `form:"email" validate:"omitempty,email"`
}
