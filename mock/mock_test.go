package mock

import "testing"

func TestNewRequest(t *testing.T) {
	request := NewRequest()

	get, err := request.Get("www.baidu.com")
	if err != nil {
		t.Error(err.Error())
		return
	}
	content := get.Content()
	t.Log(content)

	user := new(struct {
		UserID uint64 `json:"userID"`
	})

	if err = get.MarshalJSON(user); err != nil {
		t.Fatal(err.Error())
		return
	}
}
