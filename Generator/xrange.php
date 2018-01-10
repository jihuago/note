<?php

//使用生成器


//定义一个生成器
function xrange($start, $end, $step = 1)
{
    for ($i = $start; $i <= $end; $i += $step) {
        yield $i;
    }
}

foreach (xrange(0, 1000000) as $v) {
    echo $v.PHP_EOL;
}
