const express = require('express')
const path = require("path")
const cookieParser = require('cookie-parser')
const bodyParser = require('body-parser')
const methodOverride = require('method-override')
const axios = require('axios');
// const buffer = require('buffer');
const app = express()
const server = require('http').Server(app)
const io = require('socket.io')(server);

io.set('transports', ['polling']);

let port = process.env.PORT || 8081;
const REST_REMOTE_SERVER = "http://worker:50052";
// const REST_REMOTE_SERVER = "http://127.0.0.1:50052";

io.sockets.on('connection', function (socket) {
  socket.emit('message', { text : 'Welcome!' });
  socket.on('subscribe', function (data) {
    socket.join(data.channel);
  });
});
getVotes()
function getVotes() {
  axios({
    method: 'post',
    url: REST_REMOTE_SERVER + '/v1/results',
    data: {"query": "result"},
  }).then(function (response) {
    // TODO: different version has different type of response
    let result = []
    if (typeof response.data === 'string' || response.data instanceof String) {
      let eles = (response.data) ? response.data.split('\n') : ''
      for (let idx in eles) {
        if (eles[idx] != '') {
          result.push(JSON.parse(eles[idx]).result)
        }
      }
    } else {
      result.push(response.data.result)
    }

    io.sockets.emit("vote-result", collectVotesFromResult(result));
  }).catch(function (error) {
    console.log(error);
  });

  setTimeout(() => getVotes() , 1000);
}

function collectVotesFromResult(result) {
  let votes = {a: 0, b: 0};
  for (idx in result) {
    if (result[idx].vote === 'cat') votes.a = result[idx].count
    if (result[idx].vote === 'dog') votes.b = result[idx].count
  }
  return votes
}

app.use(cookieParser());
app.use(bodyParser.urlencoded({
  extended: true
}));
app.use(bodyParser.json());
app.use(methodOverride('X-HTTP-Method-Override'));
app.use(function(req, res, next) {
  res.header("Access-Control-Allow-Origin", "*");
  res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
  res.header("Access-Control-Allow-Methods", "PUT, GET, POST, DELETE, OPTIONS");
  next();
});

app.use(express.static(__dirname + '/views'));

app.get('/', function (req, res) {
  res.sendFile(path.resolve(__dirname + '/views/index.html'));
});

server.listen(port, function () {
  let port = server.address().port;
  console.log('App running on port ' + port);
});
