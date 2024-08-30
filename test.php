<?php

/**
 * Этот PHP-скрипт предназначен для тестирования различных опций API message_service.
 * Он позволяет отправлять HTTP-запросы к разным конечным точкам API и проверять ответы.
 * 
 * Включает тесты для:
 * - Регистрация и вход пользователя
 * - Отправка и получение сообщений
 * - Работа с чатами
 * - Оценка техподдержки
 * - Работа с вложениями
 */

// Базовый URL вашего API
$baseUrl = 'http://localhost:3336';

// Функция для отправки POST-запроса
function sendPostRequest($url, $data) {
    $options = [
        'http' => [
            'header'  => "Content-type: application/json\r\n",
            'method'  => 'POST',
            'content' => json_encode($data),
        ],
    ];
    $context  = stream_context_create($options);
    return file_get_contents($url, false, $context);
}

// Функция для отправки GET-запроса
function sendGetRequest($url) {
    $options = [
        'http' => [
            'method'  => 'GET',
    ]];
    $context  = stream_context_create($options);
    return file_get_contents($url, false, $context);
}

// Функция для регистрации пользователя
function testUserRegistration($username, $password) {
    global $baseUrl;
    $url = $baseUrl . '/register';
    $data = [
        'username' => $username,
        'password' => $password,
    ];
    return sendPostRequest($url, $data);
}

// Функция для входа пользователя
function testUserLogin($username, $password) {
    global $baseUrl;
    $url = $baseUrl . '/login';
    $data = [
        'username' => $username,
        'password' => $password,
    ];
    $response = sendPostRequest($url, $data);
    $responseData = json_decode($response, true);
    return $responseData['token'] ?? null;
}

// Функция для отправки сообщения
function testSendMessage($token, $chatId, $message) {
    global $baseUrl;
    $url = $baseUrl . '/message';
    $data = [
        'chat_id' => $chatId,
        'message' => $message,
    ];
    $options = [
        'http' => [
            'header'  => "Content-type: application/json\r\n" .
                         "Authorization: Bearer $token\r\n",
            'method'  => 'POST',
            'content' => json_encode($data),
        ],
    ];
    $context  = stream_context_create($options);
    return file_get_contents($url, false, $context);
}

// Функция для получения сообщений
function testGetMessages($token, $chatId) {
    global $baseUrl;
    $url = $baseUrl . "/messages?chat_id=$chatId";
    $options = [
        'http' => [
            'header'  => "Authorization: Bearer $token\r\n",
            'method'  => 'GET',
        ],
    ];
    $context  = stream_context_create($options);
    return file_get_contents($url, false, $context);
}

// Функция для создания чата
function testCreateChat($token, $chatName, $users) {
    global $baseUrl;
    $url = $baseUrl . '/chat';
    $data = [
        'name' => $chatName,
        'users' => $users,
    ];
    $options = [
        'http' => [
            'header'  => "Content-type: application/json\r\n" .
                         "Authorization: Bearer $token\r\n",
            'method'  => 'POST',
            'content' => json_encode($data),
        ],
    ];
    $context  = stream_context_create($options);
    return file_get_contents($url, false, $context);
}

// Функция для оценки техподдержки
function testAddSupportRating($token, $chatId, $rating, $comment) {
    global $baseUrl;
    $url = $baseUrl . '/support/rating';
    $data = [
        'chat_id' => $chatId,
        'rating' => $rating,
        'comment' => $comment,
    ];
    $options = [
        'http' => [
            'header'  => "Content-type: application/json\r\n" .
                         "Authorization: Bearer $token\r\n",
            'method'  => 'POST',
            'content' => json_encode($data),
        ],
    ];
    $context  = stream_context_create($options);
    return file_get_contents($url, false, $context);
}

// Функция для загрузки вложения
function testUploadAttachment($token, $messageId, $filePath) {
    global $baseUrl;
    $url = $baseUrl . '/attachment';
    $data = [
        'message_id' => $messageId,
        'file' => new CURLFile($filePath),
    ];
    $options = [
        'http' => [
            'header'  => "Authorization: Bearer $token\r\n",
            'method'  => 'POST',
        ],
    ];
    $context  = stream_context_create($options);
    return file_get_contents($url, false, $context);
}

// Функция для получения вложения
function testDownloadAttachment($token, $attachmentId) {
    global $baseUrl;
    $url = $baseUrl . "/attachment/$attachmentId";
    $options = [
        'http' => [
            'header'  => "Authorization: Bearer $token\r\n",
            'method'  => 'GET',
        ],
    ];
    $context  = stream_context_create($options);
    return file_get_contents($url, false, $context);
}

// Тестовые данные
$username = "testuser";
$password = "testpassword";

// 1. Регистрация пользователя
echo "Регистрация пользователя:\n";
$response = testUserRegistration($username, $password);
echo $response . "\n";

// 2. Вход пользователя
echo "Вход пользователя:\n";
$token = testUserLogin($username, $password);
echo "Token: " . $token . "\n";

if ($token) {
    // 3. Создание чата
    echo "Создание чата:\n";
    $chatName = "Test Chat";
    $users = [1, 2]; // Замените на реальные ID пользователей
    $response = testCreateChat($token, $chatName, $users);
    echo $response . "\n";

    // 4. Отправка сообщения
    echo "Отправка сообщения:\n";
    $chatId = 1; // Замените на реальный ID чата
    $message = "Привет, это тестовое сообщение!";
    $response = testSendMessage($token, $chatId, $message);
    echo $response . "\n";

    // 5. Получение сообщений
    echo "Получение сообщений:\n";
    $response = testGetMessages($token, $chatId);
    echo $response . "\n";

    // 6. Добавление оценки техподдержке
    echo "Добавление оценки техподдержке:\n";
    $rating = 5;
    $comment = "Отличная работа!";
    $response = testAddSupportRating($token, $chatId, $rating, $comment);
    echo $response . "\n";

    // 7. Загрузка вложения
    echo "Загрузка вложения:\n";
    $filePath = "/path/to/your/file.txt"; // Замените на реальный путь к файлу
    $response = testUploadAttachment($token, $chatId, $filePath);
    echo $response . "\n";

    // 8. Загрузка вложения
    echo "Загрузка вложения:\n";
    $attachmentId = 1; // Замените на реальный ID вложения
    $response = testDownloadAttachment($token, $attachmentId);
    echo $response . "\n";
}

?>
