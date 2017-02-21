var constant = require('../utils/constant.js')
var post = require('./post.js')

function list(callback) {
    post.post('bizcard/list', {
        eventId: constant.eventId
    }, callback);
}
module.exports = {
  list: list
}