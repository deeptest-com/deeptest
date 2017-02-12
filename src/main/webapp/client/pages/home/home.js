var constant = require('../../utils/constant.js')
var util = require('../../utils/util.js')
var sign = require('../../services/sign.js')
var home = require('../../services/home.js')

var app = getApp()
Page({
  data: {url: constant.serviceUrl},

  gotoPage: function(event) {
    var page = event.currentTarget.dataset.name;
    console.log(page);

    if(page == 'sign') {
      sign.sign();
    } else {
      wx.navigateTo({
        url: '../event/' + page + '/' + page
      });
    }
  },
  onLoad: function () {
    var that = this
    console.log('onLoad');

    that.init();
    
    home.getData(function(json) {
      constant.event = json.event;

      var label, name;
      if (constant.event.status == 'in_progress') {
        label = '现场';
        name = 'live';
      } else if (constant.event.status == 'end') {
        label = '回顾';
        name = 'review';
      } else if (constant.event.status == 'sign') {
        label = '签到';
        name = 'sign';
      } else {
        label = '报名';
        name = 'register';
      }
      that.setData({
        event: constant.event,
        buttons: [
            {label: '介绍', name: 'introduction'},
            {label: '日程', name: 'schedule'},
            {label: '嘉宾', name: 'guest'},
            
             {label: label, name: name},
             //{label: '现场', name: 'live'},
             //{label: '回顾', name: 'review'},
             //{label: '签到', name: 'sign'},
            
            {label: '服务', name: 'service'},
            {label: '问答', name: 'qa'}
          ]
      });
    });
  },
  init: function () {
    var that = this
    var width = 750;
    var colHeight = width / 3;
    
    that.setData({
      gridHeight: colHeight * 2,
      colHeight: colHeight
    });
  },

  onReady: function () {
    var that = this;
    console.log('onReady');
    that.setNavTtile();
  },

  onShow: function () {
    var that = this;
    console.log('onShow');
    that.setNavTtile();
  },

  setNavTtile: function(title){
    var time = 0;
    var interval = setInterval(function() {
      if (constant.event) {
        wx.setNavigationBarTitle({
          title: constant.event.title
        });
        clearInterval(interval);
      } else {
        if (time++ > 10) {
          clearInterval(interval);
        }
      }
    }, 50)
  }
})