<?php


$redis = new Redis();
$redis->connect('localhost', 6379);



// $res = $redis->set('test11', 'aa', array('nx', 'ex' => 1000));
$res = $redis->set('test22', 'bb', array('nx', 'px' => 10000));

var_dump($res);