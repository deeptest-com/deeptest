var constant = require('../utils/constant.js')
var post = require('./post.js')

function getData(callback) {
    post.post('mobiu/index', {
        eventId: constant.EVNET_ID
    }, callback);
}
module.exports = {
  getData: getData
}