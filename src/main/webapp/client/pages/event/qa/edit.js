var constant = require('../../../utils/constant.js')
var util = require('../../../utils/util.js')
var qa = require('../../../services/qa.js')

var app = getApp()
Page({
  data: {
    qustion: ''
  },

  onLoad: function () {
    var that = this;
    
  },
  bindKeyInput: function(e) {
    var that = this;

    that.setData({
      qustion: e.detail.value
    })
  },
  save: function(e) {
    var that = this;
    qa.save(that.data.qustion, function(json) {
      wx.navigateTo({
        url: './qa'
      });
    });
    console.log(that.data);
  }
})