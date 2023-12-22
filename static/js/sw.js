importScripts('/static/js/go_http_sw.js');

registerWasmHTTPListener('/static/wasm/server.wasm', {
    base: '/api',
});

// Skip installed stage and jump to activating stage
addEventListener('install', (event) => {
    console.log('installing service worker');
    event.waitUntil(skipWaiting())
})

// Start controlling clients as soon as the SW is activated
addEventListener('activate', event => {
    console.log('activating service worker');
    event.waitUntil(clients.claim());
})
