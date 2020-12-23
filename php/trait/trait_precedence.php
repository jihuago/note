<?php

/*

优先级

从父类继承的成员方法会被trait插入的成员覆盖。

优先顺序：

1. 当前类的成员
2. trait的成员
3. 从父类继承的成员
 */

class Parents
{
	public function demo()
	{
		print('parent demo');
	}
}

trait Shop
{
	public function demo()
	{
		print('trait demo');
	}
}

class Child extends Parents
{
	use Shop;
}

(new Child)->demo();//输出 trait demo,因为trait的demo覆盖了父类的demo方法