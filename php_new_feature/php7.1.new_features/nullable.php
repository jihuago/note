<?php

/**
 *
 * 可为空(Nullable)类型
 *
 * 类型现在允许为空，当启用这个特性时，传入的参数或者函数返回的结果要么是给定的类型，要么是null。
 *
 * 可以通过在类型前面加上一个问号来使之成为可为空的。
 */

//只能给$name传递null或者是string类型的值
function test(?string $name)
{
	var_dump($name);
}

test();//报错Uncaught Error: Too few arguments to function test(), 0 passed in...

test(null);//ok
test('ff');//ok