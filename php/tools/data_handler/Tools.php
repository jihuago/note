<?php

/**
 * 工具函数类
 * Trait Tools
 */
trait Tools
{
    // 根据二维数组中某个key排序
    public function array_sort($array, $on, $order=SORT_ASC)
    {
        $new_array = array();
        $sortable_array = array();

        if (count($array) > 0) {
            foreach ($array as $k => $v) {
                if (is_array($v)) {
                    foreach ($v as $k2 => $v2) {
                        if ($k2 == $on) {
                            $sortable_array[$k] = $v2;
                        }
                    }
                } else {
                    $sortable_array[$k] = $v;
                }
            }

            switch ($order) {
                case SORT_ASC:
                    asort($sortable_array);
                    break;
                case SORT_DESC:
                    arsort($sortable_array);
                    break;
            }

            foreach ($sortable_array as $k => $v) {
                $new_array[$k] = $array[$k];
            }
        }

        return $new_array;
    }

    public function toCSV(array $data, array $colHeaders = array(), $asString = false)
    {
        $stream = ($asString)
            ? fopen("php://temp/maxmemory", "w+")
            : fopen("php://output", "w");

        if (!empty($colHeaders)) {
            fputcsv($stream, $colHeaders);
        }

        foreach ($data as $record) {
            fputcsv($stream, $record);
        }

        if ($asString) {
            rewind($stream);
            $returnVal = stream_get_contents($stream);
            fclose($stream);
            return $returnVal;
        }
        else {
            fclose($stream);
        }
    }

    // 获取测试单数据
    public function getTestOrders($filename = './data/test_orders.txt')
    {
        $result = '';
        if (file_exists($filename)) {
            $result = file_get_contents($filename);
        }

        return $result;
    }
}
