var app = require('express')();
var http = require('http').createServer(app);
var io = require('socket.io')(http);

app.get('/', function(req, res){
  res.sendFile(__dirname + '/index.html');
});


io.on('connection', function(socket){
  //console.log('a user connected');
  socket.on('chat message', function(msg){
    socket.broadcast.emit('User joined');
    io.emit('chat message', {message: msg.message, username:msg.username });
  });

  socket.on('new user', function(msg){
      console.log("executing");
    socket.broadcast.emit('new user','User joined');
  })

  socket.on('typing', function(data){
      socket.broadcast.emit('typing', { username: data.username});
  })

});

http.listen(3000, function(){
  console.log('listening on *:3000');
});