<?php


/*

修改方法的访问控制

使用as语法还可以用来调整方法的访问控制。但是定义在trait中方法的访问控制没有改变



 */


trait Test
{

	public function say()
	{
		print('say method is called'.PHP_EOL);
	}

	protected function demo()
	{
		print('demo method is called'.PHP_EOL);
	}

	private function run()
	{
		print('run method is called'.PHP_EOL);
	}
}

class A
{
	use Test {
		// 将Test的say方法访问控制由public改为protected
		say as protected;

		//将Test的demo方法访问控制由protected改为public
		demo as public;

		//将Test的run方法访问控制由private改为public,并且改名为newRun
		run as public newRun;
	}

}

$a = new A;

// $a->say();//报错，因为访问控制变成protected,类外无法调用
$a->demo();

$a->newRun();