<?php

namespace ShoppingList;

use mysqli;
use Slim\App;
use Slim\Exception\HttpNotFoundException;
use Slim\Factory\AppFactory;
use Slim\Handlers\Strategies\RequestHandler;
use Slim\Psr7\Factory\ResponseFactory;
use Slim\Psr7\Stream;
use Slim\Routing\RouteCollectorProxy;
use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\Routing\RouteContext;

// Add compression

class Alexandra
{

    public static $db;
    public static $response;

    public function __construct()
    {
        Alexandra::$db       = $this->connectDB();
        Alexandra::$response = new \ShoppingList\Response();

        $app = AppFactory::create();
        $app->add(CorsMiddleware::class);
        $app->addRoutingMiddleware();

        $this->routes($app);
        $this->renderErrors($app);

        $app->run();
    }

    public function routes(App $app)
    {
        $app->options('/image', function (Request $request, Response $response): Response {
            return $response;
        });

        $app->group('/image', function (RouteCollectorProxy $group) {
            // Upload image
            $group->post('', function (Request $request, Response $response) {
                $files = $_FILES;

                if (empty($files)) {
                    return Alexandra::$response->response(
                        ['error' => 'no file supplied'],
                        $response
                    );
                }

                if (empty($files['file'])) {
                    return Alexandra::$response->response(
                        ['error' => 'file supplied is not an image'],
                        $response
                    );
                }

                if (sizeof($files) > 1) {
                    return Alexandra::$response->response(
                        ['error' => 'currently no support for multiple images at a time'],
                        $response
                    );
                }

                $createNewImage = "INSERT INTO images (id) VALUES (null);";
                $result         = Alexandra::$db->query($createNewImage);

                if ( ! $result) {
                    return Alexandra::$response->response(
                        ['error' => 'failed to insert row in database'],
                        $response
                    );
                }

                $maxId  = "SELECT MAX(id) FROM images;";
                $result = Alexandra::$db->query($maxId);

                if ( ! $result) {
                    return Alexandra::$response->response(
                        ['error' => 'failed to insert row in database'],
                        $response
                    );
                }

                $id = $result->fetch_row()[0];

                $file_type     = $files['file']['type'];
                $file_ext      = explode('/', $file_type)[1];
                $file_name     = $id . '.' . $file_ext;
                $file_tmp_name = $files['file']['tmp_name'];

                $saveExtension = "UPDATE images SET ext = '" . $file_ext . "' WHERE id = " . $id . ";";
                $result        = Alexandra::$db->query($saveExtension);

                $moved_file = move_uploaded_file($file_tmp_name, 'images/' . $file_name);

                if ( ! $moved_file) {
                    return Alexandra::$response->response(
                        ['error' => 'failed to move file from temporary location'],
                        $response
                    );
                }

                return Alexandra::$response->response(
                    ['image_id' => intval($id)],
                    $response
                );
            });

            // Retrieve image
            $group->get('/{id}', function (Request $request, Response $response, $args) {
                $getId  = "SELECT * FROM images WHERE id = " . $args['id'];
                $result = Alexandra::$db->query($getId);

                $image = $result->fetch_assoc();

                if ( ! $result) {
                    return Alexandra::$response->response(
                        ['error' => 'failed to query database'],
                        $response
                    );
                }


                $file     = "images/" . $image['id'] . '.' . $image['ext'];
                $openFile = fopen($file, 'rb');
                $stream   = new Stream($openFile);

                return $response->withBody($stream)
                                ->withHeader('Access-Control-Allow-Origin', '*')
                                ->withHeader('Access-Control-Allow-Headers',
                                    'X-Requested-With, Content-Type, Accept, Origin, Authorization')
                                ->withHeader('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE, PATCH, OPTIONS');
            });

            $group->map(['GET', 'POST', 'PUT', 'DELETE', 'PATCH'], '/{routes:.+}', function ($request, $response) {
                throw new HttpNotFoundException($request);
            });
        });
    }

    private function renderErrors(App $app)
    {
        switch ($_ENV) {
            case 'development':
                $error = [true, true, true];
                break;
            default:
                $error = [false, false, false];
                break;
        }
        $errorMiddleware = $app->addErrorMiddleware($error[0], $error[1], $error[2]);
        $errorHandler    = $errorMiddleware->getDefaultErrorHandler();
        $errorHandler->registerErrorRenderer('text/html', ErrorRenderer::class);
    }

    function connectDB()
    {
        $mysqli = new mysqli(
            '127.0.0.1',
            'root',
            'password',
            'alexandra',
            '3306'
        );

        if ($mysqli->connect_errno) {
            echo "Error: Failed to make a MySQL connection, here is why: \n";
            echo "Errno: " . $mysqli->connect_errno . "\n";
            echo "Error: " . $mysqli->connect_error . "\n";
            echo "Please report this error to your webmaster.";
            exit;
        }

        return $mysqli;
    }

}