<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use App\Http\Services\Statistics\StatisticsDailyReport as StatisticsDailyReportTool;

/*
 * 命令执行： php artisan laravel:StatisticDaily 开始日期(2021-06-01) 结束日期(2021-06-09)
 */
class StatisticsDailyReport extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'laravel:StatisticDaily {startDate}{endDate}';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = '每日报表统计，俏手艺门店数据说明';

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();
    }

    /**
     * Execute the console command.
     *
     * @return int
     */
    public function handle()
    {
        StatisticsDailyReportTool::run($this->argument("startDate"), $this->argument("endDate"));
    }
}
