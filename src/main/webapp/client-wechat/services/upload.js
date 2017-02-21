var constant = require('../utils/constant.js')
var post = require('./post.js')

function sign(filePath, formData, callback) {
    var url = constant.apiUrl + 'sign?Authorization=' + constant.token;
    console.log(url);

    wx.uploadFile({
        url: url,
        filePath: filePath,
        name: 'file',
        formData: formData,
        success: function(res) {
            console.log('上传成功', res);
            callback(res);
        },
        fail: function(res) {
            console.log('上传失败', res);
        },
        complete: function(res) {
            console.log('上传完成', res);
        }
    });
}

module.exports = {
        sign: sign
}
