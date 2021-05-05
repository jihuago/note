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
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did nt expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}

		t.Run("return an error if a server doesn't respond within 10s", func(t *testing.T) {
			server := makeDelayedServer(25 * time.Millisecond)

			defer server.Close()

			_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

			if err == nil {
				t.Errorf("expected an error but didn't get one")
			}
		})
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server  {
	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(delay)
		writer.WriteHeader(http.StatusOK)
	}))
}