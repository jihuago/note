package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayerServer(t *testing.T)  {
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)

		// net/http/httptest自带一个名为ResponseRecorder的监听器，它有很多有用的方法可以检查应答被写入了什么
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
}