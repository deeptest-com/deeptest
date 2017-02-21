var constant = require('../utils/constant.js')
var post = require('./post.js')

function list(callback) {
    post.post('service/list', {
        eventId: constant.eventId
    }, callback);
}
function get(serviceId, callback) {
    post.post('service/get', {
        serviceId: serviceId
    }, callback);
}
module.exports = {
  list: list,
  get: get
}