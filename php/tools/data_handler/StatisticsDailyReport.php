<?php

require_once('./BaseModel.php');
require_once('./Tools.php');

/*

	俏手艺门店数据说明
		数据统计日期：2021.6.1-6.6

		1、门店总营业额：30291元
		      洗剪吹：18848元
		      烫染直：8114元
		      护理：915元
		      套餐次卡：1836元
		      会员：576元

		2、总单数：514单
		      洗剪吹：472单
		      烫染直：33单
		      护理：7单
		      套餐次卡：2单

		3、业绩/订单排名
		     业绩前五
		     李子文：3736元
		     李飞：2210元
		     孙守法：2149元
		     朱志兵：1990元
		     陈海龙：1824元

		     订单前五
		     李子文：44单
		     孙守法：42单
		     江志峰：40单
		     朱志兵：34单
		     李源：33单

		4、昨日销售情况
		营业额：
		纵横广场店：1056元
		曡景中路店：1090元
		保利中宇店：759元
		愉景南苑店：284元
		仲恺店：39元
		江怡路店：772元
		叠彩园店：319元

		订单数：
		纵横广场店：16单
		曡景中路店：21单
		保利中宇店：16单
		愉景南苑店：7单
		仲恺店：1单
		江怡路店：20单
		叠彩园店：11单

		5、美团数据(6.1-6.6)
		总业绩(未去除平台扣点及代运营公司分成)
		纵横广场店：99元
		曡景中路店：0元
		保利中宇店：633元
		愉景南苑店：238元
		仲恺店：158元
		江怡路店：98元
		叠彩园店：0元
		总计：1227元

		6、周环比增长
		(5.24-5.31 与 6.1-6.6)对比
		门店业绩增长：
		叠彩园店：208.54%
		愉景南苑店：18.36%
		曡景中路店：-14.29%
		保利中宇店：-26.69%
		纵横广场店：-36.59%
		江怡路店：-40.6%
		仲恺店：-72.95%

		7、个人业绩增长：
		“正增长”前五
		①陈海龙，128.86%
		②江志峰，123.77%
		③齐晶晶，40.75%
		④李飞，36.94%
		⑤李源，6.6%

		“负增长”前五
		①李登，-94.37%
		②孙童，-50.95%
		③罗立讯，-48%
		④陈华相，-47.88%
		⑤向涛，-47.3%

		8、分析
		①六月第一周同比相比业绩从37148元下降到30291元，掉了18.46%
		②订单数同比下降25.87%，从692单下降至513单
		③六月第一周业绩31441元，相对于五月第一周的39943元，下降21.29%
		④总体业绩/订单都是下滑状态，疫情影响！办理会员也同比下降62%

		以上数据具体拆分已制作成图片，发送至上方


 */

/**
 *
 * php StatisticsDailyReport.php 2021-06-1(累计开始时间) 2021-06-10(累计结束时间) 2021-06-10(要统计的单日，如昨日)
 *
 * Class StatisticsDailyReport
 */
class StatisticsDailyReport extends BaseModel
{
    protected $startDate;
    protected $endDate;
    protected $yestody;

    use Tools;

    public function __construct($startDate, $endDate, $yestoday)
    {
        $this->startDate = $startDate;
        $this->endDate = $endDate;
        $this->yestody = $yestoday;

        parent::__construct();
    }

    public function run()
    {
        $this->outPutString();
    }

    private function calcuBaseModel(string $startDate, string $endDate)
    {
        return $this->db->table('order as o')
            ->where([
                'status' => 3,
            ])
            ->where('create_time between ? and ?', [$startDate, $endDate]);

    }

