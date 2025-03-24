package billing

import "time"

// ContractsDetails struct
type ContractsDetails struct {
	ID                  uint64     `gorm:"column:id;primaryKey;autoIncrement"`
	ContractCode        string     `gorm:"column:contract_code"`
	Addendum            string     `gorm:"column:addendum"`
	CustomerCode        string     `gorm:"column:customer_code"`
	CustomerName        string     `gorm:"column:customer_name"`
	ExtNumber          string     `gorm:"column:ext_number"`
	CategoriesCode      string     `gorm:"column:categories_code"`
	CategoriesExpand    string     `gorm:"column:categories_expand"`
	UserCode            string     `gorm:"column:user_code"`
	UserName            string     `gorm:"column:user_name"`
	SignedAt           time.Time  `gorm:"column:signed_at"`
	ActivatedAt        time.Time  `gorm:"column:activated_at"`
	ExpirationAt       time.Time  `gorm:"column:expiration_at"`
	SuspensionAt       *time.Time `gorm:"column:suspension_at"`
	CancellationAt     *time.Time `gorm:"column:cancellation_at"`
	SuspensionReason   *string    `gorm:"column:suspension_reason"`
	CancellationReason *string    `gorm:"column:cancellation_reason"`
	Description        *string    `gorm:"column:description"`
	Discount           *float64   `gorm:"column:discount"`
	Total              *float64   `gorm:"column:total"`
	RoamingFee         float64    `gorm:"column:RoamingFee"`
	CostExpand         float64    `gorm:"column:cost_expand"`
	ActiveFee          float64    `gorm:"column:active_fee"`
	PaymentCycle       int        `gorm:"column:payment_cycle"`
	Status             string     `gorm:"column:status"`
	LastUpdatedAt      *time.Time `gorm:"column:last_updated_at"`
	UserCreated        *string    `gorm:"column:user_created"`
	UserUpdated        *string    `gorm:"column:user_updated"`
	Log                *string    `gorm:"column:log"`
	IsShow             bool       `gorm:"column:IsShow"`
	ReportPaymentCycle int        `gorm:"column:report_payment_cycle"`
	ClassifyID         *int       `gorm:"column:classify_id"`
	ShowSelect         int        `gorm:"column:show_select"`
	ShowSelectBy       string     `gorm:"column:show_select_by"`
	NumberLink         *string    `gorm:"column:number_link"`
	NoteAdmin          *string    `gorm:"column:note_admin"`
	IsVIP              *string    `gorm:"column:is_vip"`
	FreeFee            *string    `gorm:"column:free_fee"`
	CommitFee          *float64   `gorm:"column:commit_fee"`
}

func (ContractsDetails) TableName() string {
	return "contracts_details"
}
