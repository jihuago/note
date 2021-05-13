<?php



// 取出消费大于等于2次的uids
// 1. 执行5月份的sql, 然后遍历结果，根据结果分别判断每月有无服务，如果当月有服务，则数据保留

class CountRepeatBuy
{

    private  $db;

    public function __construct()
    {
        require __DIR__ . '/vendor/autoload.php';

        $config = array(
            'dsn' => 'mysql:host=192.168.10.10;dbname=fecmall',
            'username' => 'root',
            'password' => '123456',
            'charset' => 'utf8',
            'tablePrefix' => 'gp_',
        );

        $this->db = new \PFinal\Database\Builder($config);
    }


    public  function count()
    {

        $beginTime = '2021-05-01 00:00:00';
        $endTime = '2021-05-31 23:59:59';

        $serveArr = $this->bigProjectLists();// 大项目
//        $serveArr = $this->littleProjectLists();// 小项目

        $serveIDs = implode(',', $serveArr);

        $sql = <<<EOT
SELECT (count(o.id) - 1) as times,
       o.`user_id`,o.store_name,e.`name`,e.stage_name,e.id
  from `gp_order` as o LEFT JOIN `gp_employee` as e on o.`employee_id` = e.`id`
 where o.`status`= 3
   and o.`create_time`<= '{$endTime}'
   and o.`serve_id` IN({$serveIDs}) and o.`order_no` not in ({$this->getTestOrders()})  GROUP BY o.`user_id`,o
   .`store_name` ,o
   .`employee_id`
HAVING count(o.id)>=2
EOT;

//        echo $sql;exit;
// 如果两个都是同一个月份的，数量应该减一
        $data = $this->db->findAllBySql($sql);
//        var_dump($data);exit;
        foreach ($data as $key => $value) {
            $currentMonthNumber = $this->db->table('order as o')
                ->where('o.employee_id = ?', [$value['id']])
                ->where('o.user_id = ?', [$value['user_id']])
                ->where('o.create_time >= ? and o.create_time <= ?', [$beginTime, $endTime])
                ->where('o.status = ?', [3])
                ->whereIn('o.serve_id', $serveArr)
                ->count();

            if ($currentMonthNumber == 0) {
                unset($data[$key]);
                continue;
            }

            if ($currentMonthNumber > 1) {

                $beforetMonthNumber = $this->db->table('order as o')
                                     ->where('o.employee_id = ?', [$value['id']])
                                     ->where('o.user_id = ?', [$value['user_id']])
                                     ->where('o.create_time <= ?', [$beginTime])
                                     ->where('o.status = ?', [3])
                                     ->whereIn('o.serve_id', $serveArr)
                                     ->count();

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

    public function getTestOrders($filename = './data/test_orders.txt')
    {
        $result = '';
        if (file_exists($filename)) {
            $result = file_get_contents($filename);
        }

        return $result;
    }

    public function bigProjectLists()
    {
        return [2,3,8,10,11,12,13,16,17,21,22,23,24,25,26,27,28,29,30,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47];
    }

    // 小项目的ID
    public function littleProjectLists()
    {
        return [1,14,15,18,19,20,31,48];
    }

    public function toCSV(array $data, array $colHeaders = array(), $asString = false) {
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
