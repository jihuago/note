<?php

/**
 *
 * php StatisticsDailyReport.php 2021-05-30(累计开始时间) 2021-05-31(累计结束时间) 2021-05-30(要统计的单日，如昨日)
 *
 * Class StatisticsDailyReport
 */
class StatisticsDailyReport
{
    protected $startDate;
    protected $endDate;
    protected $yestody;
    private  $db;

    public function __construct($startDate, $endDate, $yestoday)
    {
        $this->startDate = $startDate;
        $this->endDate = $endDate;
        $this->yestody = $yestoday;

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

    protected function getTestOrders($filename = './data/test_orders.txt')
    {
        $result = '';
        if (file_exists($filename)) {
            $result = file_get_contents($filename);
        }

        return $result;
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

    // 测试号发型师
    private function getTestHairCuters(): array
    {
        return [
            5682 => '李泽',
            5681 => '张晓一',
            5678 => '梁宁',
            5646 => '李辉',
            5676 => '陈龙',
            5672 => '范宇',
            5675 => '钟伟健',
        ];
    }

    public function hairCuters():array
    {
        $hairCuters = $this->db->table('employee as e')
            ->field(['id', 'name'])
            ->where([
//                'e.position' => 1,
                'status' => 1,
                'deleted' => 0
            ])
            ->whereIn('job_status', [1, 2])
            ->whereIn('e.role', [1, 2])
            ->where('name != ?', 'Tony')
            ->findAll();


        foreach ($hairCuters as $key => $hairCuter) {
            if (in_array($hairCuter['id'], array_keys($this->getTestHairCuters()))) {
                unset($hairCuters[$key]);
            }
        }

        return $hairCuters;
    }

    // 门店ID
    protected function storeInfo()
    {
        return [
            495 => '天河保利中宇店',
            553 => '海珠江怡路店',
            497 => '海珠区叠景中路店',
            500 => '海珠区愉景南苑店',
            501 => '海珠区纵横广场店',
            552 => '海珠区仲恺店',
            556 => '叠彩园店',
        ];
    }

    public function run()
    {
        $this->outPutString();
    }

    // 统计周期内的： 1、 门店总营业额 2、 总单数 3、 业绩/订单排名
    protected function calcuStoreData()
    {
        $data = [];
        $model = $this->db->table('order as o')
            ->where([
                'status' => 3,
            ])
            ->where('create_time between ? and ?', [$this->startDate, $this->endDate]);

        // 门店总营业额
        $sum = $model->sum('pay_price');
        $sum /= 100;

        $data['门店总营业额'] = $sum;

        // 服务类型
        $serve_realated_types = [
            'serve_realated_type' => [
                1 => '洗剪吹',
                2 => '烫染直发',
                3 => '护理',
            ],
            'type' => [
                2 => '购买会员',
                3 => '购买次卡',
            ],
        ];

        foreach ($serve_realated_types as $fieldName => $serve_realated_type) {

            foreach ($serve_realated_type as $type_key => $type_name) {
                $data[$type_name]['monery'] = $this->db->table('order as o')
                    ->where([
                        'status' => 3,
                    ])
                    ->where('create_time between ? and ?', [$this->startDate, $this->endDate])
                    ->where([
                        $fieldName => $type_key
                    ])->sum('pay_price');

                $data[$type_name]['monery'] /= 100;

                $data[$type_name]['number'] = $this->db->table('order as o')
                    ->where([
                        'status' => 3,
                    ])
                    ->where('create_time between ? and ?', [$this->startDate, $this->endDate])
                    ->where([
                        $fieldName => $type_key
                    ])->count();
            }

        }

        // 统计发型师订单数
        foreach ($this->hairCuters() as $hairCuter) {
            $data['发型师'][$hairCuter['name']]['number'] = $this->db->table('order as o')
                ->where([
                    'status' => 3,
                ])
                ->where('create_time between ? and ?', [$this->startDate, $this->endDate])
                ->where([
                    'employee_id' => $hairCuter['id']
                ])->count();

            $data['发型师'][$hairCuter['name']]['price'] = $this->db->table('order as o')
                ->where([
                    'status' => 3,
                ])
                ->where('create_time between ? and ?', [$this->startDate, $this->endDate])
                ->where([
                    'employee_id' => $hairCuter['id']
                ])->sum('pay_price');

            $data['发型师'][$hairCuter['name']]['price'] /= 100;
        }

        $data['业绩排名']['number'] = array_slice($this->array_sort($data['发型师'], 'number', SORT_DESC), 0, 5);
        $data['业绩排名']['price'] =  array_slice($this->array_sort($data['发型师'], 'price', SORT_DESC),0, 5);
        unset($data['发型师']);
//        var_dump($data);exit;

        return $data;
    }

    // 昨日各门店销售情况
    public function calcuYestodayData()
    {
        $data = [];
        foreach ($this->storeInfo() as $storeID =>  $storeName) {
            $data['order_number'][$storeName] = $this->getTotalModel()->where(['store_id' => $storeID])->count();
            $data['order_money'][$storeName] = $this->getTotalModel()->where(['store_id' => $storeID])->sum('pay_price') / 100;
        }

        return $data;
    }

    protected function getTotalModel()
    {
        $model = $this->db->table('order as o')
            ->where([
                'status' => 3,
            ])
            ->where('create_time between ? and ?', [$this->yestody . ' 00:00:00', $this->yestody . ' 23:59:59']);

        return $model;
    }

    protected function outPutString()
    {
        $storeData = $this->calcuStoreData();
        $yestodayData = $this->calcuYestodayData();

        $numberRankingStr = "订单前五\r\n";
        foreach ($storeData['业绩排名']['number'] as $hairCuterName => $value) {
            $numberRankingStr .= "\t\t" . $hairCuterName . ':' . $value['number'] . "单\r\n";
        }

        $moneyRankingStr = "业绩前五\r\n";
        foreach ($storeData['业绩排名']['price'] as $hairCuterName => $value) {
            $moneyRankingStr .= "\t\t" . $hairCuterName . ':' . $value['price'] . "元\r\n";
        }

        // 昨日销售情况
        $yestodaySaleStr = "营业额：\r\n";
        foreach ($yestodayData['order_money'] as $storeName => $money) {
            $yestodaySaleStr .= "\t" . $storeName . ':' . $money . "元\r\n";
        }

        $yestodayOrderStr = "订单数：\r\n";
        foreach ($yestodayData['order_number'] as $storeName => $number) {
            $yestodayOrderStr .= "\t" . $storeName . ':' . $number . "单\r\n";
        }

        $str = <<<EOT
俏手艺门店数据说明
数据统计日期：{$this->startDate}-{$this->endDate}

1、门店总营业额：{$storeData['门店总营业额']}元
      洗剪吹：{$storeData['洗剪吹']['monery']}元
      烫染直：{$storeData['烫染直发']['monery']}元
      护理：{$storeData['护理']['monery']}元
      套餐次卡：{$storeData['购买次卡']['monery']}元
      会员：{$storeData['购买会员']['monery']}元
2、总单数：725单
      洗剪吹：{$storeData['洗剪吹']['number']}单
      烫染直：{$storeData['烫染直发']['number']}单
      护理：{$storeData['护理']['number']}单
      套餐次卡：{$storeData['购买次卡']['number']}单

3、业绩/订单排名
\t{$moneyRankingStr}
\t{$numberRankingStr}

4、昨日销售情况
{$yestodaySaleStr}
{$yestodayOrderStr}
EOT;

        echo $str;

    }

    protected function array_sort($array, $on, $order=SORT_ASC)
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

}

$beginTime = $argv[1] . ' 00:00:00';
$endTime = $argv[2] . ' 23:59:59';
$yestoday = $argv[3];

(new StatisticsDailyReport($beginTime, $endTime, $yestoday))->run();
