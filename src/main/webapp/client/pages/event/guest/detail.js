var constant = require('../../../utils/constant.js')
var util = require('../../../utils/util.js')
var guest = require('../../../services/guest.js')

var app = getApp()
Page({
  data: {
    url: constant.serviceUrl
  },

  onLoad: function (options) {
    var that = this;
    console.log('onLoad');

    console.log(options.guestId);
    guest.get(options.guestId, function(json) {
      that.setData({
        guest: json.guest
      });
    });
  }
})