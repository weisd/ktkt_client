<?php
/**
 * Created by PhpStorm.
 * User: root
 * Date: 15-4-22
 * Time: 下午4:39
 */
require './hprose-php/Hprose.php';
$conf = require './conf.php';
$client = new HproseHttpClient($conf['server']);
