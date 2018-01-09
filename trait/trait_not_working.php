<?php


/*

这个代码总结了trait一些错误的用法，以及一些可用，但不常见的用法



 */


/*class B
{

}

// 1. trait继承一个类(错误)
trait A extends B
{

}*/

//=============================
// 2. 实例化trait是错误的
/*trait Test
{
	public static function demo()
	{
		return 'demo method is calling';
	}

}

// new Test;//错误

//但是可以直接调用trait的静态方法
print(Test::demo());*/

//=================================
//也可以在trait中定义一个抽象方法，然后组合类实现
trait Test
{
	public static function demo()
	{
		return 'demo method is calling';
	}

	abstract public function db();
}

class AA
{
	use Test;

	//实现trait中的抽象方法db
	public function db()
	{
		return 'db';
	}
}

//调用db方法
print (new AA)->db();