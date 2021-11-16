<?php

require_once('./BaseModel.php');
require_once('./Tools.php');

// 根据用户表来计算每个店铺剩余的价格
class StatisticsFinal extends BaseModel
{
    use Tools;

    public function caculateUserMoney()
    {

        $storeNames = [
            495 => 0,
            553 => 0,
            497 => 0,
            501 => 0,
            556 => 0,
            555 => 0,
            557 => 0,
            558 => 0,
            496 => 0,
            0 => 0, // 系统
        ];

        $this->db->table("user u")
            ->field(['id', 'balance'])
            ->where('balance > ?', 0)
            ->where([
                'fake' => 0,
                'deleted' => 0,
            ])->chunk(1000, function ($users) use (&$storeNames) {
                foreach ($users as $user) {
                    $orders = $this->db->table('order')
                        ->field(['store_id', 'user_id'])
                        ->where('status_time >= ?', '2021-06-16 00:00:00')
                        ->where([
                            'status' => 3,
                            'type' => 2,
                            'user_id' => $user['id']
                        ])->findOne();

                    if (is_null($orders)) {
                        continue;
                    }
                    $storeNames[$orders['store_id']] += $user['balance'];
                }

            });

        print_r($storeNames);
    }
}

(new StatisticsFinal())->caculateUserMoney();
