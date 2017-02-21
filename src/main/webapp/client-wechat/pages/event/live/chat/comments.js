var constant = require('../../../../utils/constant.js')
var util = require('../../../../utils/util.js')
var comments = require('../../../../services/comments.js')

var app = getApp()
Page({
  data: {
    content: '',
    modalShow: false
  },
  init: function () {
    var that = this;

  },

  onLoad: function () {
    var that = this
    that.init();

    comments.getData(function(json) {
      that.setData({
        info: json.data
      });
    });
  },
  popup: function(e) {
    var that = this;

    that.setData({
      modalShow: true
    });
  },
  bindKeyInput: function(e) {
    var that = this;

    that.setData({
      content: e.detail.value
    });
  },
  modalSubmit: function(e) {
    var that = this;
    that.setData({
      modalShow: false
    })
  },
  modalCancel: function(e) {
    var that = this;
    that.setData({
      modalShow: false
    })
  },
})