var constant = require('../utils/constant.js')
var post = require('./post.js')

function get(callback) {
    post.post('event/get', {
        eventId: constant.eventId
    }, callback);
}
module.exports = {
        get: get
}