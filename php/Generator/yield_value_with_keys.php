<?php

/*

指定键名来生成值

PHP的数组支持关联键值对数组，生成器也一样支持。

所以除了生成简单的值，你也可以生成值得时候指定键名。

如下面例子所示，生成一个键值对与定义一个关联数组十分相似

 */


function yieldData($data)
{
	foreach ($data as $k => $v) {

		//这样指定键名来生成值
		yield $k => $v;
	}
}

echo '<pre>';

foreach (yieldData(['id' => 1]) as $id => $val) {

	echo $id,'=>',$val;
}
