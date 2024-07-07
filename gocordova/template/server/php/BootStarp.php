<?php

define('DS',DIRECTORY_SEPARATOR);
define('APP_ROOT',realpath(__DIR__));
//var_dump(APP_ROOT);
class BootStarp
{
    static function Run(){
        self::autoload();
        self::parseUrl();
    }

    /**
     * 解析请求
     * @return void
     */
    private static function parseUrl(){
        $url = $_SERVER['REQUEST_URI'];
        $url = substr($url,1);
        $data = explode('/',$url);
        define('Controller',ucfirst($data[0] ?? 'Index'));
        define('Action',$data[1] ?? 'index');
        $class = 'App\\Controller\\'.Controller.'Controller';
        $method = Action.'Action';
        try {
            $cls = new $class;
            $cls->before();
            try {
                $ret = call_user_func([$cls,$method]);
                $cls->after($ret);
            }catch (\Exception $e){
                $cls->after('',$e);
            }
        }catch (\Exception $e){
            echo json_encode([
                'code' => 500,
                'msg' => '系统异常'
            ]);
            die();
        }
    }

    /**
     * 注册自动加载
     * @return void
     */
    private static function autoload(){
        $r = spl_autoload_register(function ($class){
            if(strpos($class,'App') === 0){
                $class = substr($class,3);
                $path = APP_ROOT.str_replace('\\','/',$class);
                include $path.'.php';
            }
        },true,true);
    }
}
class BaseController{
    function getIP(){

    }

    function isPost()
    {
        return $_SERVER['REQUEST_METHOD'] == 'POST';
    }
    function before(){}

    function after($data,$excption = null)
    {
        if(is_null($excption)){
            echo json_encode([
                'code'=>200,
                'msg' => 'success',
                'data' => $data
            ]);
        }else{
            echo json_encode([
                'code'=>500,
                'msg' => $excption->getMessage(),
                'data' => []
            ]);
        }

    }
}

//权限检查
class AuthCheck{
    static function Check(){}

    private static function apiCheck(){}
}
trait GBaseSingle{
    /**
     * @var static
     */
    private static $_self = null;
    static function ins(){
        if(self::$_self == null){
            self::$_self = new static();
        }
        return self::$_self;
    }
    private function clone(){}

}
class DB{
    use GBaseSingle;

    /**
     * @var PDO
     */
    private $_pdo;
    private function __construct()
    {
        $pdo = new PDO(sprintf('mysql:host=%s;dbname=%s;port=%s;charset=utf8',Env::get('database.host'),Env::get('database.dbname'),Env::get('database.port')),Env::get('database.user'),Env::get('database.user'));
        $pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
        $pdo->exec('set names utf8');
        $this->_pdo = $pdo;
    }
    function begin(callable $func){
        self::$_pdo->beginTransaction();
        try {
            $ret = $func();
            self::$_pdo->commit();
            return $ret;
        }catch (\Exception $e){
            self::$_pdo->rollBack();
            return false;
        }
    }

    function query($sql,$args = []){
        $start = microtime(true);
        Log::sql($sql);
        $st =  self::$_pdo->prepare($sql);
        $st->execute($args);
        if($st->errorCode() != '2012'){
            $end = microtime(true);
//            Log::sql(' %f %s %s %s',$end - $start,$st->errorCode(),$sql,json_encode($args,true));
            return $st;

        }
        $this->connect(true);
        $st =  self::$_pdo->prepare($sql);
        $st->execute($args);
        $end = microtime(true);
//        Log::sql(' %f %s %s %s',$end - $start,$st->errorCode(),$sql,json_encode($args,true));
        return $st;
    }
    function fetch($sql,$args = []){
        return $this->query($sql,$args)->fetch(\PDO::FETCH_ASSOC);
    }

    function fetchAll($sql,$args = [])
    {
        return $this->query($sql,$args)->fetchAll(\PDO::FETCH_ASSOC);
    }
    function delete($sql,$args = []){
        return $this->query($sql,$args)->rowCount();
    }

    function update($sql,$args = [])
    {
        return $this->query($sql,$args)->rowCount();
    }
    function insert($sql,$args = []){
        $s = $this->query($sql,$args);
        switch($s->rowCount()){
            case 1:
                return self::$_pdo->lastInsertId();
            default:
                return $s->rowCount();
        }
    }

}
class Cache{
    use GBaseSingle;
    private $reids;
    private function __construct()
    {
        $this->reids = new Redis();
        if(!$this->reids->connect(Env::get('redis.host','127.0.0.1'),Env::get('redis.port',6379))){
            throw new Exception('配置错误');
        }
        $passwd = Env::get('redis.passwd');
        if($passwd){
            if(!$this->reids->auth($passwd)){
                throw new Exception('配置错误');
            }
        }
    }
    public function __call($name, $arguments)
    {
        return call_user_func_array([$this->reids,$name],$arguments);
    }
}

/**
 * 配置信息
 */
class Env{
    private static $data = null;

    static function get($key,$def = ''){
        if(is_null(self::$data)){
            $path = APP_ROOT.DS.'.env';
            if(file_exists($path)){
                self::$data = parse_ini_file($path);
            }else{
                self::$data = [];
            }
        }
        return self::$data[$key] ?? $def;
    }
}

