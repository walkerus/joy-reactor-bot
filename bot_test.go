package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"joy-reactor/pkgs"
	"strings"
	"testing"
)

func TestGetPostId(t *testing.T)  {
	expectedID := `12345`
	reader := strings.NewReader(fmt.Sprintf("<a title=\"ссылка на пост\" class=\"link\" href=\"/post/%s\">ссылка</a>", expectedID))
	id, _ := pkgs.GetPostID(reader)

	assert.Equal(t, expectedID, id)

	reader = strings.NewReader(fmt.Sprintf(`<a title="ссылка на пост" class="link" href="/post/">ссылка</a>`))
	id, err := pkgs.GetPostID(reader)

	assert.Equal(t, ``, id)
	assert.Equal(t, `post not found`, err.Error())
}
