var constant = require('../../../../utils/constant.js')
var util = require('../../../../utils/util.js')
var feedback = require('../../../../services/feedback.js')

var app = getApp()
Page({
  data: {
    types: ['建议', '投诉'],
    index: 0,
    feedbackType: '',
    content: ''
  },
  init: function () {
    var that = this;

  },

  onLoad: function () {
    var that = this
    that.init();

    feedback.list(function(json) {
      that.setData({
        list: json.list
      });
    });
  },

   bindPickerChange: function(e) {
      console.log('picker发送选择改变，携带值为', e.detail.value)
      this.setData({
        index: e.detail.value
      })
    }
})