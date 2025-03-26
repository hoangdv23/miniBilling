package billing

import "time"

type CdrRecord struct {
	ID              uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	Server          string    `gorm:"column:server;type:varchar(50);not null;default:''" json:"server"`
	CdrID           string    `gorm:"column:CdrID;type:varchar(150);not null;default:''" json:"cdr_id"`
	Caller          string    `gorm:"column:Caller;type:varchar(20);not null;default:''" json:"caller"`
	Callee          string    `gorm:"column:Callee;type:varchar(20);not null;default:''" json:"callee"`
	Duration        *uint     `gorm:"column:duration" json:"duration"`
	Time            time.Time `gorm:"column:time;not null" json:"time"`
	TimeEnd         *time.Time `gorm:"column:time_end" json:"time_end"`
	Minute          uint      `gorm:"column:minute;default:0" json:"minute"`
	LocalMinute     *uint     `gorm:"column:local_minute" json:"local_minute"`
	CallerObject    string    `gorm:"column:caller_object;type:varchar(30);not null;default:''" json:"caller_object"`
	CalleeObject    string    `gorm:"column:callee_object;type:varchar(30);not null;default:''" json:"callee_object"`
	CallType        string    `gorm:"column:call_type;type:varchar(30);not null;default:''" json:"call_type"`
	FixedType       string    `gorm:"column:fixed_type;type:varchar(30);not null;default:''" json:"fixed_type"`
	CallerGw        string    `gorm:"column:caller_gw;type:varchar(100);not null;default:''" json:"caller_gw"`
	CalleeGw        string    `gorm:"column:callee_gw;type:varchar(100);not null;default:''" json:"callee_gw"`
	CallerIP        string    `gorm:"column:caller_ip;type:varchar(50);not null;default:''" json:"caller_ip"`
	CalleeIP        string    `gorm:"column:callee_ip;type:varchar(50);not null;default:''" json:"callee_ip"`
	ContractCode    string    `gorm:"column:contract_code;type:varchar(255);not null" json:"contract_code"`
	CustomerCode    string    `gorm:"column:customer_code;type:varchar(30);not null;default:''" json:"customer_code"`
	CustomerName    string    `gorm:"column:customer_name;type:varchar(255);not null;default:''" json:"customer_name"`
	UserCode        string    `gorm:"column:user_code;type:varchar(30);not null;default:''" json:"user_code"`
	UserName        string    `gorm:"column:user_name;type:varchar(255);not null;default:''" json:"user_name"`
	UserGroup       string    `gorm:"column:user_group;type:varchar(30);not null;default:''" json:"user_group"`
	AgentCode       string    `gorm:"column:agent_code;type:varchar(30);not null;default:''" json:"agent_code"`
	AgentName       string    `gorm:"column:agent_name;type:varchar(255);not null;default:''" json:"agent_name"`
	CategoriesCode  *string   `gorm:"column:categories_code;type:varchar(50)" json:"categories_code"`
	CategoriesExpand *string  `gorm:"column:categories_expand;type:varchar(75)" json:"categories_expand"`
	NumberID        string    `gorm:"column:number_id;type:varchar(20);not null;default:''" json:"number_id"`
	Provider        string    `gorm:"column:provider;type:varchar(50);not null;default:''" json:"provider"`
	CallPrice       float64   `gorm:"column:call_price;default:0" json:"call_price"`
	LocalCost       float64   `gorm:"column:local_cost;default:0" json:"local_cost"`
	AgentCost       float64   `gorm:"column:agent_cost;default:0" json:"agent_cost"`
	ApiCost         float64   `gorm:"column:api_cost;default:0" json:"api_cost"`
	ProviderCost    float64   `gorm:"column:provider_cost;default:0" json:"provider_cost"`
	Cost            float64   `gorm:"column:cost;default:0" json:"cost"`
	Service         uint      `gorm:"column:service;default:0" json:"service"`
}
