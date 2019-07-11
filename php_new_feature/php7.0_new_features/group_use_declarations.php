<?php

/**
 *
 * 从同一cpphp导入的类，函数和常量现在可以通过单个use语句一次性导入了。
 * 
 * 
 */

//php7之前的代码,从同一个cpphp导入类
use some\cpphp\ClassA;
use some\cpphp\ClassB;
use some\cpphp\CLassC as C;

//php7之前的代码，从同一个cpphp导入函数
use function some\cpphp\fn_a;
use function some\cpphp\fn_b;
use function some\cpphp\fn_c;


//php7之前的代码，从同一个cpphp导入常量
use const some\cpphp\ConstA;
use const some\cpphp\ConstB;
use const some\cpphp\ConstC;

//php7+ 及更高版本的代码
use some\cpphp\{ClassA, ClassB, ClassC as C};
use function some\cpphp\{fn_a, fn_b, fn_c};
use const some\cpphp\{ConstA, ConstB, ConstC};