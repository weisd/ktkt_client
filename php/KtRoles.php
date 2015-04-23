<?php
/**
 * Created by PhpStorm.
 * User: root
 * Date: 15-4-22
 * Time: 下午4:57
 */
//
class KtRoles {
    public $id;
    public $title;
    public $keywords;
    public $deletedAt;
    public $createdAt;
    public $updatedAt;

}

class KtRolesClient extends KtClient {

    public static  function Create(KtRoles $info){
        $client = KtClient::getClient();

        return $client->KtRolesCreate($info);
    }


    public static function Find(Wherer $where = null, Limiter  $limit = null, Sorter $sort = null){
        $client = KtClient::getClient();
        if($limit == null){
            $limit = new  Limiter();
        }

        if($where == null){
            $where = new  Wherer();
        }

        if($sort == null){
            $sort = new  Sorter();
        }

        return $client->KtRolesGetList($where, $limit, $sort);
    }
}