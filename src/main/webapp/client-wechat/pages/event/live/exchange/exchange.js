var constant = require('../../../../utils/constant.js')
var util = require('../../../../utils/util.js')
var exchange = require('../../../../services/exchange.js')

var app = getApp()
Page({
  data: {
    searching: false
  },
  init: function () {
    var that = this;
    var width = 750;
    var colHeight = width / 4;
    var imgHeight = colHeight * 0.8 * 0.8;

    that.setData({
        colHeight: colHeight,
        imgHeight: imgHeight
    });
  },

  onLoad: function () {
    var that = this
    that.init();

    exchange.getData(function(json) {
      that.setData({
        info: json.data
      });
    });
  },
  startSearch: function () {
    var that = this;

    that.setData({
      searching: !that.data.searching
    });
  }
})