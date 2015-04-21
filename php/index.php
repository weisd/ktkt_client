<?php
require './hprose-php/Hprose.php';
$client = new HproseHttpClient('http://127.0.0.1:3000');

class StrategyStock {
    public $id;
    public $strategyId;
    public $stkname;
    public $stkcode;
    public $codeNo;
    public $fmlprice;
    public $fmltime;
    public $fmlname;
    public $date;
    public $up;
    public $down;
    public $createdAt;
    public $updatedAt;
}

class wherer {
    public $where;
    public $args;
}

class limiter {
    public $start;
    public $offset;
}

class sorter {
    public $orderby;
}




// $res = $client->NewStrategyStock($s);
// var_dump($res);

if (!isset($argv[1])) {
    die("no method given");
}

$method = $argv[1];

$classMap = ["strategygetlistcache", "strategygetlist", "strategystockgetlist", "strategystockgetlistcache", "strategystockcreate", "strategyusergetlistcache"];

$method = strtolower($method);


if ($method == "strategystockcreate") {
    $s = new StrategyStock;
    $s->stkname = "stkname";
    $s->strategyId = 1;
    $s->stkcode = "stkcode";
    $s->codeNo = "codeNo";
    $s->fmlprice = "fmlprice";
    $s->fmlname = "fmlname";
    $s->fmltime = "fmltime";
    $s->date = "date";
    $s->up = 8;
    $s->down = 89;

    $res = $client->$method($s);
    echo "sss ssresult: \n";
    var_dump($res);
    exit(1);

}else if(false !== strpos($method, "getone")){
    $wherer = new wherer;
    $wherer->where = "id = ?";
    $wherer->args = [2];
    $res = $client->$method($wherer);
    echo "sss ssresult: \n";
    var_dump($res);
    exit(1);

}else if (in_array($method, $classMap)) {
    echo "sss strtolower: \n";

    $wherer = new wherer;
    $wherer->where = "id = ?";
    $wherer->args = [2];
    $res = $client->$method($wherer, new limiter, new sorter);
    echo "sss ssresult: \n";
    var_dump($res);
    exit(1);
}



$arg = array_slice($argv, 2, $argc - 2);

echo "do method : {$method}  with args : ".json_encode($arg)."\n";

$res = call_user_func_array(array($client, $method), $arg);

// $res = $client->$method();
echo "result: \n";
var_dump($res);