<?php

/*
 * 计算发型师复购次数脚本
 *    如计算5月复购次数
 *      1. 5月前无消费：复购次数 = 5月总单数 - 1
 *      2. 5月前有消费：复购次数 = 5月总单数
 */

class CountRepeatBuy
{

    private  $db;

    public function __construct()
    {
        require __DIR__ . '/vendor/autoload.php';

        $config = array(
            'dsn' => 'mysql:host=192.168.10.10;dbname=gipin',
            'username' => 'homestead',
            'password' => 'secret',
            'charset' => 'utf8',
            'tablePrefix' => 'gp_',
        );

        $this->db = new \PFinal\Database\Builder($config);
    }


    protected function countTime()
    {
        return [
            []
        ];
    }

    public function count()
    {

        $beginTime = '2021-05-01 00:00:00';
        $endTime = '2021-05-31 23:59:59';

        $serveArr = $this->bigProjectLists();// 大项目
//                $serveArr = $this->littleProjectLists();// 小项目

        $serveIDs = implode(',', $serveArr);

        $sql = <<<EOT
SELECT (count(o.id) - 1) as times,
       o.`user_id`,o.store_name,e.`name`,e.stage_name,e.id
  from `gp_order` as o LEFT JOIN `gp_employee` as e on o.`employee_id` = e.`id`
 where o.`status`= 3
   and o.`create_time`<= '{$endTime}'
   and o.`serve_id` IN({$serveIDs}) and o.`order_no` not in ({$this->getTestOrders()})  GROUP BY o.`user_id`,
   o.`store_name`,
   o.`employee_id`
HAVING count(o.id)>=2
EOT;

        $data = $this->db->findAllBySql($sql);

        foreach ($data as $key => $value) {
            $sql = <<<EOT
SELECT count(*) as number
  from `gp_order` as o
 where o.`employee_id` = {$value['id']}
   and o.`user_id` = {$value['user_id']}
   and o.`create_time` >= '{$beginTime}' AND o.`create_time`<= '{$endTime}'
   and o.`status`= 3
   and o.`serve_id` IN({$serveIDs}) 
   and o.`order_no` not in ({$this->getTestOrders()})
EOT;

            $currentMonthNumber = $this->db->findOneBySql($sql)['number'];

            if ($currentMonthNumber == 0) {
                unset($data[$key]);
                continue;
            }

            if ($currentMonthNumber > 1) {
                $sql = <<<EPT
SELECT count(*) as beforenumber
  from `gp_order` as o
 where o.`employee_id` = {$value['id']}
   and o.`user_id` = {$value['user_id']}
   and o.`create_time`<= '{$beginTime}'
   and o.`status`= 3
   and o.`serve_id` IN({$serveIDs}) 
   and o.`order_no` not in ({$this->getTestOrders()})
EPT;

                $beforetMonthNumber = $this->db->findOneBySql($sql)['beforenumber'];

                // 说明此前没有订单:当月总单-1
                if ($beforetMonthNumber == 0) {
                    $data[$key]['times'] = $currentMonthNumber - 1;
                } else {
                    // 此前有订单，当月也有订单：直接统计当月所有订单总数
                    $data[$key]['times'] = $currentMonthNumber;
                }



            } else if ($currentMonthNumber == 1) {
                $data[$key]['times'] = 1;
            }


        }

        $str = $this->toCSV($data, ['次数', '用户id', '门店', '发型师', '艺名', '发型师id'], true);

        date_default_timezone_set("PRC");
        file_put_contents('./data/'.date('Y-m-dHis').'.csv', $str);

    }

    protected function getTestOrders($filename = './data/test_orders.txt')
    {
        $result = '';
        if (file_exists($filename)) {
            $result = file_get_contents($filename);
        }

        return $result;
    }

    protected function bigProjectLists()
    {
        return [
            2,3,8,10,11,12,13,16,17,21,22,23,24,25,26,27,
            28,29,30,32,33,34,35,36,37,38,39,40,41,42,43,
            44,45,46,47,52,53,54,55,56,57,58,59,60,61,62,
            63,64,65,66,67,68,69,70,71,72,73,74,75,76,77,
            78,79,80
        ];
    }

    // 小项目的ID
    protected function littleProjectLists()
    {
        return [1,14,15,18,19,20,31,48,49,50,51];
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
}



(new CountRepeatBuy())->count();
