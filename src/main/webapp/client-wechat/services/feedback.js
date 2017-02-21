var constant = require('../utils/constant.js')
var post = require('./post.js')

function list(callback) {
    post.post('feedback/list', {
        eventId: constant.eventId
    }, callback);
}
function save(callback) {
    post.post('feedback/save', {
        eventId: constant.eventId
    }, callback);
}
module.exports = {
  list: list,
  save: save
}