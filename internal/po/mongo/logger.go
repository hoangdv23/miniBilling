package mongo

import (
	"miniBilling/internal/constant"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/kamva/mgm/v3"
)

type Logs struct {
	mgm.DefaultModel `bson:",inline"`
	
	User     *UserInfo                `bson:"user" json:"user,omitempty"`
	Status 			*constant.LogStatus 		`bson:"status" json:"status,omitempty"`
	Action	 		*string 					`bson:"action1" json:"Action,omitempty"`
	Desc	 		*string 					`bson:"desc" json:"Action,omitempty"`
	FileName 		*string 					`bson:"filename" json:"filename,omitempty"`
}

type UserInfo struct {
	ID    			primitive.ObjectID 			`bson:"_id,omitempty" json:"_id,omitempty"`
	UserName 		*string						`bson:"username" json:"username,omitempty"`
	TeleId			*int64						`bson:"tele_id" json:"tele_id,omitempty"`
	TeleUsername	*string						`bson:"tele_user" json:"tele_user,omitempty"`
	UserCode 		*string 					`bson:"user_code" json:"usercode,omitempty"`
	Role 			*string						`bson:"role" json:"role,omitempty"`
	Company 		*string						`bson:"company" json:"company,omitempty"`
}