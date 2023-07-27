package service

import (
	"os"
	"testing"

	"go-project-example/repository"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	repository.Init("../data/")
	os.Exit(m.Run())
}
func TestQueryPageInfo(t *testing.T) {
	pageInfo, _ := QueryPageInfo(1)
	assert.NotEqual(t, nil, pageInfo)
	assert.Equal(t, 5, len(pageInfo.PostList))
}
