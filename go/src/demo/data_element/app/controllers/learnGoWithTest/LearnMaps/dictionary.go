package LearnMaps

// 创建了一个Dictionary类型，它是对map的简单封装。
type Dictionary map[string]string

const (
	ErrNotFound = DictionaryErr("could not find the word you were looing for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// 将错误声明为常量，可以让错误更具可重用性和不可变性
type DictionaryErr string

func (e DictionaryErr) Error() string  {
	return string(e)
}

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

func (d Dictionary) Update(word, definition string) error  {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string)  {
	delete(d, word)
}
