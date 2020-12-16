<?php

/*

整数除法函数intdiv()

int intdiv ( int $dividend , int $divisor )
dividend 除以 divisor 的商，对该商取整

如果传递的参数是浮点数，则先拿浮点数的整数部分，然后再做除法运算
例如： intdiv(12.8, 3.2);  => intdiv(12, 3)  => 4

新加的函数intdiv()用来进行整数的除法运算。

手册： http://php.net/manual/zh/function.intdiv.php

*/

var_dump(intdiv(10, 3));//int(3)

//相当于12/3
var_dump(intdiv(12.8, 3));//int(4)