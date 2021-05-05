package concurrency

// version 1

//type WebsiteChecker func(string) bool


/*func CheckWebsites(wc WebsiteChecker, urls []string) map[string] bool {
	results := make(map[string]bool)


	for _, url := range urls {


				以下代码存在一个问题：
					* 可能存在两个goroutine完全同时写入results map。Go的maps不喜欢多个事物视图一次性写入，可能会导致fatal error
					* 这是一种race condition(竞争条件)，当软件的输出取决于事件发生的时间和顺序时，因为我们无法控制，bug就会出现。
					* 因为我们无法准确控制每个goroutine写入结果map的时间，两个goroutine同一时间写入时程序将非常脆弱
					* 通过内置的race detector来发现竞争条件
						go test -race

					* 通过channels协调我们的goroutines来解决数据竞争
						channel是一个Go数据结构，可以同时接受和发送值。这些操作以及细节允许不同进程之间的通信。



		go func(u string) {
			results[u] = wc(u)
		}(url)
	}



	time.Sleep(2 * time.Second)
	return results
}*/

//
// version 2
// channels 用来组织和控制不同进程之间的交流，使我们能够避免race condition的问题
// the race detector(竞争探测器)帮助我们调试并发代码的问题
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, webSites []string) map[string]bool {
	results := make(map[string]bool)

	// channel
	resultsChannel := make(chan result)

	for _, url := range webSites {

		// 避免map竞争条件，放channel处理
		go func(url string) {
			resultsChannel <- result{url, wc(url)}
		}(url)
	}

	for i := 0; i < len(webSites); i++ {
		result1 := <- resultsChannel

		results[result1.string] = result1.bool
	}

	return results
}
