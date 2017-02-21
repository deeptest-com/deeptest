var constant = require('../../../../utils/constant.js')
var util = require('../../../../utils/util.js')
var chat = require('../../../../services/chat.js')

var app = getApp()
Page({
  data: {
    
  },
  init: function () {
    var that = this;

  },

  onLoad: function () {
    var that = this
    that.init();

    chat.getData(function(json) {
      that.setData({
        info: json.data
      });
    });
  },

  comments: function() {
    wx.navigateTo({
      url: './comments'
    })
  }
})