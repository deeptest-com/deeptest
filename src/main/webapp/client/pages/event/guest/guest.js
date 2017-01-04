var constant = require('../../../utils/constant.js')
var util = require('../../../utils/util.js')
var guest = require('../../../services/guest.js')

var app = getApp()
Page({
  data: {
    url: constant.serviceUrl,
  },

  onLoad: function (options) {
    var that = this;
    console.log('onLoad');

    guest.list(function(json) {
      that.setData({
        guests: json.guests
      });
    });
  },

  gotoDetail: function(event) {
    var guestId = event.currentTarget.dataset.guestId;
    console.log(guestId);
    
    wx.navigateTo({
      url: './detail?guestId=' + guestId
    })
  }

})
