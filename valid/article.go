package valid

type ArticleGet struct {
	ID uint32 `form:"id" binding:"required,gte=3"`
}

type ArticleCreate struct {
	Title   string `form:"title" binding:"required,min=2,max=30"`
	Desc    string `form:"desc" binding:"required,min=10,max=255"`
	Content string `form:"content" binding:"required,min=10,max=300"`
}
