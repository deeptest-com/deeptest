var constant = require('../utils/constant.js')
var post = require('./post.js')

function getInfo(callback) {
    post.post('register/getInfo', {
        eventId: constant.eventId
    }, callback);
}
function register(callback) {
    post.post('register/register', {
        eventId: constant.eventId
    }, callback);
}
module.exports = {
  getInfo: getInfo,
  register: register
}