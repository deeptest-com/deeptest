var serviceUrl = 'http://localhost:8080/events/';
var websocketPath = 'ws';
var eventId = -1;
var token = '123456';

module.exports = {
  serviceUrl: serviceUrl,
  apiUrl: serviceUrl + 'api/client/v1/',
  websocketPath: websocketPath,
  eventId: eventId + '',
  token: token,
  
  userInfo: null,
  event: null,
  W: null,
  H: null
}
