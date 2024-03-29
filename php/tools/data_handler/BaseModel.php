<?php


class BaseModel
{
    protected  $db;

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

    // 测试号发型师
    protected function getTestHairCuters(): array
    {
        return [
            5682 => '李泽',
            5681 => '张晓一',
            5678 => '梁宁',
            5646 => '李辉',
            5676 => '陈龙',
            5672 => '范宇',
            5675 => '钟伟健',
            5645 => '陈华',
            5644 => '阮赞益',
            5701 => '许华明'
        ];
    }

    // 发型师列表
    protected function hairCuters():array
    {
        $hairCuters = $this->db->table('employee as e')
            ->field(['id', 'name', 'hairdresser_store_id'])
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

    protected function cardLevel():array
    {
        return [
            4 => '钻石',
            3 => '铂金',
            2 => '黄金'
        ];
    }

    // 门店ID
    protected function storeInfo()
    {
        return [
            495 => '天河保利中宇店',
            553 => '海珠江怡路店',
            497 => '叠景中路店',
//            500 => '海珠区愉景南苑店',
            501 => '纵横广场店',
//            552 => '海珠区仲恺店',
            556 => '叠彩园店',
            555 => '康乐西路店',
            557 => '天河棠东店',
            558 => '赤岗店',
        ];
    }

    /*

    Array
(
    [天河保利中宇店] => 113188
    [海珠江怡路店] => 90097.65
    [叠景中路店] => 143376
    [纵横广场店] => 152776
    [叠彩园店] => 83112
    [康乐西路店] => 40856
    [天河棠东店] => 35524
    [赤岗店] => 35200
    [金海] => 5344
    [系统] => 1500
)



     */
}