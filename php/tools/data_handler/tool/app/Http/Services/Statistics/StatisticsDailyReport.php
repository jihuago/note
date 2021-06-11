<?php


namespace App\Http\Services\Statistics;

use App\Models\Orders;

/**
 * 每日数据统计
 * Class StatisticsDailyReport
 * @package App\Http\Services\Statistics
 */
class StatisticsDailyReport
{

    protected static $startDate;
    protected static $endDate;

    public static function run(string $startDate, string $endDate)
    {
        self::$startDate = $startDate;
        self::$endDate = $endDate;

        self::storeData();
    }

    /*

    门店总营业额

     */
    protected static function storeData()
    {
        $data =  Orders::whereBetween('create_time', [self::$startDate, self::$endDate])
            ->where('status', 3);
//            ->
//        dd($data);
    }
}