    // 统计周期内的： 1、 门店总营业额 2、 总单数 3、 业绩/订单排名
    protected function calcuStoreData(string $startDate, string $endDate)
    {
        $data = [];

        // 门店总营业额
        $sum = $this->calcuBaseModel($startDate, $endDate)->sum('pay_price');
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
                $data[$type_name]['monery'] = $this->calcuBaseModel($startDate, $endDate)
                    ->where([
                        $fieldName => $type_key
                    ])->sum('pay_price');

                $data[$type_name]['monery'] /= 100;

                $data[$type_name]['number'] = $this->calcuBaseModel($startDate, $endDate)
                    ->where([
                        $fieldName => $type_key
                    ])->count();
            }

        }

        // 统计发型师订单数、业绩
        $hairCuterData = $this->calcuHairCuterSaleData($startDate, $endDate);
        $data = array_merge($hairCuterData, $data);

        $data['业绩排名']['number'] = array_slice($this->array_sort($data['发型师'], 'number', SORT_DESC), 0, 5);
        $data['业绩排名']['price'] =  array_slice($this->array_sort($data['发型师'], 'price', SORT_DESC),0, 5);
        unset($data['发型师']);

        return $data;
    }

    // 统计发型师订单数、业绩
    protected function calcuHairCuterSaleData($startDate, $endDate):array
    {
        $data = [];

        foreach ($this->hairCuters() as $hairCuter) {
            $data['发型师'][$hairCuter['name']]['number'] = $this->calcuBaseModel($startDate, $endDate)
                ->where([
                    'employee_id' => $hairCuter['id']
                ])->count();

            $data['发型师'][$hairCuter['name']]['price'] = $this->calcuBaseModel($startDate, $endDate)
                ->where([
                    'employee_id' => $hairCuter['id']
                ])->sum('pay_price');

            $data['发型师'][$hairCuter['name']]['price'] /= 100;
        }

        return $data;
    }

    // 各各门店销售情况
    protected function calcuEveryStoreData($startDate, $endDate)
    {
        $data = [];
        foreach ($this->storeInfo() as $storeID => $storeName) {
            $data['order_number'][$storeName] = $this->calcuBaseModel($startDate, $endDate)->where(['store_id' => $storeID])->count();
            $data['order_money'][$storeName] = $this->calcuBaseModel($startDate, $endDate)->where(['store_id' => $storeID])->sum('pay_price') / 100;
        }

        return $data;
    }

    // 计算周各门店业绩环比
    private function calcuGrowth()
    {
        $growth = [];
        foreach ($this->lastWeekPeriod() as $key => $arr) {
            $growth[$key] = $this->calcuEveryStoreData($arr['start'], $arr['end'])['order_money'];
        }

        $result = [];
        foreach (array_values($this->storeInfo()) as $storeName) {
            $tmp = $growth['two'][$storeName] > 0? $growth['two'][$storeName] : 1;
            $result[$storeName] = round(
                ($growth['last'][$storeName] - $growth['two'][$storeName]) / $tmp * 100,
                2
                ) . '%';
        }


        return $result;
    }

    // 周环比增长率文本
    protected function growthStr()
    {
        $dateStr = '';
        $i = 0;
        foreach ($this->lastWeekPeriod() as $dates) {
            $dateStr .= $dates['start'] . '-' . $dates['end'] . '与';
            $i++;
        }

        $dateStr = trim($dateStr, '与');

        $str = "6、周环比增长 \r\n （{$dateStr}）对比 \r\n";
        foreach ($this->calcuGrowth() as $storeName => $item) {
            $str .= $storeName . ':' . $item . "\r\n";
        }

        return $str;
    }

    // 发型师个人业绩增长
    public function hairCuterGrowth()
    {
        $growth = [];
        foreach ($this->lastWeekPeriod() as $key => $arr) {
            $growth[$key] = $this->calcuHairCuterSaleData($arr['start'], $arr['end'])['发型师'];
        }

        $data = [];
        foreach ($this->hairCuters() as $hairCuter) {
            $tmp = $growth['two'][$hairCuter['name']]['price'] > 0 ? $growth['two'][$hairCuter['name']]['price'] : 1;
            $data[$hairCuter['name']] = round(($growth['last'][$hairCuter['name']]['price'] - $growth['two'][$hairCuter['name']]['price']) / $tmp * 100, 2);

            if ($data[$hairCuter['name']] > 100) {
                unset($data[$hairCuter['name']]);
                continue;
            }
            $data[$hairCuter['name']] .= '%';

        }

        return $data;
    }

    // 发型师个人业绩前五正负增长字符串
    public function hairCuterGrowthStr()
    {

        $data = $this->hairCuterGrowth();
        asort($data, SORT_NUMERIC);

        $str = "7、个人业绩增长：\r\n 正增长前五 \r\n";

        // 负增长
        $negativeGrowth = array_slice($data, 0, 5);
        // 正增长
        $positiveGrowth = array_slice($data, -5);

        $i = 1;
        foreach ($positiveGrowth as $hairCuter => $value) {
            $str .= "\t" . $i . '. ' . $hairCuter . ':' . $value . "\r\n";
            $i++;
        }

        $str .= "负增长前五\n\r";
        $i = 1;
        foreach ($negativeGrowth as $hairCuter => $value) {
            $str .= "\t" .  $i . '. ' . $hairCuter . ':' . $value . "\r\n";
            $i++;
        }

        return $str;
    }

    protected function outPutString()
    {
        $storeData = $this->calcuStoreData($this->startDate, $this->endDate);
        $yestodayData = $this->calcuEveryStoreData($this->yestody . ' 00:00:00', $this->yestody . ' 23:59:59');

        $this->calcuGrowth();

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

        $tmpStartDate = date('Y-m-d', strtotime($this->startDate));
        $tmpEndDate = date('Y-m-d', strtotime($this->endDate));
        $str = <<<EOT
俏手艺门店数据说明
数据统计日期：{$tmpStartDate}至{$tmpEndDate}

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
{$this->growthStr()}
{$this->hairCuterGrowthStr()}
EOT;

        echo $str;

    }

    /**
     * 上一周的开始，结束日期
     *
     * @return array[]
     */
    public function lastWeekPeriod():array
    {

        $twoWeekAgoStartDate = date('Y-m-d', strtotime('-2 week', strtotime($this->endDate)));
        $twoWeekAgoEndDate = date('Y-m-d', strtotime('+6 day', strtotime($twoWeekAgoStartDate)));

        $lastWeekStartDate = date('Y-m-d', strtotime('-1 week', strtotime($this->endDate)));
        $lastWeekAgoEndDate = date('Y-m-d', strtotime('+6 day', strtotime($lastWeekStartDate)));

        return [
            'two' => [
                'start' => $twoWeekAgoStartDate,
                'end' => $twoWeekAgoEndDate,
            ],
            'last' => [
                'start' => $lastWeekStartDate,
                'end' => $lastWeekAgoEndDate
            ],
        ];
    }

}

$beginTime = $argv[1] . ' 00:00:00';
$endTime = $argv[2] . ' 23:59:59';
$yestoday = $argv[3];

(new StatisticsDailyReport($beginTime, $endTime, $yestoday))->run();
//(new StatisticsDailyReport($beginTime, $endTime, $yestoday))->hairCuterGrowthStr();
