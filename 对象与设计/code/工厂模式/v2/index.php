<?php

require './ApptEncoder.php';
require './BloggsApptEncoder.php';
require './CommManager.php';
require './MegaApptEncoder.php';
require './MegaCommManager.php';
require './BloggsCommManager.php';

//v2
//$comm = CommManager::getApptEncoder(CommManager::MEGA);
$comm = CommManager::getApptEncoder(CommManager::BLOGGS);
$headerText = $comm->getHeaderText();
var_dump($comm);
echo '<br/>',$headerText;
