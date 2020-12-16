<?php

/*
	通过define()定义常量数组：

	Array类型的常量在php7.0后可以通过define()来定义。在php5.6中仅能通过const定义。

 */

define('ANIMALS', [
	'dog',
	'cat',
	'bird'
]);

echo ANIMALS[1];//cat