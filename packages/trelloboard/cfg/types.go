package cfg

/*-------------------------------------------------------------------------------------*/

type User struct {
	ID     int    `json:"user_id" gorm:"column:id;PRIMARY_KEY;AUTO_INCREMENT"`
	Email  string `json:"email" gorm:"column:email;unique;not null"`
	RoleID *int   `json:"role_id" gorm:"column:role_id"`
}

func (User) TableName() string {
	return "user"
}

/*-------------------------------------------------------------------------------------*/
