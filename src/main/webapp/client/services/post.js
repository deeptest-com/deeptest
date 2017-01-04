var constant = require('../utils/constant.js')

function post(path, data, callback) {
    var url = constant.apiUrl + path + '?token=' + constant.token;
    console.log(url);

     wx.request({
        url: url,
        data: data,
        method: 'POST',
        header: {
            'Content-Type': 'application/json'
        },
        success: function(json) {
          console.log(json);
          if (json.data.code === 1) {
              callback(json.data);
          } else {
              console.log('request error')
          }
        }
      });
}
module.exports = {
  post: post
}