import {Injectable} from '@angular/core';

import {CONSTANT} from './constant';

@Injectable()
export class Utils {
  constructor() { }

  static ClientCofig (){
    Utils.SetScreanSize();
    Utils.SetServiceUrl();
    Utils.SetPlatformType();
  }

  static SetScreanSize (){
    CONSTANT.W = Utils.getScreenSize().w;
    CONSTANT.H = Utils.getScreenSize().h;
  }
  static SetServiceUrl (){
    let host = window.location.host;
    if (!CONSTANT.SERVICE_URL) {
      if (host.indexOf("localhost") > -1 || host.indexOf("127.") > -1 || host.indexOf("10.") > -1) {
        CONSTANT.SERVICE_URL = CONSTANT._SERVICE_URL_DEV;
      } else {    // production
        CONSTANT.SERVICE_URL = CONSTANT._SERVICE_URL_PRODUCTION;
      }
    }
  }

  static SetPlatformType (){
    let sUserAgent = navigator.userAgent;
    // Android iPhone iPad

    if ( sUserAgent.toLowerCase().indexOf('android') > -1 ) {
      CONSTANT.PLATFORM = 'ios';
    } else if ( sUserAgent.toLowerCase().indexOf('iphone') > -1
      || sUserAgent.toLowerCase().indexOf('ipad') > -1 ) {
      CONSTANT.PLATFORM = 'android';
    }
    console.log('CONSTANT.PLATFORM = ' + CONSTANT.PLATFORM);
  }

  static IsAndroid(){
    return CONSTANT.PLATFORM === 'android';
  }
  static IsIos(){
    return CONSTANT.PLATFORM === 'ios';
  }
  static ImgUrl(url: string, external: boolean){
    if (!url) {
//        if (defaultt) {
//            return defaultt;
//        } else {
            return 'assets/img/none.png';
//        }
    }

    if (!external)  {
        external = true;
    }

    if (external) {
        url = CONSTANT.SERVICE_URL + url;
    }
    return url;
  }

  static backgroundImage(url: string, external: boolean): any {
    if (!url) {
       return {};
    }

    if (external) {
        url = CONSTANT.SERVICE_URL + url;
    }
    return {'background-image': 'url(' + url + ')'};
  }

  static ImgSize(url: string){
    if (CONSTANT.W > 640) {
        url = url.replace('.', '@2x.');
    }
    return url;
  }
    
  static getScreenSize () {
    var sh = window.screen.height;
    if (document.body.clientHeight < sh) {
      sh = document.body.clientHeight;
    }
    
    var sw = window.screen.width;
    if (document.body.clientWidth < sw) {
      sw = document.body.clientWidth;
    }
    
    //if (this.landscape && this.landscape() && sh > sw) {
    //    var temp = sh;
    //    sh = sw;
    //    sw = temp;
    //}
    
    return {h: sh, w: sw};
  }
  
}
