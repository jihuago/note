<?php

$arr = [5, 4, 1, 6, 3, 2];
//$arr = [5, 4];

// 从小到大排序
function insertSort ($arr) {
    $len = count($arr);

    /*
     *  1.
     *      把4当成有序的，拿5与4对比  =>  5,4    拿$arr[1]
     *              $i = 1
     *              $j = 0
     *       $arr[0] > $arr[1] 真
     *            将$arr[0]元素后移，
     *            $arr[1] 插入到0下标
     *
     *      再拿1与4,5对比       拿$arr[2]，与$arr[1]，$arr[0]对比
     *              $i = 2
     *              $j = 0 => 1
     *
     */
    // [5, 4]
    for ($i = 1; $i < $len; $i++) { // 负责拿无序的元素 $i
        // 保存插入的值
        $value = $arr[$i];

        // [5, 4, 1, 6, 3, 2] => [1, 4, 5, 6, 3, 2]
        // 拿无序的元素倒序与有序的元素进行对比
        for ($j = $i - 1; $j >= 0; $j--) {// 因为$j 负责倒序拿有序的元素，当
            // 对比
            if ($arr[$j] > $value) {
                // 数据移动(后移)
                $arr[$j+1] = $arr[$j];
            } else {
                break;
            }

//            break 2;
        }

        echo 'i:',$i,'j:',$j,PHP_EOL;
        // 数据插入  $j + 1怎么得到的
        $arr[$j+1] = $value;

    }

    return $arr;

}

$res = insertSort($arr);
print_r($res);