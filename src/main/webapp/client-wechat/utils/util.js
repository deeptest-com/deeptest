var constant = require('./constant.js')

function formatTime(date) {
  var year = date.getFullYear()
  var month = date.getMonth() + 1
  var day = date.getDate()

  var hour = date.getHours()
  var minute = date.getMinutes()
  var second = date.getSeconds()


  return [year, month, day].map(formatNumber).join('/') + ' ' + [hour, minute, second].map(formatNumber).join(':')
}

function formatNumber(n) {
  n = n.toString()
  return n[1] ? n : '0' + n
}


function getScreenSize () {
  var sh = window.screen.height;
  if (document.body.clientHeight < sh) {
    sh = document.body.clientHeight;
  }
  
  var sw = window.screen.width;
  if (document.body.clientWidth < sw) {
    sw = document.body.clientWidth;
  }
  
  return {h: sh, w: sw};
}
function setScreanSize (){
  var size = getScreenSize();
  constant.W = size.w;
  constant.H = size.h;
}

module.exports = {
  formatTime: formatTime,
  setScreanSize: setScreanSize
}