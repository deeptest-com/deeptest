var constant = require('../utils/constant.js')
var post = require('./post.js')

function getData(callback) {
    post.post('find/index', {
        eventId: constant.EVNET_ID
    }, callback);
}
module.exports = {
  getData: getData
}
