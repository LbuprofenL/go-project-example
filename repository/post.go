package repository

import (
	"sync"
)

type Post struct {
	Id         int64  `json:"id"`
	ParentId   int64  `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

type PostDao struct {
}

var (
	postDao  *PostDao  //Data Access Object
	postOnce sync.Once //sync.Once是同步原语，用于执行仅需执行一次的操作
)

func NewPostDaoInstance() *PostDao {
	//Do以函数作为参数
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

// 使用匿名函数封装是为了将 postDao = &PostDao{}
// 这一初始化操作作为一个单独的函数传递给 postOnce.Do()，
// 确保其只被执行一次。这样在每次调用 NewPostDaoInstance() 时，
// 都会先检查 postDao 是否已经被初始化，如果没有，则执行初始化操作，
// 否则直接返回已初始化的 postDao。

// 多次调用NewPostDaoInstance后Do方法不会被执行，但postDao依旧为首次被调用的值

func (*PostDao) QueryPostByParentId(ParentId int64) []*Post {
	return postIndexMap[ParentId]
}
