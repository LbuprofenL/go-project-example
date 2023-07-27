package controller

import (
	"net/url"
	"project0/repository"
	"project0/service"
	"strconv"
	"time"
)

/*
	func AppendPost(parentIdStr string, contentStr string, createTimeStr string, filePath string) *PageData {
		parentId, err := strconv.ParseInt(parentIdStr, 10, 64)
		if err != nil {
			return &PageData{Code: -1,
				Msg: err.Error(),
			}
		}

		createTime, err := strconv.ParseInt(createTimeStr, 10, 64)
		if err != nil {
			return &PageData{Code: -1,
				Msg: err.Error(),
			}
		}

		pageData, err := service.AppendPost(&repository.Post{ParentId: parentId, Content: contentStr, CreateTime: createTime}, filePath)
		if err != nil {
			return &PageData{Code: -1,
				Msg: err.Error(),
			}
		}
		return &PageData{Code: 0, Msg: "success", Data: pageData}
	}
*/
func AppendPost(values *url.Values, filePath string) *PageData {
	parentId, err := strconv.ParseInt(values.Get("parent_id"), 10, 64)
	if err != nil {
		return &PageData{Code: -1,
			Msg: err.Error(),
		}
	}
	createTime := time.Now().UnixNano()
	content := values.Get("content")
	pageData, err := service.AppendPost(&repository.Post{ParentId: parentId, Content: content, CreateTime: createTime}, filePath)
	if err != nil {
		return &PageData{Code: -1,
			Msg: err.Error(),
		}
	}
	return &PageData{Code: 0, Msg: "success", Data: pageData}
}
