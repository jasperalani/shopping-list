<?php

namespace ShoppingList;

class Response {
    function response ($message, $response) {
        $message = json_encode($message);
        $response->getBody()->write($message);
        return $response;
    }
}