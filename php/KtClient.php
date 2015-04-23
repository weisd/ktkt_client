<?php
/**
 * Created by PhpStorm.
 * User: root
 * Date: 15-4-22
 * Time: 下午4:58
 */
class KtClient {
    public static $_client;

    public static function getClient($server = "http://127.0.0.1:3030"){
        if(self::$_client === null){
            require './hprose-php/Hprose.php';
            self::$_client = new  HproseHttpClient($server);
        }

        return self::$_client;
    }
}

class Wherer {
    public $where;
    public $args;
}

class Limiter {
    public $start;
    public $offset;
}

class Sorter {
    public $orderby;
}
