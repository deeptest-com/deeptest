var constant = require('../../utils/constant.js')
var about = require('../../services/about.js')

var app = getApp()
Page({
  data: {
    
  },

  onLoad: function () {
    var that = this;
    console.log('onLoad');
  },
  onReady: function () {
    var that = this;
    console.log('onReady');
    wx.setNavigationBarTitle({
      title: constant.event.title
    });
  },
  onShow: function () {
    var that = this;
    console.log('onShow');

  }
})