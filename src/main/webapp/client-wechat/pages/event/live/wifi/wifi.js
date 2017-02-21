var constant = require('../../../../utils/constant.js')
var util = require('../../../../utils/util.js')
var wifi = require('../../../../services/wifi.js')

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

    wifi.getData(function(json) {
      that.setData({
        info: json.data
      });
    });
  }
})