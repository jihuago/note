package controllers

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	// 结构体tag
	Name string `json:"name"`

	// omitempty 序列化结果去掉空值字段
	Email string `json:"email,omitempty"`
	Password string `json:"password"`
}

func OmitemptyDemo()  {
	u1 := User{Name: "jack"}

	// json序列化
	b, err := json.Marshal(u1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("str:%s\n", b)

	intAndStringDemo()

	// 使用自定义序列化方法
	o1 := Order{
		ID: 123,
		Title: "《测试》",
		CreatedTime: time.Now(),
	}
	//r, err := o1.MarshalJSON()
	r, err := json.Marshal(&o1)
	if err != nil {
		fmt.Printf("json.Marshal o1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", r)

}

// 优雅处理字符串格式的数字
// 有时候，前端在传递来的json数据中可能会使用字符串类型的数字，这个时候可以在结构体tag中添加string告诉json包从字符串中解析相应字段的数据
type Card struct {
	ID int64 `json:"id,string"`
	Score float64 `json:"score,int"`
}
func intAndStringDemo() {
	jsonStr1 := `{"id":"12345", "score":88.9}`
	var c1 Card
	if err := json.Unmarshal([]byte(jsonStr1), &c1); err != nil {
		fmt.Printf("json.Unmarshal jsonStr1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("c1:%#v\n", c1)
}

// 不修改原结构体且忽略空值字段
// 需要json序列化User，但是不想把密码也序列化，又不想修改User结构体
type PublicUser struct {
	*User // 匿名嵌套
	Password *struct{} `json:"password,omitempty"`
}

// 自定义MarshalJSON和UnmarshalJSON方法
type Order struct {
	ID int `json:"id"`
	Title string `json:"title"`
	CreatedTime time.Time `json:"created_time"`
}
const layout = "2006-01-02 15:04:05"
// 为Order类型实现自定义的MarshalJSON方法
// 如果你能够为某个类型实现了MarshalJSON()([]byte, error) 和 UnmarshalJSON(b []byte) error 方法，
// 那么这个类型在序列化和反序列化时就会使用你定制的相应方法
func (o *Order) MarshalJSON() ([]byte, error) {
	type TempOrder Order // 定义与Order字段一致的新类型

	return json.Marshal(struct {
		CreatedTime string `json:"created_time"`
		*TempOrder // 避免直接嵌套Order进入死循环
	}{
		CreatedTime: o.CreatedTime.Format(layout),
		TempOrder: (*TempOrder)(o),
	})
}

// UnmarshalJSON为order
func (o *Order) UnmarshalJSON(data []byte) error  {
	type TempOrder Order
	ot := struct {
		CreatedTime string `json:"created_time"`
		*TempOrder
	}{
		TempOrder: (*TempOrder)(o),
	}

	if err := json.Unmarshal(data, &ot); err != nil {
		return err
	}

	var err error
	o.CreatedTime, err = time.Parse(layout, ot.CreatedTime)
	if err != nil {
		return err
	}
	return nil
}

// 使用匿名结构体添加字段
// 使用内嵌结构体能够扩展结构体的字段，但有时候没有必要单独定义新的结构体，可以使用匿名结构体简化操作
type UserInfo struct {
	ID int `json:"id"`
	Name string `json:"name"`
}
func anonymousStructDemo() {
	u1 := UserInfo{
		ID: 123456,
		Name: "起腻",
	}
	// 使用匿名结构体内嵌User并添加额外字段token
	b, err := json.Marshal(struct {
		*UserInfo
		Token string `json:"token"`
	}{
		&u1,
		"fdsalfjl",
	})

	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}

	fmt.Printf("str:%s", b)

}

