<?php

/*

// 需要去掉测试单

算法： 单位时间内，复购人数 / 总客数
比如，4-5月复购率。 总客数 = 4-5月总客数    复购人数 = （4-5月消费次数>=2的人数）  +  （4-5月消费次数等于1，且之前有过消费的）

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

*/

class CalculateRepurchaseRate
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

    // 统计期间的客户样本数量
    protected function getConsumeNumber($employeeId)
    {
//        $startTime = '2021-04-01 00:00:00';
//        $endTime = '2021-05-31 23:59:59';
        $startTime = '2021-01-01 00:00:00';
        $endTime = '2021-03-31 23:59:59';

        $sql = <<<EOT
SELECT
	count(
	DISTINCT ( user_id ) ) as number
FROM
	gp_order AS o 
WHERE
	create_time between '{$startTime}'
	AND '{$endTime}'
	AND employee_id = ? and order_no not in({$this->getTestOrders()})
	AND `status` = 3
EOT;

        // 取4 5    1 3 月份
        $allUserNumber = $this->db->findOneBySql($sql, [$employeeId]);

        return $allUserNumber['number'] ?? 1;
    }

    // 获取复购人数 = （4-5月消费次数>=2的人数）  +  （4-5月消费次数等于1，且之前有过消费的）
    protected function getRepeatNumber($employeeId)
    {
        $this->getRepeatNumberGt2($employeeId);
    }

    // （4-5月消费次数>=2的人数）
    private function getRepeatNumberGt2($employeeId)
    {
        $startTime = '2021-01-01 00:00:00';
        $endTime = '2021-03-31 23:59:59';

        $sql = <<<EOT
SELECT
	count(*) as number,
	user_id
FROM
	gp_order AS o 
WHERE
	employee_id = ? 
	AND `status` = 3
	AND create_time between '{$startTime}'
	AND '{$endTime}'
	AND order_no NOT IN ({$this->getTestOrders()}) 
	GROUP BY user_id HAVING count(*) >= 2
EOT;

        $repeatNumberGt2 = $this->db->findOneBySql($sql, [$employeeId]);

        return $repeatNumberGt2['number'] ?? 0;
//        var_dump($repeatNumberGt2);exit;
//        return count($repeatNumberGt2);
    }

    // （4-5月消费次数等于1，且之前有过消费的）
    private function getRepeatNumberEq1($employeeId)
    {
        $startTime = '2021-01-01 00:00:00';
        $endTime = '2021-03-31 23:59:59';

        $sql = <<<EOT
SELECT
	count(*) as number,
	user_id
FROM
	gp_order AS o 
WHERE
	employee_id = ? 
	AND `status` = 3
	AND create_time between '{$startTime}'
	AND '{$endTime}'
	AND order_no NOT IN ({$this->getTestOrders()}) 
	GROUP BY user_id HAVING count(*) = 1
EOT;

        // 4-5月消费次数等于1的
        $repeatNumbers = $this->db->findAllBySql($sql, [$employeeId]);

        // 之前有过消费的
        $sqlBefore = <<<EOT
SELECT
	user_id
FROM
	gp_order AS o 
WHERE
	employee_id = ? 
	AND `status` = 3
	AND create_time < '{$startTime}'
	AND user_id = ?
	AND order_no NOT IN ({$this->getTestOrders()})
EOT;

        foreach ($repeatNumbers as $key => $value) {
            $result = $this->db->findOneBySql($sqlBefore, [$employeeId, $value['user_id']]);

            if (empty($result['user_id'])) {
                unset($repeatNumbers[$key]);
            }
        }

        return count($repeatNumbers);
    }

    public function countFinal()
    {
        $hairCuters = $this->p3LevelHairCuters();

        foreach ($hairCuters as $key => $hairCuter) {

            // 排除测试号发型师
            if (
                in_array($hairCuter['id'], array_keys($this->getTestHairCuters()))
            ) {
                unset($hairCuters[$key]);
                continue;
            }

            // 统计期间的客户样本数量
            $hairCuters[$key]['consumeNumber'] = $this->getConsumeNumber($hairCuter['id']);

            $hairCuters[$key]['repeatNumberGt2'] = $this->getRepeatNumberGt2($hairCuter['id']);
            $hairCuters[$key]['repeatNumberEq1'] = $this->getRepeatNumberEq1($hairCuter['id']);
            $repeatNumberTotal = ($hairCuters[$key]['repeatNumberGt2'] + $hairCuters[$key]['repeatNumberEq1']);
            $hairCuters[$key]['repurchaseRate'] = round($repeatNumberTotal / $hairCuters[$key]['consumeNumber'], 3) * 100 . '%';

        }

        $str = $this->toCSV($hairCuters, ['发型师ID', '发型师', '客户数', '重复消费客户数2', '重复消费客户数1', '复购率'], true);

        date_default_timezone_set("PRC");
        file_put_contents('./data/综合复购率'.date('Y-m-dHis').'.csv', $str);
//        var_dump($hairCuters);
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
//                'e.position' => 1,
                'status' => 1,
                'deleted' => 0
            ])
            ->whereIn('job_status', [1, 2])
            ->whereIn('e.role', [1, 2])
            ->where('name != ?', 'Tony')
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

(new CalculateRepurchaseRate())->countFinal();