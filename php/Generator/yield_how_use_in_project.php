<?php

/*

迭代大型数据集或数列时最适合使用生成器，因为这样占用的系统内存最少。


一个常用的使用案例就是使用生成器迭代流资源（文件、音频等）。假设我们想要迭代一个大小为4GB的CSV文件，而虚拟私有服务器（VPS）只允许PHP使用1GB内存，因此不能把整个文件都加载到内存中，下面的代码展示了如何使用生成器完成这种操作

*/
function getRows($file) {
    $handle = fopen($file, 'rb');
    if ($handle == FALSE) {
        throw new Exception();
    }
    while (feof($handle) === FALSE) {
        yield fgetcsv($handle);
    }
    fclose($handle);
}

foreach ($getRows($file) as $row) {
    print_r($row);
}

//上述示例一次只会为CSV文件中的一行分配内存，而不会把整个4GB的CSV文件都读取到内存中。