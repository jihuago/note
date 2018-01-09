<?php

/*

trait组成trait

正如class能够组成trait一样，其他trait也能都使用trait。

在trait定义时通过使用一个或多个trait,能够组合其他trait中的部分或全部成员。


 */

trait Hello
{
	public function sayHello()
	{
		print('Hello'.PHP_EOL);
	}
}

trait World
{
	public function sayWorld()
	{
		print('World'.PHP_EOL);
	}
}

trait A
{
	// trait A组合了trait Hello和Word
	use Hello, World;
}

class B
{
	use A;
}

$b = new B;
$b->sayHello();
$b->sayWorld();