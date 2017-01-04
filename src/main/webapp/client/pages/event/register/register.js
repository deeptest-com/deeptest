var constant = require('../../../utils/constant.js')
var util = require('../../../utils/util.js')
var register = require('../../../services/register.js')

var app = getApp()
Page({
  data: {
    
  },
  init: function () {
    var that = this
    var width = 750;
    var colHeight = width / 3;

    that.setData({
      colHeight: colHeight
    });
  },

  onLoad: function () {
    var that = this
    that.init();

    register.getInfo(function(json) {
      that.setData({
        info: json.data
      });
    });
  },

  viewBizcard: function () {
    var that = this;
    console.log('viewBizcard');

    wx.navigateTo({
      url: '../bizcard/view'
    })
  },
  editBizcard: function () {
    var that = this;
    console.log('editBizcard');

    wx.navigateTo({
      url: '../bizcard/edit'
    })
  }

})