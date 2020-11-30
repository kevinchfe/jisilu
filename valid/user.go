package valid

type UserReq struct {
	Id   string `form:"id";binding:"required,min=0"`
	Name string `form:"name";binding:"required"`
}
