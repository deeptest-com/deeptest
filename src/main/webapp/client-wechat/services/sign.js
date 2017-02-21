var constant = require('../utils/constant.js')
var post = require('./post.js')
var upload = require('./upload.js')

function sign(callback) {
    wx.chooseImage({
        count: 1,
        sourceType: ['camera'],
        success: function (res) {
            var tempFilePaths = res.tempFilePaths
            console.log(tempFilePaths);

            var filePath = tempFilePaths[0];
            var formData = {
                eventId: constant.eventId,
                extName: 'png'
            };
            upload.sign(filePath, formData, function(res) {
                console.log(res);
                // post.post('chat/getData', {
                //     eventId: constant.eventId
                // }, callback);
            });
        }
    });

}
module.exports = {
  sign: sign
}