<?php

/*

yield关键字

生成器函数的核心是yield关键字。

它最简单的调用形式看起来像一个return声明，不同之处在于return返回值
会终止函数的执行，而yield会返回一个值给循环调用此生成器的代码并且只是暂停执行生成器函数。


 */

// 一个简单的生成器的例子
function getSomeNumber()
{
	for ($i = 0; $i < 3; $i++) {
		yield $i;
	}
}

//使用生成器
foreach (getSomeNumber() as $v) {
	echo $v,PHP_EOL;
}

/*

以上例子会输出：

1
2
3
*/