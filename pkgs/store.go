package pkgs

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Store struct {
	FileName string
}

func (store *Store) GetChats() ([]int64, error)  {
	jsonFile, err := os.Open(store.FileName)

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	if len(byteValue) == 0 {
		return make([]int64, 0), nil
	}

	var result map[string][]int64
	unmarshalError := json.Unmarshal([]byte(byteValue), &result)

	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return result[`chats`], nil
}

func (store *Store) AddChat(chatID int64) error  {
	chats, getChatsError := store.GetChats()

	if getChatsError != nil {
		return getChatsError
	}

	for _, v := range chats {
		if v == chatID {
			return nil
		}
	}

	result := make(map[string][]int64)
	result[`chats`] = chats
	jsonFile, err :=  os.OpenFile(store.FileName, os.O_RDWR | os.O_TRUNC, 0777)

	if err != nil {
		return err
	}

	defer jsonFile.Close()

	result[`chats`] = append(result[`chats`], chatID)
	newFileContent, marshallingError := json.Marshal(result)

	if marshallingError != nil {
		return marshallingError
	}

	_, writingError := jsonFile.Write(newFileContent)

	if writingError != nil {
		return writingError
	}

	return nil
}
