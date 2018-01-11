<?php

//在匿名函数中返回值
$helloworld = function () { 
    return "each hello world is different... ".date("His"); 
}; 

echo $helloworld(); 