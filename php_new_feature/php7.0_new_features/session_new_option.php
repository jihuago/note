<?php

/*

session_start()可以接收一个array作为参数，用来覆盖php.ini文件中设置的会话配置选项。

在调用session_start()的时候，传入的选项参数中也支持session.lazy_write行为，默认情况下这个配置是打开的。

session.lazy_write选项的作用是控制PHP只有在会话中的数据发生变化的时候才写入会话存储文件，如果会话中的数据没有发生改变，那么php会在读取完会话数据之后，立即关闭会话存储文件，不做任何修改，可以通过设置read_and_close来实现。

*/


session_start([
	'read_and_close' => true,
	'save_handler' => 'mysql',
]);