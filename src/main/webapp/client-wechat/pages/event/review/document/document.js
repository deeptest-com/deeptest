var constant = require('../../../../utils/constant.js')
var util = require('../../../../utils/util.js')
var document = require('../../../../services/document.js')

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

    document.list(function(json) {
      that.setData({
        list: json.list
      });
    });
  }
})