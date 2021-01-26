<?php
/*

    统计每个发型师的服务次数，购买会员数，回头人数，回头率 = 回头人数 / 服务人数

        1. 每个发型师的服务人数、购买会员数 sql
          SELECT COUNT(DISTINCT(o.`user_id`)),COUNT(o.id),
               o.`employee_id`,
               o.`store_name`,
               o.`type`
          from `gp_employee` as e
          INNER JOIN `gp_order` as o on e.id= o.`employee_id`
          where o.`status`= 3 AND e.`status`  = 1 and e.`job_status`  != 3
           and e.`role`  in(1, 2) and e.`id`  not IN (5675,5676,5677,5678,5679,5680,5681,5682)
          GROUP BY o.`employee_id`, o.`type`

        2. 得到回头人数

            SELECT COUNT(o.`user_id`) as 服务次数,
               o.`user_id`,
               o.`employee_id`,
               e.`name` as 发型师,
               o.`store_name` as 店铺,
               if(o.`type`= 1, '项目消费', '购买会员') as 消费类型
          from `gp_employee` as e
          INNER JOIN `gp_order` as o on e.id= o.`employee_id`
         where o.`status`= 3
           AND e.`status`= 1
           and e.`job_status`!= 3
           and e.`role` in(1, 2)
           and e.`id` not IN(5675, 5676, 5677, 5678, 5679, 5680, 5681, 5682)
         GROUP BY o.`employee_id`,
                 o.`type`,
                 o.`user_id`
        HAVING COUNT(o.`user_id`)> 1;
*/
