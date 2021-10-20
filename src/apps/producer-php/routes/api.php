<?php

use App\Models\Film;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use Junges\Kafka\Facades\Kafka;
use Junges\Kafka\Message\Message;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::post('/', function (Request $request) {
    try {
        $film = Film::create($request->all());
    }
    catch (\Exception) {
        abort(500, 'Falha ao salvar o filme');
    }

    $message = new Message(
        'topico-exemplo',
        body: $request->all(),
        key: $film->id
    );

    $sent = Kafka::publishOn('broker:29092', 'topico-exemplo')
        ->withMessage($message)
        ->send();
    
    if (! $sent) {
        abort(500, 'Falha ao publicar o evento');
    }

    return response()->noContent();
});
