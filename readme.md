# Gossip

Gossip is websocket based Messaging Service written in Golang, Gossip uses advantage of concurrency, also it is build around room based messaging architecture.

- Client Limit per Room: 256 (configurable)
- Maximum Message Size: 1024 (configurable)


# Start
```shell
  $ cp env.example .env
  $ docker-compose up --build
```

# Example App (Included)

![Preview Screenshot 1](https://github.com/dkvilo/gossip/blob/master/public/images/chat_s_1.png)
![Preview Screenshot 2](https://github.com/dkvilo/gossip/blob/master/public/images/chat_s_2.png)

# Web Client Example 
```js

// HMAC Token - Available in terminal
const accessToken = "..."

// Room Should be provided by user
const room = "random"

/*
* Make sure that http://localhost:3000 is listed in your .env file as a ALLOWED_ORIGINS
* ALLOWED_ORIGINS=http://localhost:3000,http://localhost:4000 
*/
const ws = new WebSocket(`http://localhost:3000/ws?room=${room}&accessToken=$${accessToken}`);

ws.addEventListener("open", () => {
  console.log("New Client Connected")
  ws.Send("Hello, World")
})

ws.addEventListener("message", (package) => {
  let messages = package.data.split("\n");
  console.log("Messages", messages)
})

```

# Todo
 - [ ] Read/Write messages from redis
 - [ ] Dump messages from redis to MongoDB after some time
 - [ ] Add Message (Package) validation
 - [ ] Generate Room/Namespace from server
 - [ ] Create Better Example app using React


**Author:** David Kviloria


