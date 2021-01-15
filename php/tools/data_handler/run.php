<?php

/**
 * Class Run
 *
 * 运行步骤：
 * 1. 导出上一周csv文件，放到data目录下
        select `user_id`,
        `store_name`,
        `create_time`
        FROM `gp_order` as o
        where o.`user_id` in (
        SELECT distinct o.`user_id`
        FROM `gp_order` as o
        where o.`create_time` BETWEEN '2021-01-03 00:00:00'
        and '2021-01-09 23:59:59'
        and o.`status`= 3)
 *  2. 运行脚本
 *      php run.php 开始日期(Y-m-d) 结束日期   文件名
 *      如： php run.php 2021-01-03 2021-01-09 ./data/sqlresult_5540750.csv
 *
 */

class Run
{
    protected $filename;
    protected $file_handler;
    protected $timeArr;

    public function __construct(string $filename, int $beginTime, int $endTime)
    {
        $this->filename = $filename;

        if (! file_exists($this->filename)) {
            throw new Exception('not found file:' . $this->filename);
        }

        $this->file_handler = fopen($this->filename, 'r');
        $this->timeArr['begin'] = $beginTime;
        $this->timeArr['end'] = $endTime;
    }

    protected function getData()
    {
        while ($row = fgetcsv($this->file_handler)) {
            yield $row;
        }
    }

    // 门店ID
    protected function storeInfo()
    {
        return [
            494 => '天河沃凯街店', 495 => '天河保利中宇店',
            496 => '天河金海花园店', 497 => '海珠区叠景中路店',
            500 => '海珠区愉景南苑店', 501 => '海珠区纵横广场店',
            552 => '海珠区仲恺店'
        ];
    }

    // 获取新户数 判断注册时间是否在上一周，在上一周时间内属于新户
    public function getNumberUser()
    {

        $data = [];
        foreach ($this->getData() as $key => $values) {
            if ($key == 0) {
                continue;
            }

            $data[$values[0]][] = ['time' => $values[1] ?? 'wrong', 'store_id' => $values[2] ?? 'wrong'];

        }

        $oldUserIds = [];
        // 新户数
        $newUserNumbers = 0;

        // 老客户数
        $oldUserNumbers = 0;

//        $buyVipKey = iconv('utf-8', 'gbk//IGNORE', '直接购买会员');
        $buyVipKey = '直接购买会员';

        $storeUserNumbser = [];
        // 初始化各门店新老客户数
        foreach (array_keys($this->storeInfo()) as $storeId) {
            $storeUserNumbser[$storeId]['new'] = 0;
            $storeUserNumbser[$storeId]['old'] = 0;
            $storeUserNumbser[$buyVipKey]['new'] = 0;
            $storeUserNumbser[$buyVipKey]['old'] = 0;
        }

        foreach ($data as $uid => $values) {
            $minOrderTime = strtotime(min(array_column($values, 'time')));

            // 计算上周新老客户人数
            if ($minOrderTime >= $this->timeArr['begin'] && $minOrderTime <= $this->timeArr['end']) {
                $newUserNumbers ++;

                // 统计各门店新增客户人数
                foreach (array_column($values, 'store_id') as $storeId) {
                    $storeUserNumbser[$storeId ? $storeId : $buyVipKey]['new'] ++;
                }

            } else {
                array_push($oldUserIds, $uid);
                $oldUserNumbers ++;

                // 统计各门店老客户回店消费人数(回头：过滤掉非上一周的时间的)
                foreach ($values as $value) {

                    $time = strtotime($value['time']);

                    if ($time >= $this->timeArr['begin'] && $time <= $this->timeArr['end']) {

                        $storeUserNumbser[$value['store_id'] ? $value['store_id'] : $buyVipKey]['old'] ++;
                    }
                }
            }

        }

        // 将$storeUserNumbser中门店ID替换成门店名字
        foreach ($this->storeInfo() as $storeId => $storeName) {
            $storeUserNumbser[$storeName] = $storeUserNumbser[$storeId];
//            $storeUserNumbser[iconv('utf-8', 'gbk//IGNORE', $storeName)] = $storeUserNumbser[$storeId];
            unset($storeUserNumbser[$storeId]);
        }

        return [
            'new' => $newUserNumbers,
            'old' => $oldUserNumbers,
//            'oldUIDs' => join(',', $oldUserIds),
            // 门店数据
            'store' => $storeUserNumbser
        ];

    }

    public static function showDataByTable($data, $beginTime, $endTime)
    {
        $str = <<<EOT
        <html>
        <head>
<meta charset="UTF-8">
</head>
<body style="margin: 100px;">
<h2>上周各门店数据客户数据统计({$beginTime}至{$endTime})</h2>
<h3>上周新增客户人数共{$data['new']}人，上周老客户再次消费共{$data['old']}人。</h3>
<table border="1" cellpadding="0" cellspacing="0" style="text-align: center;">
<tr>
    <td></td>
    <td>新增客户人数</td>
    <td>老客户再次消费人数</td>
</tr>
EOT;

        foreach ($data['store'] as $key => $value) {
            $str .= <<<EOT

    <tr>
        <td>{$key}</td>
        <td>{$value['new']}</td>
        <td>{$value['old']}</td>
    </tr>

EOT;

        }

        $str .= '</table></body></html>';

        header('content-type:text/html;charset=utf-8');
        file_put_contents('./data/' . date('Y-m-d') . '.html', $str);

    }

    public function __destruct()
    {
        fclose($this->file_handler);
    }
}


if (count($argv) < 4) {
    throw new Exception('wrong params');
}

$beginTime = $argv[1] . ' 00:00:00';
$endTime = $argv[2] . ' 23:59:59';

$result = (new Run(
    $argv[3],
    strtotime($beginTime),
    strtotime($endTime)
))->getNumberUser();


Run::showDataByTable($result, $beginTime, $endTime);
