var constant = require('../utils/constant.js')
var post = require('./post.js')

function getData(callback) {
    post.post('chat/getData', {
        eventId: constant.eventId
    }, callback);
}
module.exports = {
  getData: getData
}