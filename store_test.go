package main

import (
	"github.com/stretchr/testify/assert"
	"joy-reactor/pkgs"
	"os"
	"testing"
)

func TestGetChats(t *testing.T)  {
	testFileName := `tests/test.json`
	text := `{"chats": [1, 2, 3]}`
	file, _ := os.Create(testFileName)

	defer os.Remove(testFileName)

	file.WriteString(text)
	file.Close()

	s := pkgs.Store{FileName: testFileName}
	ids, _ := s.GetChats()

	assert.Equal(t, []int64{1, 2, 3}, ids)
}

func TestGetChatsFileNotFound(t *testing.T)  {
	testFileName := `tests/test.json`

	s := pkgs.Store{FileName: testFileName}
	ids, err := s.GetChats()

	assert.Equal(t, []int64(nil), ids)
	assert.IsType(t, new(os.PathError), err)
}

func TestAddChat(t *testing.T)  {
	testFileName := `tests/test.json`
	file, _ := os.Create(testFileName)

	defer os.Remove(testFileName)

	file.WriteString(`{"chats": [1, 2, 3]}`)
	file.Close()

	s := pkgs.Store{FileName: testFileName}
	s.AddChat(4)
	ids, _ := s.GetChats()

	assert.Equal(t, []int64{1, 2, 3, 4}, ids)
}

func TestAddChatFirst(t *testing.T)  {
	testFileName := `tests/test.json`
	file, _ := os.Create(testFileName)

	defer os.Remove(testFileName)

	file.Close()

	s := pkgs.Store{FileName: testFileName}
	s.AddChat(1)
	ids, _ := s.GetChats()

	assert.Equal(t, []int64{1}, ids)
}

func TestAddChatNotUniqId(t *testing.T)  {
	testFileName := `tests/test.json`
	file, _ := os.Create(testFileName)

	defer os.Remove(testFileName)

	file.WriteString(`{"chats": [1, 2, 3]}`)
	file.Close()

	s := pkgs.Store{FileName: testFileName}
	s.AddChat(2)
	ids, _ := s.GetChats()

	assert.Equal(t, []int64{1, 2, 3}, ids)
}
