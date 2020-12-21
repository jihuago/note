<?php

//格式管理者
class CommManager
{
    //bloggs格式
    const BLOGGS = 1;
    //mega格式
    const MEGA = 2;

    public function __construct($mode)
    {
        $this->mode = $mode;
    }

    //获取到编码对象
	public function getApptEncoder()
	{
	    switch ($this->mode) {
            case (self::BLOGGS):
                return new BloggsApptEncoder();
                break;
            default:
                return new MegaApptEncoder();
                break;
        }

	}

	//当需要提供其他功能时，比如提供页眉功能。那么判断的代码又将重复
    public function getHeaderText()
    {
        switch ($this->mode) {
            case (self::BLOGGS) :
                return new BloggsApptEncoder();
                break;
            default:
                return new MegaApptEncoder();
            
        }
    }
}