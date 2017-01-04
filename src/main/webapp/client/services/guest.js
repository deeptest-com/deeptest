var constant = require('../utils/constant.js')
var post = require('./post.js')

function list(callback) {
    post.post('guest/list', {
        eventId: constant.eventId
    }, callback);
}
function get(guestId, callback) {
    post.post('guest/get', {
        guestId: guestId
    }, callback);
}

module.exports = {
  list: list,
  get: get
}