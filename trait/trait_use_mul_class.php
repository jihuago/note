<?php

/**
 * 多个trait
 * 通过逗号分隔，在use声明列出多个trait,可以都插入到一个类中
 */

trait Test
{

	public function say()
	{
		print(1);
	}

	public static function staticMethod()
	{
		print('calling trait staticMethod');
	}
}

trait Demo
{
	public function run()
	{
		print('run');
	}
}

//在A类中需要Test类与Demo类的方法
class A
{
	// 将Test,Demo类的方法插入到A类中，A类的实例可以调用Test,Demo类的所有方法
	use Test,Demo;

	public function aFunc()
	{
		print('A类的方法');
	}
}

$a = new A;
$a->say();//调用Test的say方法
$a->run();//调用Demo类的run方法
$a->aFunc();//自身的类也可以调用

echo PHP_EOL;

// 可以直接调用trait的静态方法
Test::staticMethod();

