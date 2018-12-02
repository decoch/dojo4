package handler_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/decoch/dojo4/kadai4/decoch/application"
	"github.com/decoch/dojo4/kadai4/decoch/handler"
	"github.com/decoch/dojo4/kadai4/decoch/model"
)

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	mock := omikujiServiceMock{}
	omikujiHandler := handler.OmikujiHandler{
		Service: &mock,
	}
	omikujiHandler.ServeHTTP(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}
	var response map[string]string
	if err = json.Unmarshal(b, &response); err != nil {
		t.Fatal("unexpected error")
	}
	if response["result"] != "中吉" {
		t.Fatalf("unexpected response: %v", response)
	}
}

type omikujiServiceMock struct {
	application.OmikujiService
}

func (o *omikujiServiceMock) Get() model.FortuneSlip {
	return model.FortuneSlip{
		Result: model.Chukichi,
	}
}
