package service

import (
	"errors"
	"project0/repository"

	"github.com/bytedance/gopkg/lang/fastrand"
)

// 控制代码流程
type AppendPostFlow struct {
	newPo *repository.Post
	Path  string
}

// 参数检查
func (a *AppendPostFlow) checkParam() error {
	if a.newPo.ParentId <= 0 {
		return errors.New("topic id must be larger than 0")
	}
	//分配一个独一无二的id,考虑高并发
	a.newPo.Id = (fastrand.Int63())
	return nil
}

// 添加到文件末尾，
func (a *AppendPostFlow) appendPost() error {
	repository.AppendPost(a.newPo, a.Path)
	return nil
}

// 执行
func (a *AppendPostFlow) Do() (*PageInfo, error) {
	if err := a.checkParam(); err != nil {
		return nil, err
	}
	if err := a.appendPost(); err != nil {
		return nil, err
	}
	pageInfo, err := QueryPageInfo(a.newPo.ParentId)
	if err != nil {
		return nil, err
	}
	return pageInfo, nil
}

// 用函数代替变量构造对象
func NewAppendPost(newPost *repository.Post, filePath string) *AppendPostFlow {
	return &AppendPostFlow{newPo: newPost, Path: filePath}
}

func AppendPost(newPost *repository.Post, filePath string) (*PageInfo, error) {
	return NewAppendPost(newPost, filePath).Do()
}
