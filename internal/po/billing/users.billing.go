package billing

type Users struct {
	Id 					int64 		`gorm:"column:id"`
	Main_id 			string 		`gorm:"column:main_id"`
	User_name 			string		`gorm:"column:user_name"`
	User_code 			string		`gorm:"column:user_code"`
	User_group 			string		`gorm:"column:user_group"`
	Role 				string		`gorm:"column:role"`
	Position_id 		int			`gorm:"column:position_id"`
	Position_name 		string		`gorm:"column:position_name"`
	Department_id 		int			`gorm:"column:department_id"`
	Department_name 	string		`gorm:"column:department_name"`
	Company_id 			int			`gorm:"column:company_id"`
	Company_name 		string		`gorm:"column:company_name"`
	Company_code 		string		`gorm:"column:company_code"`
	Parent_id 			int			`gorm:"column:parent_id"`
	Parent_name 		string		`gorm:"column:parent_name"`
	Phone				string		`gorm:"column:phone"`
	Email 				string		`gorm:"column:email"`
	Address				string		`gorm:"column:address"`
	Password 			string		`gorm:"column:password"`
	Password_show		string		`gorm:"column:password_show"`
	Two_factor_secret 	string		`gorm:"column:two_factor_secret"`
	Current_team_id 	string		`gorm:"column:current_team_id"`
	Status				int		`gorm:"column:status"`
}														

func (Users) TableName() string {
	return "users"
}
// Sau này chỉnh giá trị role vì hiện tại nó đang là ["role"]