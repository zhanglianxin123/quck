package model

type Conversa struct {
	Id      string `form:"id" json:"id" query:"id"`
	Role    string `form:"role" json:"role" query:"role"`
	Content string `form:"content" json:"content" query:"content"`
}

func (c *Conversa) TableName() string {
	return "conversa"
}
