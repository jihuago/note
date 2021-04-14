package LearnMaps

import "testing"

// map类型的键类型只能是一个可比较的类型，因为如果不能判断两个键是否相等
// 值的类型可以是任意类型，甚至可以是另一个map

// Map是引用类型，就像指针一样。
// 引用类型引入了maps可以是nil值，如果尝试使用一个nil的map，会得到一个nil指针异常
// 由于nil指针异常，所以不应该初始化一个空的map变量。如 var m map[string][string]

// 可以这样初始化空map 或使用make创建map
// d = map[string]string{}   d = make(map[string]string)
func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	got, _ := dictionary.Search("test")
	want := "this is just a test"

	assertStrings(t, got, want)
}

func assertStrings(t *testing.T, got, want string)  {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
