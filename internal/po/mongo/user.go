package mongo

import (
	"miniBilling/internal/constant"

	"github.com/kamva/mgm/v3"
)

type Users struct {
	mgm.DefaultModel `bson:",inline"`
	
	UserCode 		*string 					`bson:"user_code" json:"usercode,omitempty"`
	UserName 		*string						`bson:"username" json:"username,omitempty"`
	TeleName 		*string						`bson:"tele_name" json:"tele_name,omitempty"`
	TeleId			*int64						`bson:"tele_id" json:"tele_id,omitempty"`
	TeleUsername	*string						`bson:"tele_user" json:"tele_user,omitempty"`
	Password 		*string						`bson:"password" json:"password,omitempty"`
	Email 			*string 					`bson:"email" json:"email,omitempty"`
	Role 			*string						`bson:"role" json:"role,omitempty"`
	Company 		*string						`bson:"company" json:"company,omitempty"`
	Status 			*constant.UserStatus 		`bson:"status" json:"status,omitempty"`
	Action1 		*string 					`bson:"action1" json:"action1,omitempty"`
	Action2 		*string 					`bson:"action2" json:"action2,omitempty"`
	Action3 		*string 					`bson:"action3" json:"action3,omitempty"`
	Action4 		*string 					`bson:"action4" json:"action4,omitempty"`
	Action5 		*string 					`bson:"action5" json:"action5,omitempty"`
	Action6 		*string 					`bson:"action6" json:"action6,omitempty"`
}

type UpdateUsers struct {
	UserCode 		*string 					`bson:"user_code" json:"usercode,omitempty"`
	UserName 		*string						`bson:"username" json:"username,omitempty"`
	TeleName 		*string						`bson:"tele_name" json:"tele_name,omitempty"`
	TeleId			*int64						`bson:"tele_id" json:"tele_id,omitempty"`
	TeleUsername	*string						`bson:"tele_user" json:"tele_user,omitempty"`
	Password 		*string						`bson:"password" json:"password,omitempty"`
	Email 			*string 					`bson:"email" json:"email,omitempty"`
	Role 			*string					`bson:"role" json:"role,omitempty"`
	Company 		*string						`bson:"company" json:"company,omitempty"`
	Status 			*constant.UserStatus 		`bson:"status" json:"status,omitempty"`
	Action1 		*string 					`bson:"action1" json:"action1,omitempty"`
	Action2 		*string 					`bson:"action2" json:"action2,omitempty"`
	Action3 		*string 					`bson:"action3" json:"action3,omitempty"`
	Action4 		*string 					`bson:"action4" json:"action4,omitempty"`
	Action5 		*string 					`bson:"action5" json:"action5,omitempty"`
	Action6 		*string 					`bson:"action6" json:"action6,omitempty"`
}

