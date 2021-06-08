<?php

/*

// 计算所有P3等级的发型师的综合复购率。不区分洗剪吹、烫染护

// 需要去掉测试单

// 总客数
SELECT
	count(
	DISTINCT ( user_id ) )
FROM
	gp_order AS o
WHERE
	create_time BETWEEN '2021-01-01 00:00:00'
	AND '2021-03-30 23:59:59'
	AND employee_id = 5688
	AND `status` = 3

// 复购人数：取行数
SELECT count(*),user_id from gp_order as o where 	employee_id = 5688
	AND `status` = 3 and create_time <= '2021-03-31' GROUP BY user_id HAVING count(*) >= 2

// 分母：一种只是拿1-3月内的， 一种拿3月前所有的

*/

class CalCulatRepeat
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

    public function count()
    {
/*        $sql = <<<EOT
SELECT
	count(
	DISTINCT ( user_id ) ) as number
FROM
	gp_order AS o 
WHERE
	create_time BETWEEN '2021-03-01 00:00:00' 
	AND '2021-05-30 23:59:59' 
	AND employee_id = ? and order_no not in({$this->getTestOrders()})
	AND `status` = 3
EOT;*/

        // 重复购买客户数量/客户样本数量。

        // 客户样本数
        $sql = <<<EOT
SELECT
	count(
	DISTINCT ( user_id ) ) as number
FROM
	gp_order AS o 
WHERE
	create_time <= '2021-05-31 23:59:59'
	AND employee_id = ? and order_no not in({$this->getTestOrders()})
	AND `status` = 3
EOT;

        // 重复购买客户数量
        $sql1 = <<<EOT
SELECT
	count(*),
	user_id 
FROM
	gp_order AS o 
WHERE
	employee_id = ? 
	AND `status` = 3 
	AND create_time <= '2021-05-31 23:59:59'  AND order_no NOT IN 
	({$this->getTestOrders()}) 
	GROUP BY user_id HAVING count(*) >= 2
EOT;

        $data = [];
        foreach ($this->p3LevelHairCuters() as $key => $hairCuter) {

            // 排除测试号发型师
            if (
                in_array($hairCuter['id'], array_keys($this->getTestHairCuters()))
            ||
                in_array($hairCuter['name'], ['Tony'])
            ) {
                continue;
            }

            $allUserNumber = $this->db->findOneBySql($sql, [$hairCuter['id']]);

            $data[$hairCuter['id']]['number'] = $allUserNumber['number'];

            $data[$hairCuter['id']]['repeatUser'] = count($this->db->findAllBySql($sql1, [$hairCuter['id']]));

            $data[$hairCuter['id']]['repeatUser'] = $data[$hairCuter['id']]['repeatUser'] ? $data[$hairCuter['id']]['repeatUser'] : 1;

            $data[$hairCuter['id']]['repurchase_rate'] =
                round($data[$hairCuter['id']]['repeatUser'] / $data[$hairCuter['id']]['number'], 3) * 100 . '%';

            $employer = $this->db->table('employee')
                ->field(['name'])
                ->where(['id' => $hairCuter['id']])
                ->findOne();

            $data[$hairCuter['id']]['name'] = $employer['name'];

        }

//        var_dump($data);exit;

        $str = $this->toCSV($data, ['客户数', '重复消费客户数', '复购率', '发型师'], true);

        date_default_timezone_set("PRC");
        file_put_contents('./data/fugou'.date('Y-m-dHis').'.csv', $str);
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

    // 获取p3等级的发型师
    public function p3LevelHairCuters():array
    {
        $results = $this->db->table('employee as e')
            ->field(['id', 'name'])
            ->where([
                'e.position' => 1,
                'status' => 1,
                'deleted' => 0
            ])
            ->whereIn('job_status', [1, 2])
            ->whereIn('e.role', [1, 2])
            ->findAll();

        return $results;
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
}

(new CalCulatRepeat())->count();