package request

type UriIDRequest struct {
	ID int `uri:"id" binding:"required,gt=0"`
}
