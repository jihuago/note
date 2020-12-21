<?php

	// header("Content-Type:text/html;charset=utf-8");
	$pdo = new PDO('mysql:host=localhost;dbname=xdl;default=utf8', 'xdl', 'xdl123456');
	
	$sqlTpl = 'select id,name,store,price from product where id = ?';

	$gid = intval(@$_GET['id']);

	if ($gid <= 0) {
		exit('非法ID');
	}

	$stmt = $pdo->prepare($sqlTpl);
	$stmt->execute([
		$gid,
	]);

	$data = $stmt->fetch(PDO::FETCH_ASSOC);
?>
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>商品详情页</title>
</head>
<body>

	商品名：<?php echo $data['name']?><br>
	价格：<?php echo $data['price']?><br>
	库存：<?php echo $data['store']?><br>
	<a href="./dobuy.php?id=<?php echo $data['id']?>">直接购买</a>
</body>
</html>