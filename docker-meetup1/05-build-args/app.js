'use strict';

const http = require('http')
    , os = require('os')
    , ip = require('ip')
    , _ = require('lodash');

const PORT=8000;

function handleRequest(request, response){

  let host = os.hostname();
  
  let resObj = JSON.stringify({
    hostName: host, 
    ip:_.split(ip.address(),'.'),
    commit: process.env.COMMIT
  });
  response.writeHead(200, {'Content-Type': 'application/json'});
  response.end(resObj);

}
 
let server = http.createServer(handleRequest);

server.listen(PORT, function(){
  console.log('Server listening on: http://localhost:%s', PORT);
});