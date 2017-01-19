import {Injectable} from '@angular/core';

import {CONSTANT} from '../utils/constant';
import { Cookie } from 'ng2-cookies/ng2-cookies';

export var Utils: any = {
  config: function() {
    var host = window.location.host;
    if (host.indexOf('localhost') > -1) {
      CONSTANT.SERVICE_URL = CONSTANT._SERVICE_URL_DEV;
    } else {
      CONSTANT.SERVICE_URL = CONSTANT._SERVICE_URL_PRODUCTION;
    }
    CONSTANT.API_URL = CONSTANT.SERVICE_URL + CONSTANT.API_PATH;

    CONSTANT.TOKEN = Cookie.get(CONSTANT.COOKIE_KEY);
  },
  getUploadUrl: function() {
    return CONSTANT.API_URL + CONSTANT.UPLOAD_URI;
  },

  strToDate: function(str: string) {
    return new Date(str);
  },
  strToTimestamp: function(str: string) {
      return new Date(str).getTime();
  },
  timestampToStr: function(tm: string, fmt: string) {
      return Utils.dateToStr(new Date(tm), fmt);
  },
  dateToStr: function(date: any, fmt: string) {
      var o = {
        "M+" : date.getMonth()+1,                 //月份
        "d+" : date.getDate(),                    //日
        "h+" : date.getHours(),                   //小时
        "m+" : date.getMinutes(),                 //分
        "s+" : date.getSeconds(),                 //秒
        "q+" : Math.floor((date.getMonth()+3)/3), //季度
        "S"  : date.getMilliseconds()             //毫秒
      };
      if(/(y+)/.test(fmt))
        fmt=fmt.replace(RegExp.$1, (date.getFullYear()+"").substr(4 - RegExp.$1.length));
      for(var k in o)
        if(new RegExp("("+ k +")").test(fmt))
      fmt = fmt.replace(RegExp.$1, (RegExp.$1.length==1) ? (o[k]) : (("00"+ o[k]).substr((""+ o[k]).length)));
      return fmt;
  },

  dateDivide: function(model: any, dateKey: any, timeKey: any, datetimeKey: string) {
    let dateStr = Utils.timestampToStr(model[datetimeKey], "yyyy-MM-dd hh:mm");
    let arr = dateStr.split(' ');
    model[dateKey] = arr[0];
    model[timeKey] = arr[1];
  },

  dateCombine: function(model: any, dateKey: any, timeKey: any, datetimeKey: string) {
    let dateStr = model[dateKey] + ' ' + model[timeKey];
    model[datetimeKey] = Utils.strToDate(dateStr);
  },

  imgUrl:function(url: string, external: boolean){
    if (!url) {
        return 'assets/img/none.png';
    }

    if (!external)  {
        external = true;
    }

    if (external) {
        url = CONSTANT.SERVICE_URL + url;
    }
    return url;
  },

  thumbUrl:function(url: string, external: boolean){
    if (!url) {
      return 'assets/img/none.png';
    }

    url = url.replace('.', '-thumb.');

    if (!external)  {
      external = true;
    }

    if (external) {
      url = CONSTANT.SERVICE_URL + url;
    }
    return url;
  }

};
