<?php

    //php7增加了对返回类型声明的支持，类似参数类型声明
    //返回类型声明指明了函数返回值的类型
    //可用的类型与参数声明中可用的类型相同。
    function getName($index):string
    {

        $namesArr = ['jack', 'mary'];

        if ($index <0 || $index > count($namesArr)-1 ) {
            throw new Exception('下标不在范围');        
        }
        return $namesArr[$index];
    }

    echo getName(0);

    class Test
    {
        //getName()方法的返回值限定为string
        public function getName():string
        {
            return 'jack';
        }
    }

    $test = new Test;
    echo $test->getName();
