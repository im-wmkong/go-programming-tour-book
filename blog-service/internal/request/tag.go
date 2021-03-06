package request

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State *uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State *uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     *uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	State     *uint8 `form:"state"  binding:"required,oneof=0 1"`
	UpdatedBy string `form:"updated_by" binding:"required,min=3,max=100"`
}
