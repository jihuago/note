<?php

/*
	由于开发中存在大量同时使用三元表达式和isset()的情况，php7添加了null合并运算符??这个语法糖。
	?? 如果变量存在且值不为NULL,它就会返回自身的值，否则返回它的第二个操作数。

	语法：

	判断 ?? 操作数2
*/

//当$_GET['id']不存在或值为NULL,$username = 1;
$username = $_GET['id'] ?? '1';

//下面代码等同于上面的代码
$username = isset($_GET['user']) ? $_GET['user'] : 'nobody';

var_dump($username);

echo '<br/>';

$a = '';

$b = $a ?? 'a';

var_dump($b);//''

//通过这个例子看，??不等同于?:三元运算符
$c = $a ?: 'a';
var_dump($c);//$c = 'a'
