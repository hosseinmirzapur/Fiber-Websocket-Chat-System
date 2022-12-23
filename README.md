
# Minimal chat application for [Fiber](https://github.com/gofiber/fiber) with [Fiber WebSocket](https://github.com/gofiber/websocket)

Uses a client hub, 
somewhat similar to the 
[gorilla chat example](https://github.com/gorilla/websocket/tree/master/examples/chat) 
and 
[fasthttp chat example](https://github.com/fasthttp/websocket/tree/master/_examples/chat).

## Usage

`go run .` and use postman to generate key and put it in browser's ``localstorage`` to connect
<br>

or e.g.

`go run . -addr=3000` and provide a front-end to the chat-application using
[Websocket library of Javascript](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API) 
to connect to the [websocket on localhost](ws://localhost:3000) endpoint.

