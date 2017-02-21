var constant = require('../../../../utils/constant.js')
var util = require('../../../../utils/util.js')
var bizcard = require('../../../../services/bizcard.js')

var app = getApp()
Page({
  data: {
    
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

    bizcard.list(function(json) {
      that.setData({
        list: json.list
      });
    });
  }
})