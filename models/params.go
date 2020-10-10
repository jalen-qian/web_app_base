package models

// 参数示例
type ParamExample struct {
	Username      string `json:"username" form:"username" binding:"required"`
	Password      string `json:"password" form:"password" binding:"required"`
	RePassword    string `json:"re_password" form:"re_password" binding:"required,eqfield=Password"`
	OneOf         string `json:"one_of" form:"one_of" binding:"required,oneof=-1 0 1"` //如果有选项
	NotViewInJson string `json:"-"`                                                    // 不在json渲染时不显示
}
