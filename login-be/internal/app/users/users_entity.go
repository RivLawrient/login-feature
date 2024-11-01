package users

type Users struct {
	ID        string  `gorm:"column:id;primaryKey"`
	Name      string  `gorm:"column:name"`
	Email     string  `gorm:"column:email;unique"`
	Password  *string `gorm:"column:password"`
	IsGoogle  bool    `gorm:"column:is_google"`
	IsGithub  bool    `gorm:"column:is_github"`
	Token     string  `gorm:"column:token"`
	CreatedAt int64   `gorm:"column:created_at;autoCreateTime:milli"`
}

func (u *Users) TableName() string {
	return "users"
}
