<?php

//格式管理者
abstract class CommManager
{
    //bloggs格式
    const BLOGGS = 1;

    //mega格式
    const MEGA = 2;

    static protected $mode;

    //获取到编码对象
	static public function getApptEncoder($mode)
    {
        self::$mode = $mode;

        switch (self::$mode) {
            case (self::BLOGGS):
                return new BloggsCommManager();
                break;
            default:
                return new MegaCommManager();
                break;
        }
    }

    //CommManager只是约定getHeaderText方法,具体实现在继承的子类
    abstract public function getHeaderText();
}