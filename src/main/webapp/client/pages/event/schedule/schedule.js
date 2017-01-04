var constant = require('../../../utils/constant.js')
var util = require('../../../utils/util.js')
var schedule = require('../../../services/schedule.js')

var app = getApp()
Page({
  data: {
    mode: 'bySession'
  },

  onLoad: function () {
    var that = this

    schedule.getData(function(json) {
      that.setData({
        all: json
      });
      that.show();
    });
  },

  show: function () {
    var that = this;
    var timelineData;

    that.setData({
      timelineData: {mode: that.data.mode, list : that.data.all[that.data.mode]}
    });
    console.log(that.data.timelineData);
  },

  changBy: function (event) {
    var that = this;

    var by = event.currentTarget.dataset.by;
    console.log(by);

    that.setData({
      mode: by
    });
    that.show();
  },
  
})