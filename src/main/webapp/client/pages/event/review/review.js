var constant = require('../../../utils/constant.js')
var util = require('../../../utils/util.js')
var review = require('../../../services/review.js')

var app = getApp()
Page({
  data: {
    
  },
  init: function () {
    var that = this;
    var width = 750;
    var colHeight = width / 3;

    that.setData({
      colHeight: colHeight
    });
  },

  onLoad: function () {
    var that = this
    that.init();

    review.getData(function(json) {
      that.setData({
        data: json.data
      });
    });
  },

  gotoPage: function(event) {
    var page = event.currentTarget.dataset.page;
    console.log(page);
    wx.navigateTo({
      url: './' + page + '/' + page
    })
  }
})