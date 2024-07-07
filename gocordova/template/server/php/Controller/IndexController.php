<?php

namespace App\Controller;

class IndexController extends \BaseController
{
    function indexAction()
    {
        return 1123;
    }

    function after($ret,$data = null)
    {
        var_dump($ret,$data);
    }
}