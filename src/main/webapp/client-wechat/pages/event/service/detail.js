var constant = require('../../../utils/constant.js')
var util = require('../../../utils/util.js')
var service = require('../../../services/service.js')

var app = getApp()
Page({
  data: {
    
  },

  onLoad: function (options) {
    var that = this
    
    var typeId = options.typeId;
    console.log(typeId);

    service.get(typeId, function(json) {
      that.setData({
        service: json.service,
        article: {article: json.service.descr}
      });
    });
  }
})