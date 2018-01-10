<?php

//使用生成器


//定义一个生成器
/*function xrange($start, $end, $step = 1)
{
    for ($i = $start; $i <= $end; $i += $step) {
        yield $i;
    }
}

//每次foreach才去yield值，而每次yield之后，内存不会释放，变量$i
//值会保留，下一次foreach就是$i++的值
foreach (xrange(0, 1000000) as $v) {
    echo $v.PHP_EOL;
}*/

//=================
//下面代码也是实现range()函数的效果，但是性能低

function makeRange($len)
{
	$dataset = [];

	for ($i = 0; $i < $len; $i++) {
		$dataset[] = $i;
	}

	return $dataset;
}

//这种写法代码存在的问题：
//1. 如果传递给makeRange函数的值很大，就需要在内存中开辟一块很大的空间存放$dataset
//2. makeRange()的return $dataset,因为makeRange()执行完后，栈会释放，必须将$dataset复制到其他栈，如果传递给makeRange的值很大，内存复制就非常耗时。
//
foreach (makeRange(1000000) as $v) {

	echo $v, PHP_EOL;
}
