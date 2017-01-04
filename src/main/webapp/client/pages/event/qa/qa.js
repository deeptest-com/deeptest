var constant = require('../../../utils/constant.js')
var util = require('../../../utils/util.js')
var qa = require('../../../services/qa.js')

var app = getApp()
Page({
  data: {
    
  },

  onLoad: function () {
    var that = this
    
    qa.list(function(json) {
      that.setData({
          qas: json.qas
      });
    });
  },
  create: function () {
    var that = this
    
    wx.navigateTo({
      url: './edit'
    })
  }
})