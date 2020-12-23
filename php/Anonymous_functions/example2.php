<?php

/*
闭包函数也可以作为变量的值来使用。

php会自动把此种表达式转换成内置类Closure的对象实例。

把一个closure对象赋值给一个变量的方式与普通变量赋值的语法是一样的，
最后也要加上分号。

 */
// 匿名函数变量赋值示例

$greet = function ($name) 
{
	print($name);
};

$greet('jack');

//===============================

/*
闭包可以从父作用域中继承变量。

任何此类变量都应该用use语言结构传递进去。

PHP7.1起，不能传入此类变量：1. 超全局变量 2. $this 3. 参数重名

 */

/*$message = 'hello';

//没有 'use'
$example = function ()
{
	var_dump($message);//NULL
};
$example();

//使用 'use'
$example = function () use ($message) 
{
	var_dump($message);//'hello'
};

$example();*/

// 这些变量都必须在函数或类的头部声明。

//错误的写法
$fruits = ['apples', 'oranges'];
$example = function () use ($fruits[0]) {
    echo $fruits[0]; 
};
$example();

//正确
$fruits = ['apples', 'oranges'];
$example = function () use ($fruits) {
    echo $fruits[0]; // will echo 'apples'
};
$example();

//正确
$fruits = ['apples', 'oranges'];
$fruit = $fruits[0];
$example = function () use ($fruit) {
    echo $fruit; // will echo 'apples'
};
$example();