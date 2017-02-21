var constant = require('../utils/constant.js')
var post = require('./post.js')

function getUserInfo(cb){
    var that = this
    if(constant.userInfo){
        typeof cb == "function" && cb(constant.userInfo)
    }else{
        wx.login({
        success: function () {
            wx.getUserInfo({
            success: function (res) {
                console.log('111' + res);
                constant.userInfo = res.userInfo
                typeof cb == "function" && cb(constant.userInfo)
            }
            })
        }
        });
    }
}
module.exports = {
  getUserInfo: getUserInfo
}