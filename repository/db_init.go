package repository

import (
	"bufio"
	"encoding/json"
	"os"
	"sync"
)

var (
	topicIndexMap map[int64]*Topic
	postIndexMap  map[int64][]*Post
	mapMutex      sync.Mutex
)

func InitTopicIndexMap(filepath string) error {
	open, err := os.Open(filepath + "topic")
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*Topic)
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmpMap[topic.Id] = &topic
	}
	mapMutex.Lock()
	topicIndexMap = topicTmpMap
	mapMutex.Unlock()
	return nil
}

func InitPostIndexMap(filepath string) error {
	open, err := os.Open(filepath + "post")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	postTmpMap := make(map[int64][]*Post)
	for scanner.Scan() {
		text := scanner.Text()
		var post Post
		if err := json.Unmarshal([]byte(text), &post); err != nil {
			return nil
		}
		posts, ok := postTmpMap[post.ParentId]
		if !ok {
			postTmpMap[post.ParentId] = []*Post{&post}
			continue
		}
		posts = append(posts, &post)
		postTmpMap[post.ParentId] = posts
	}
	mapMutex.Lock()
	postIndexMap = postTmpMap
	mapMutex.Unlock()
	return nil
}

func Init(filePath string) error {
	if err := InitTopicIndexMap(filePath); err != nil {
		return err
	}
	if err := InitPostIndexMap(filePath); err != nil {
		return err
	}
	return nil
}

func AppendPost(newPost *Post, filepath string) error {
	jsonData, err := json.Marshal(newPost)
	if err != nil {
		return err
	}
	open, err := os.OpenFile(filepath+"post", os.O_APPEND, 0644)

	if err != nil {
		return err
	}
	if _, err := open.WriteString(string(jsonData) + "\n"); err != nil {
		return err
	}

	if err := InitPostIndexMap(filepath); err != nil {
		return err
	}
	return nil
}
