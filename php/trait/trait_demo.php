<?php

/*

trait的一些例子
 */

trait User
{
	public function getUser()
	{
		return $this->name;
	}
}

class Login
{
	use User;

	public $name = 'jack';

	public function showLogin()
	{
		//调用trait User的getUser方法
		return $this->getUser();
	}
}

$login = new Login;
print $login->showLogin();