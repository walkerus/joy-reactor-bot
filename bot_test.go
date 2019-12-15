package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"joy-reactor/pkgs"
	"strings"
	"testing"
)

func TestGetPostId(t *testing.T)  {
	expectedId := `12345`
	reader := strings.NewReader(fmt.Sprintf("<a title=\"ссылка на пост\" class=\"link\" href=\"/post/%s\">ссылка</a>", expectedId))
	id, _ := pkgs.GetPostId(reader)

	assert.Equal(t, expectedId, id)

	reader = strings.NewReader(fmt.Sprintf(`<a title="ссылка на пост" class="link" href="/post/">ссылка</a>`))
	id, err := pkgs.GetPostId(reader)

	assert.Equal(t, ``, id)
	assert.Equal(t, `post not found`, err.Error())
}
