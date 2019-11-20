package idCards

type UserModelValidator struct {
	Article struct {
		Title       string   `form:"title" json:"title" binding:"exists,min=4"`
		Description string   `form:"description" json:"description" binding:"max=2048"`
		Body        string   `form:"body" json:"body" binding:"max=2048"`
		Tags        []string `form:"tagList" json:"tagList"`
	} `json:"article"`
	userModel UserModel `json:"-"`
}