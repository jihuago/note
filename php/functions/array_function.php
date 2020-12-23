<?php

/** 
 * 数组函数：
 *  1. array array_change_key_case(array $array [, int $case = CASE_LOWER])
 *     将数组中的所有键名修改为全大写或小写。本函数不改变数字索引
 *     $array 需要操作的数组
 *     $case 可以在这里用两个常量，CASE_UPPER或CASE_LOWER 
 *  2. array array_chunk(array $array, int $size [, bool $preserve_keys = false])
 *     将一个数组分割成多个数组，其中每个数组的单元数组由size决定。最后一个数组的单元数组可能会少于size个 
 *     $array 需要操作的数组
 *     $size  每个数组的单元数组
 *     $preserve_keys 设为TRUE,可以使PHP保留输入数组中原来的键名。如果你指定了False,那每个结果数组将用从零开始的新数字索引
 *
 *  3. array array_column(array $input, mixed $column_key [, mixed $index_key = null])
 *     返回数组中指定的一列， array_column()返回input数组中键值为column_key的列，如果指定了可选参数index_key，那么input数组中的这一列的值将作为返回数组中对应值得键
 *
 *  4. array array_combine(array $keys, array $values)
 *     创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值
 *
 *     返回值
 *          返回合并后的array,如果两个数组的单元数不同则返回FALSE
 *  5. array array_count_values(array $array)
 *
 *     统计数组中所有的值出现的次数
 *
 *     返回值：
 *
 *       返回一个关联数组，用array数组中的值作为键名，该值在数组中出现的次数作为值
 */


//array_change_key_case()函数例子
$names = array(
    '姓名' => 'jack',
    '性别' => '男',
);

$names = array(
    'name' => 'jack',
    'sex' => '男',
);
print_r(array_change_key_case($names, CASE_UPPER));
echo PHP_EOL;

//============================================
//array_chunk()函数例子
$chunk = array('a', 'b', 'c', 'last' => 'd');

echo '<pre>';
print_r(array_chunk($chunk, 3));
echo PHP_EOL;

//保留键名
echo '<pre>';
print_r(array_chunk($chunk, 3, true));
echo PHP_EOL;


//=============================================
//array_column()函数例子
$data = array(
    array(
        'name' => 'jack',
        'sex'  => '男',
    ),
    array(
        'name' => 'mary',
        'sex'  => '女',
    ),
);

//取出所有name键的值
echo '<pre>';
print_r(array_column($data, 'name'));

//取出所有name键的值,并且用相应的sex的值作为键
echo '<pre>';
print_r(array_column($data, 'name', 'sex'));


//=============================================
//array_combine()函数的例子
$a = array('green', 'red', 'yellow');
$b = array('apple', 'banana', 'orange');

echo '<pre>';
//$a数组的值作为键，$b作为新数组的值
print_r(array_combine($a, $b));


//$a与$b的个数一样
$a = array('green', 'red', 'name' => 'yellow');
$b = array('apple', 'banana', 'orange');

echo '<pre>';
//$a数组的值作为键，$b作为新数组的值
print_r(array_combine($a, $b));


//如果两个键相同，第二个生效
echo '<pre>';
print_r(array_combine(array('a', 'a', 'b'), array(1, 2, 3)));

/*
 *  返回值：
 *  array(
 *      'a' => 2
 *      'b' => 3
 *  )
 *
 */


//===============================================
//array_count_values()函数例子
$array = array(1, 'hello', 1, 'world');

//数组里面的每个不是string和integer类型的元素抛出一个警告错误
//$array = array(1, 'hello', 1, 'world', ['a']);

echo '<pre>';
print_r(array_count_values($array));
