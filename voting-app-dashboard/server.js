const express = require('express')
const async = require('async')
const path = require("path")
const cookieParser = require('cookie-parser')
const bodyParser = require('body-parser')
const methodOverride = require('method-override')
const app = express()
const server = require('http').Server(app)
const io = require('socket.io')(server);
const grpc = require("grpc");
const protoLoader = require("@grpc/proto-loader");

io.set('transports', ['polling']);

let port = process.env.PORT || 8081;
const REMOTE_SERVER = "localhost:50051";
let pb = grpc.loadPackageDefinition(
  protoLoader.loadSync("../voting-app-pb/vote.proto", {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
  })
).pb;
let client = new pb.VoteWorkerService(
  REMOTE_SERVER,
  grpc.credentials.createInsecure()
);

io.sockets.on('connection', function (socket) {
  socket.emit('message', { text : 'Welcome!' });
  socket.on('subscribe', function (data) {
    socket.join(data.channel);
  });
});

getVotes()
function getVotes() {
  let res = []
  let getVotesResults = client.GetVotesResults({ query: 'username' })
  getVotesResults.on("data", (msg) => res.push(msg))
  getVotesResults.on('error', () => console.log("error"));
  getVotesResults.on('end', () => {
    io.sockets.emit("vote-result", collectVotesFromResult(res));
  });

  setTimeout(() => getVotes(client) , 1000);
}

function collectVotesFromResult(result) {
    var votes = {a: 0, b: 0};
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
  var port = server.address().port;
  console.log('App running on port ' + port);
});
