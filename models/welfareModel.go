package models

type (
	WelfareModel struct {
		Model
		WelfareName string `json:"welfare_name"`
		Type        int    `json:"type"`
		RightId     int    `json:"right_id"`
		RightName   string `json:"right_name"`
		Price       int    `json:"price"`
	}
)

// 自定义表名
func (WelfareModel) TableName() string {
	return "welfare"
}
