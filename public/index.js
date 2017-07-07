var callback = function (message) {
  console.log(message)
}

var handler = function (config) {
  var centrifuge = new Centrifuge({
    url: config.url,
    user: config.user,
    timestamp: config.timestamp,
    token: config.token,
    sockJS: SockJS
  })

  config.channels.forEach(channel => {
    centrifuge.subscribe(channel, callback)
  })

  centrifuge.connect()
}

var get = function (url) {
  let request = new XMLHttpRequest()
  request.open('GET', url)

  request.onreadystatechange = () => {
    if (request.readyState === 4) { // DONE
      if (request.status === 200) { // OK
        handler(JSON.parse(request.responseText))
      }
    }
  }
  request.send()
}

var url = 'http://localhost:8080/websocket/user'
get(url)
