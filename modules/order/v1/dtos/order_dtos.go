package order_dtos

type OrderFilter struct {
	UserID    uint
	Status    string
	Sort      string // "ASC" or "DESC"
	StartDate string // Optional: Start date for filtering
	EndDate   string // Optional: End date for filtering
	UserName  string // Optional: Filter by user name
	Email     string // Optional: Filter by user email

}
