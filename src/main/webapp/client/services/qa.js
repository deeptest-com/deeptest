var constant = require('../utils/constant.js')
var post = require('./post.js')

function list(callback) {
    post.post('qa/list', {
        eventId: constant.eventId
    }, callback);
}
function save(callback) {
    post.post('qa/save', {
        eventId: constant.eventId
    }, callback);
}
module.exports = {
  list: list,
  save: save
}