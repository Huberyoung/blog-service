package errorcode

var (
	ErrorGetTagListFail = NewError(20010001, "获取标签错误")
	ErrorCreateTagFail  = NewError(20010002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(20010003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20010004, "删除标签失败")
	ErrorCountTagFail   = NewError(20010005, "统计标签失败")

	ErrorGetArticleListFail = NewError(20010011, "获取标签错误")
	ErrorCreateArticleFail  = NewError(20010012, "创建标签失败")
	ErrorUpdateArticleFail  = NewError(20010013, "更新标签失败")
	ErrorDeleteArticleFail  = NewError(20010014, "删除标签失败")
	ErrorCountArticleFail   = NewError(20010015, "统计标签失败")
	ErrorUploadFileFail     = NewError(20030001, "上传文件失败")
)
