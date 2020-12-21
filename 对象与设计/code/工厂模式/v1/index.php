<?php

require './ApptEncoder.php';
require './BloggsApptEncoder.php';
require './CommManager.php';
require './MegaApptEncoder.php';

//$comm = new CommManager(CommManager::BLOGGS);
$comm = new CommManager(CommManager::MEGA);
$blog = $comm->getApptEncoder();
var_dump($blog);