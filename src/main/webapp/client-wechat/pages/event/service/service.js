var constant = require('../../../utils/constant.js')
var util = require('../../../utils/util.js')
var service = require('../../../services/service.js')

var app = getApp()
Page({
  data: {
    
  },

  onLoad: function () {
    var that = this
    
    console.log('onLoad');
    that.init();

    service.list(function(json) {
      that.setData({
        services: json.services
      });
    });
  },

  init: function () {
    var that = this
    var width = 750;
    var colHeight = width / 3;

    that.setData({
      colHeight: colHeight
    });
  },

  gotoDetail: function (event) {
    var typeId = event.currentTarget.dataset.typeId;
    wx.navigateTo({
      url: './detail?typeId=' + typeId
    })
  }

})