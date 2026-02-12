<?php

$apiKey = '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';

$ch = curl_init('https://integrate.api.nvidia.com/v1/chat/completions');

$payload = [
    'model' => 'moonshotai/kimi-k2.5',
    'messages' => [
        ['role' => 'user', 'content' => 'Reply only with: okkay working']
    ],
    'max_tokens' => 10,
];

curl_setopt_array($ch, [
    CURLOPT_POST => true,
    CURLOPT_RETURNTRANSFER => true,
    CURLOPT_HTTPHEADER => [
        "Authorization: Bearer {$apiKey}",
        'Content-Type: application/json',
    ],
    CURLOPT_POSTFIELDS => json_encode($payload),
]);

$response = curl_exec($ch);

if ($response === false) {
    die(curl_error($ch));
}

curl_close($ch);

echo $response;
