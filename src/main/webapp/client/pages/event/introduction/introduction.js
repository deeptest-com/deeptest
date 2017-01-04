var constant = require('../../../utils/constant.js')
var util = require('../../../utils/util.js')
var event = require('../../../services/event.js')

var app = getApp()
Page({
  data: {
    
  },

  onLoad: function (options) {
    var that = this;
    console.log('onLoad');

    event.get(function(json) {
      that.setData({
        event: json.event
      });
    });

    // 连接WebSocket
    var url = constant.serviceUrl + constant.websocketPath + "?t=" + new Date().getTime();
    url = url.replace('http', 'ws');
    wx.connectSocket({
      url: url,
      data:{
        clientId: -1,
        eventId: -1
      },
      header:{ 
        'content-type': 'application/json'
      },
      success: function() {console.log('WebSocket连接成功')}, 
      fail:  function() {console.log('WebSocket连接失败')}, 
      complete:  function() {console.log('WebSocket连接结束')}
    });
    wx.onSocketOpen(function(res){
      console.log('WebSocket连接打开');
      var data = {
        act: 'enter_chat_room',
        eventId: -1,
        clientId: -1
      };
      wx.sendSocketMessage({
        data: JSON.stringify(data)
      })
    });
    wx.onSocketError(function(res){
      console.log('WebSocket连接错误');
    });
    wx.onSocketMessage(function(res) {
      console.log('WebSocket消息：' + res.data);
    })
  }

})