package concurrency

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// 下面测试代码存在的问题：
// 1. 在理想情况下不能依赖外部服务来进行测试，因为可能速度慢，不可靠，无法进行边界条件测试
// 建议是改用模拟测试，这样就可以控制可靠的服务器来测试了
/*func TestRacer(t *testing.T)  {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.co.uk"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}*/

func TestRacer(t *testing.T)  {
	slowServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(20 * time.Millisecond)
		writer.WriteHeader(http.StatusOK)
	}))

	defer slowServer.Close()

	fastServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}))

	defer fastServer.Close()

	slowURL := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got := Racer(slowURL, fastUrl)
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}


}