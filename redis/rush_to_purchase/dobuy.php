<?php

//购买页面
$id = intval(@$_GET['id']);

if ($id <= 0) {

	exit('非法ID');
}

$productData = getProduct($id);


//检查商品是否存在
if (empty($productData)) {
	exit('没有该商品，或者商品已下架');
}

$redis = new Redis();
$redis->connect('localhost', 6379);

$lockValue = uniqid();

//购买之前，先上锁，保证同时只能一个人进行
if (!getLock($redis, $id, $lockValue, 100)) {
	exit('购买中');
}

//检查库存是否足够
if ($productData['store'] <= 0) {

	exit('商品库存不足，无法购买');
}

//进行购买，减少库存
if (buyProduct($id)) {
	exit('购买成功');
}

exit('购买失败,库存不足');


function getProduct(int $id):array
{
	$pdo = new PDO('mysql:host=localhost;dbname=xdl;default=utf8', 'xdl', 'xdl123456');

	$sqlTpl = 'select * from product where id = ?';

	$stmt = $pdo->prepare($sqlTpl);
	$stmt->execute([
		$id,
	]);

	$data = $stmt->fetch(PDO::FETCH_ASSOC);

	if (empty($data)) {
		return [];
	}
	return $data;
	
}

//购买商品
function buyProduct(int $id)
{
	$pdo = new PDO('mysql:host=localhost;dbname=xdl;default=utf8', 'xdl', 'xdl123456');

	$sql = 'update product set store=store-1 where id = ?';

	$stmt = $pdo->prepare($sql);

	$stmt->execute([
		$id,
	]);

	if ($stmt->rowCount()) {

		$redis = new Redis();
		$redis->connect('localhost', 6379);
		//加入订单,模拟的
		$id = $redis->incr('order:data');

		//这里需要释放锁,否则就算库存足够，其他人在锁定时间内无法购买
		releaseLock($redis, $id, $lockValue);

		return true;
	}

	return false;
}

//加锁
function getLock($redis, $lockName, $lockValue, $expire)
{  
    //px 为毫秒
    return $redis->set('lock:'.$lockName, $lockValue, array('nx', 'px' => $expire));
}

//释放锁
function releaseLock($redis, $lockName, $lockValue)
{
	$lua = "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end";

	$result = $redis->eval($lua, 1, $lockName, $lockValue);

	return $result;
}