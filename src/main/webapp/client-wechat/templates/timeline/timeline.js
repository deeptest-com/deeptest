var constant = require('../../utils/constant.js')
var mobiu = require('../../service/mobiu.js')

var app = getApp()
Page({
  data: {
    mode: 'bySession',
    list: []
  },

  onLoad: function () {
    var that = this;
    console.log('onLoad');
  },
})