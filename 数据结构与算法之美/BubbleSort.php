<?php

/*
 *  冒泡排序
 *      0. 概念、原理
 *          如：有一个数组[1, 3, 2, 5, 6, 4]，要求从小到大排序。
 *          第一次冒泡： 1 3 对比，不交换；3 2 对比，交换 => [1, 2, 3, 5, 6, 4]
 *                      3 5 对比，不交换； 5 6 对比，不交换
 *                      6 4 对比，交换 => [1, 2, 3, 5, 4, 6] 结束第一次冒泡
 *          要做到一次冒泡，只需要循环一次即可：如bubbleOne()函数所示。
 *          第一次冒泡完后，我们得到的是[1, 2, 3, 5, 4, 6]，所以我们不可能只冒泡一次的，当然也有只冒泡一次就能得到符合要求的。
 *          所以，我们可能有多次冒泡，具体多少次，要看数据特点。但冒泡最大次数是不会超过数组元素个数的。
 *        * 如何让冒泡一直进行下去，直到符合排序要求
 *              按照上面说的，最大冒泡次数不超过数组长度；我们假设每个数组都需要进行最大次数的冒泡。那么想要让冒泡一直进行下去，
 * 直到符合排序要去，就只需要在bubbleOne()的循环外再来一次循环即可。
 *
 *      1. 时间复杂度： o(n²)
 *          一次冒泡的时间复杂度是O(n)，因为只需要进行一次循环。如果需要进行最大次数的冒泡，时间复杂度就是O(n²)。
 *          通常我们所说的时间复杂度是指最坏时间复杂度，所以时间复杂度是O(n²)
 *      2. 优化
 *          bubbleSort()代码需要优化。回想一下冒泡的过程：
 *              第一次冒泡结束后，得到的数组是[1, 2, 3, 5, 4, 6] ,经过第一次冒泡后，最大的元素已经在数组最右边了。
 *
 *              第二次冒泡过程应该是： 1  2， 2  3， 3  5， 5 4（交换） =》 [1, 2, 3, 4, 5, 6]，也就是说不需要进行
 *              下标 4 与 下标 5的元素对比了。此时： $j < 4 即可，也就是只要比较到 $arr[3],$arr[4]即可。
 *              这时： $j < 4 , 冒泡次数：第2次冒泡， $i = 1，所以： $j < 4，这个4 = $len - 1 - $i;
 *
 *              如果有第三次冒泡，同理：因为经过前两次冒泡，最大和次大的数都已经排好了。所以我们只需要比较到$arr[2],$arr[3]即可
 *               这时：$j < 3，第三次冒泡，$i = 2，运用上面公式 $len - 1 - $i = $len - 1 - 2 也是等于 3的
 *
 *          经过上面的优化，可以得到bubbleSortFinal()
 *
 */

// 只冒泡一次
$arr = [1, 3, 2, 5, 6, 4];
function bubbleOne ($arr) {
    $len = count($arr);

    for ($j = 0; $j < $len - 1; $j++) {

        if ($arr[$j] > $arr[$j+1]) {
            $tmp = $arr[$j];
            $arr[$j] = $arr[$j+1];
            $arr[$j+1] = $tmp;
        }
    }

    return $arr;
}
$res = bubbleOne($arr);
//print_r($res);

// =====================================
$arr = [1, 3, 2, 5, 6, 4];
function bubbleSort ($arr) {
    $len = count($arr);

    for ($i = 0; $i < $len; $i++) {
        for ($j = 0; $j < $len - 1; $j++) {

            // 交换
            if ($arr[$j] > $arr[$j+1]) {
                $tmp = $arr[$j];
                $arr[$j] = $arr[$j+1];
                $arr[$j+1] = $tmp;
            }
        }
    }

    return $arr;
}
$res = bubbleSort($arr);
print_r($res);

//==================================
$arr = [1, 3, 2, 5, 6, 4];
//$arr = ['a', 'b', 'd', 'c', 'f', 'e'];
/**
 * 冒泡排序最终版
 * @param array $arr
 * @return array|void
 */
function bubbleSortFinal (array $arr) {

    $len = count($arr);

    // 只有一个元素，不必进行排序
    if ($len <= 1) {
        return;
    }

    for ($i = 0; $i < $len; ++$i) {
        $flag = false;

        /*
         *   $j < $Len - 1 - $i 是如何得到的：
         *
         *    $len = 6   < 5   $i = 0
         *    $len = 6   < 4   $i = 1  因为上一次后，最后一个肯定是最大的了。
         */
        for ($j = 0; $j < $len - 1 - $i; ++$j) {

            // 进行数据交换
            if ($arr[$j] > $arr[$j+1]) {
                $tmp = $arr[$j];
                $arr[$j] = $arr[$j+1];
                $arr[$j+1] = $tmp;

                // 有数据交换
                $flag = true;
            }
        }

        // 第一轮冒泡就已经知道接下来的冒泡有没有需要交换的数据，所以如果第一轮没有交换，意味着剩下的也不必进行了。
        if (! $flag) {
            break;
        }
    }

    return $arr;
}
$res = bubbleSortFinal($arr);

//print_r($res);

