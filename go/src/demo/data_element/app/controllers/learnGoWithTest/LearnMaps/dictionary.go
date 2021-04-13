package LearnMaps

import "errors"

// 创建了一个Dictionary类型，它是对map的简单封装。
type Dictionary map[string]string

var (
	ErrNotFound = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)


func (d Dictionary) Search(word string) (string, error)  {

	// map查找的一个特性，可以返回两个值。第二个值是一个布尔值，表示是否成功找到key
	// 通过这个特性，可以区分单词不存在还是未定义
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
			d[word] = definition
	case nil:
		return ErrWordExists

	default:
		return err
	}

	return nil
}
