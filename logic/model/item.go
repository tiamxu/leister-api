package model

// Item 项目信息表
type Item struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	CodeID        int    `json:"code_id" gorm:"column:code_id"`
	AppName       string `json:"app_name" gorm:"column:app_name"`
	AppGroup      string `json:"app_group" gorm:"column:app_group"`
	AppType       string `json:"app_type" gorm:"column:app_type"`
	SSHURLToRepo  string `json:"ssh_url_to_repo" gorm:"column:ssh_url_to_repo"`
	HTTPURLToRepo string `json:"http_url_to_repo" gorm:"column:http_url_to_repo"`
	CreatedAt     int64  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     int64  `json:"updated_at" gorm:"column:updated_at"`
}

// TableName 指定表名
func (Item) TableName() string {
	return "item"
}
