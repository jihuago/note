<?php

/*

trait冲突的解决

如果两个trait都插入了一个同名的方法，如果没有明确解决冲突将会产生一个致命错误。

为了解决多个trait在同一个类中的命名冲突，需要使用insteadof操作符来明确指定使用冲突方法中的哪一个。

通过as操作符为某个方法起别名。当冲突的方法都想使用，就需要使用到as操作符。
 */

trait Test
{

	public function say()
	{
		print('test');
	}
}

trait Demo
{
	//Demo与Test的say方法同名了
	public function say()
	{
		print('Demo');
	}
}

class A
{
	use Test, Demo {

		//指定使用Demo中的say方法
		//也可以理解为使用Demo的say方法替代了Test的say
		Demo::say insteadof Test;

		//如果也想使用Test的say方法，那就需要给Test的say方法起别名
		// 将Test的say方法起个别名
		Test::say as testSay;
	}

	public function aFunc()
	{
		print('A类的方法');
	}
}

$a = new A;
$a->say();//调用Demo的say方法
$a->testSay();//调用了Test的say方法
