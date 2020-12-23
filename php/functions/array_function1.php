<?php


/**
 * array array_diff_assoc(array $arry1, array $arry2 [, array $...])
 *  带索引检查计算数组的差集，array_diff_assoc()返回一个数组，该数组包括了所有在
 *  array1中但是不在任何其他参数数组中的值，注意该函数键名也用于比较
 *
 * array array_diff_key(array $arry1, array $array [, array $...])
 * 使用键名比较数组的差集，根据array1中的键名和array2比较，返回不同键名的项,并且键的类型需要一样
 *
 * array array_diff_uassoc(array $array1, array $array2 [, array $...], callable $func)
 * 用用户提供的回调函数做索引检查来计算数组的差集,键名也会比较
 *
 * $func
 * 	在第一个参数小于，等于或大于第二个参数时，该比较函数必须相应地返回一个小于，等于或大于0的整数
 *
 * array array_diff(array $array1, array $array2 [, array $...])
 *  只根据值来计算数组的差集,返回存在$array1中但不存在其他array里的值
 *  
 * 
 */

//array_diff_assoc()例子
print_r(array_diff_assoc(
	['a' => 'green', 'b' => 'brown', 'c' => 'blue', 'red', 'sex' => 1],
	['a' => 'green', 'yellow', 'red', 'sex' => '1']
));

/*
输出：
因为'a'=>'green'在两个数组都有，因此不在输出中，与此不同的，键值对0=>'red'出现在输出中是因为第二个参数中'red'的键名是1。'sex'=>1输出，是因为array_diff_assoc使用了严格检查，字符串的表达式必须相同
array(
	'b' => 'brown',
	'c' => 'blue',
	'0' => 'red',
	'sex' => 1
)


 */