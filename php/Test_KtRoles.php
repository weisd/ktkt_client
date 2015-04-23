<?php
/**
 * Created by PhpStorm.
 * User: root
 * Date: 15-4-22
 * Time: 下午5:13
 */

require "include.php";

try{

    $new = new KtRoles;
    $new->title = "haha";
    $new->keywords = "nimei";

    $aft = KtRolesClient::Create($new);
    var_dump($aft);

//    $f = new wherer();
//    $f->where="id=?";
//    $f->args=[10];

    $res = KtRolesClient::Find();
    var_dump($res);

}catch (Exception $e){
    echo "error : ".$e->getMessage();
}
