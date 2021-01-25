<?php
/*

    统计每个发型师的服务人数，购买会员数，回头人数，回头率 = 回头人数 / 客单量

        1. 每个发型师的服务人数sql
            explain SELECT COUNT(DISTINCT(o.`user_id`)),
               o.`employee_id`,
               o.`store_name`
          from `gp_employee` as e
          INNER JOIN `gp_order` as o on e.id= o.`employee_id`
          where o.`status`= 3
          GROUP BY o.`employee_id`


*/
