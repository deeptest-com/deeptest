var constant = require('../../utils/constant.js')
var find = require('../../services/find.js')

var app = getApp()
Page({
  data: {
    test: 'aa'
  },
  onLoad: function () {
    var that = this;
    console.log('onLoad');
  },
  onReady: function () {
    var that = this;
    console.log('onReady');
  },
  onShow: function () {
    var that = this;
    console.log('onShow');

  }
})