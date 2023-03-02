package service

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100,min=3"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100,min=3"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"max=100,min=3"`
	CreatedBy string `form:"created_by" binding:"required,max=100,min=3"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	Id         uint32 `form:"id" binding:"required gte=1"`
	Name       string `form:"name" binding:"max=100,min=3"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,max=100,min=3"`
}

type DeleteTagRequest struct {
	Id uint32 `form:"id" binding:"required,gte=1"`
}
