<?php

/*

trait的静态成员

trait可以定义静态属性、静态方法
 */

trait StaticExample {
	// 在trait中定义静态属性
	public static $name = 'jack';

	// 在trait中定义成员属性
	public $sex = 'woman';

	public $test = 1;

	// 在trait中定义静态方法
    public static function doSomething() 
    {
        return 'Doing something';
    }
}

class Example {
    use StaticExample;

    public $test = 1;//不报错，与trait $test属性可见度，默认值一样

    // Trait 定义了一个属性后，类就不能定义同样名称的属性，否则会产生 fatal error。 有种情况例外：同样的访问可见度、初始默认值
    // public $sex;
}

// 调用trait静态方法
echo Example::doSomething().PHP_EOL;

// 调用trait静态属性
echo Example::$name.PHP_EOL;

// 访问成员属性sex
print (new Example)->sex;