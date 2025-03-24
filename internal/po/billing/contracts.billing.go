package billing
import "time"

type Contracts struct {
	ID                  uint64    `gorm:"column:id; primaryKey; autoIncrement"`
	VosName             string    `gorm:"column:vos_name"`
	VosStatus           string    `gorm:"column:vos_status"`
	BillingStatus       string    `gorm:"column:billing_status"`
	ContractCode        string    `gorm:"column:contract_code"`
	Addendum            string    `gorm:"column:addendum"`
	CustomerCode        string    `gorm:"column:customer_code"`
	CustomerName        string    `gorm:"column:customer_name"`
	TaxCode             string    `gorm:"column:tax_code"`
	UserCode            string    `gorm:"column:user_code"`
	UserName            string    `gorm:"column:user_name"`
	UserGroup           string    `gorm:"column:user_group"`
	AgentCode           string    `gorm:"column:agent_code"`
	AgentName           string    `gorm:"column:agent_name"`
	CategoriesCode      string    `gorm:"column:categories_code"`
	CategoriesExpand    string    `gorm:"column:categories_expand"`
	CustomerEmail       string    `gorm:"column:customer_email"`
	UserEmail           string    `gorm:"column:user_email"`
	SignedAt           time.Time `gorm:"column:signed_at"`
	ActivatedAt        time.Time `gorm:"column:activated_at"`
	ExpirationAt       time.Time `gorm:"column:expiration_at"`
	Status             string    `gorm:"column:status"`
	CountNumber        int       `gorm:"column:count_number"`
	CountExt           int       `gorm:"column:count_ext"`
	Menu               string    `gorm:"column:menu"`
	Classify           int       `gorm:"column:classify"`
	Ratio              float64   `gorm:"column:ratio"`
	LimitMoney         float64   `gorm:"column:limit_money"`
	NowCost            float64   `gorm:"column:now_cost"`
	NowMinute          float64   `gorm:"column:now_minute"`
	CallCost           float64   `gorm:"column:call_cost"`
	TotalCost          float64   `gorm:"column:total_cost"`
	SumLi              float64   `gorm:"column:sum_li"`
	SumPa              float64   `gorm:"column:sum_pa"`
	PermanentCost      float64   `gorm:"column:permanent_cost"`
	Wallet             float64   `gorm:"column:wallet"`
	ReceiveMail        float64   `gorm:"column:receive_mail"`
	ActivePayment      float64   `gorm:"column:active_payment"`
	Condition0         float64   `gorm:"column:condition_0"`
	Condition1         float64   `gorm:"column:condition_1"`
	Condition2         float64   `gorm:"column:condition_2"`
	Condition3         float64   `gorm:"column:condition_3"`
	Condition4         float64   `gorm:"column:condition_4"`
	Condition5         float64   `gorm:"column:condition_5"`
	Condition6         float64   `gorm:"column:condition_6"`
	ConditionFinal     float64   `gorm:"column:condition_final"`
	FinalCostCallCenter float64  `gorm:"column:final_cost_callcenter"`
	FinalCostCrm       float64   `gorm:"column:final_cost_crm"`
	TotalPrice         *float64  `gorm:"column:total_price"`
	Description        *string   `gorm:"column:description"`
	SendChargeNotice   int       `gorm:"column:send_charge_notice"`
	LastUpdatedAt      *time.Time `gorm:"column:last_updated_at"`
	UserCreated        *string   `gorm:"column:user_created"`
	UserUpdated        *string   `gorm:"column:user_updated"`
	Log                *string   `gorm:"column:log"`
	IsShow             bool      `gorm:"column:IsShow"`
	ReportPaymentCycle int       `gorm:"column:report_payment_cycle"`
	ModuleCrm          *string   `gorm:"column:module_crm"`
	ShowSelect         int       `gorm:"column:show_select"`
	ShowSelectBy       string    `gorm:"column:show_select_by"`
	ReportNumberBlock  *string   `gorm:"column:report_number_block"`
}

func (Contracts) TableName() string {
	return "contracts"
}
