var constant = require('../../../utils/constant.js')
var util = require('../../../utils/util.js')
var bizcard = require('../../../services/bizcard.js')

var app = getApp()
Page({
  data: {
    
  },

  onLoad: function () {
    var that = this
    
  },
  
  save: function () {
    var that = this;
    
    console.log('save');
  },

  formSubmit: function (e) {
    var that = this;
    
    console.log(e.detail);
  }
})