import {Injectable} from '@angular/core';

import { NgbDateStruct } from '@ng-bootstrap/ng-bootstrap';

import {CONSTANT} from '../utils/constant';

declare var unescape;

export var Utils: any = {
  config: function() {
    var host = window.location.host;
    if (host.indexOf('localhost') > -1) {
      CONSTANT.SERVICE_URL = CONSTANT._SERVICE_URL_DEV;
    } else {
      CONSTANT.SERVICE_URL = CONSTANT._SERVICE_URL_PRODUCTION;
    }
    CONSTANT.API_URL = CONSTANT.SERVICE_URL + CONSTANT.API_PATH;
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
  dateStructToDate: function(struct: NgbDateStruct) {
    let date = new Date();
    date.setFullYear(struct.year, struct.month - 1, struct.day);

    return date;
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

  imgUrl:function(url: string, external: boolean, defaultt: string){
    if (!url) {
      if (defaultt) {
        url = defaultt
      } else {
        url = 'img/1-1.png';
      }
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
      url = 'img/none.png';
    }

    url = url.replace('.', '-thumb.');

    if (!external)  {
      external = true;
    }

    if (external) {
      url = CONSTANT.SERVICE_URL + url;
    }
    return url;
  },
  getUrlParam: function (pname) {
    var rt = '';
    var url = unescape(window.location.href);
    url = url.split('#')[1];

    var allArgs = url.split("?")[1];
    if (!allArgs) {
      return '';
    }
    var args = allArgs.split("&");
    for (var i = 0; i < args.length; i++) {
      var arg = args[i].split("=");
      if (arg[0] == pname) {
        console.log('find url param: ' + arg[0] + '="' + arg[1] + '";');
        rt = arg[1];
        return rt;
      }
    }
    return rt;
  },

  getRouterUrlParam: function (url: string, param: string) {
    var rt = '';
    let reg = new RegExp('.*' + param + '=(.*)(;|$)');
    let r = url.match(reg);
    if(r != null) rt = decodeURIComponent(r[1]);
    return rt;
  },

  getScreenSize: function() {
    var sh = window.screen.height;
    if (document.body.clientHeight < sh) {
      sh = document.body.clientHeight;
    }

    var sw = window.screen.width;
    if (document.body.clientWidth < sw) {
      sw = document.body.clientWidth;
    }

    return {h: sh, w: sw};
  },

  getContainerHeight: function (h: number) {
    return CONSTANT.ScreenSize.h - h + 'px'
  },

  getOrgAndPrjId: function (url: string) {
    let orgId, prjId;

    // #/pages/org/139/prj/179/implement/plan/list
    if (url.indexOf('pages/org/') > -1) {
      let str = url.split('org/')[1];
      orgId = str.split('/')[0];
      if (str.indexOf('prj/') > -1) {
        prjId = str.split('prj/')[1].split('/')[0];;
      }
    }
    let ret = {orgId: orgId, prjId: prjId};
    console.log('url params: ', ret);
    return ret;
  }

};

export class Deferred {

  promise: Promise<any>;

  resolve: any;
  reject: any;

  constructor() {
    this.promise = new Promise((resolve, reject) => {
      this.resolve = resolve;
      this.reject = reject;
    });
  }
}
