<?php

// 匿名函数对事件处理程序非常有用,下面使用匿名函数模拟一下Laravel框架的路由实现。
error_reporting(0);
class Event
{

	public static $events = [];

	//绑定
	public static function bind($event, $callback, $obj = null)
	{
		if (self::$events[$event]) {

			self::$events[$event] = [];
		}

		//将匿名函数，或对象与方法存入属性
		self::$events[$event][] = ($obj === null) ? $callback : [$obj, $callback];
	}

	//执行相关事件函数
	public static function run($event)
	{
		if (!self::$events[$event]) return;

		foreach (self::$events[$event] as $callback) {

			if (call_user_func($callback) === false) break;
		}
	}
}



//绑定
Event::bind('test', function () {

	echo 'hi,test', PHP_EOL;
});



//重复绑定test。之后绑定的覆盖之前的
Event::bind('test', function () {

	echo 'hi,aa', PHP_EOL;
});


Event::run('test');


//调用某个类的某个方法
class Index
{
	public function showIndex()
	{
		print('this is Class Index Method');
	}
}

Event::bind('index', 'showIndex', new Index);
Event::run('index');
