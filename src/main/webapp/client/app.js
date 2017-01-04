var constant = require('./utils/constant.js')
var utils = require('./utils/util.js')

App({
  onLaunch: function () {
    // utils.setScreanSize();

    //调用API从本地缓存中获取数据
    var logs = wx.getStorageSync('logs') || [];
    logs.unshift(Date.now());
    wx.setStorageSync('logs', logs);

  },
  
  globalData: {
    
  }
})