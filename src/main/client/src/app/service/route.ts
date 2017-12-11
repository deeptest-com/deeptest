import {Injectable} from "@angular/core";
import {Router} from "@angular/router";

import {CONSTANT} from '../utils/constant';

@Injectable()
export class RouteService {
  constructor(private _router: Router) {
  }

  navTo(url: string) {
    let that = this;
    // let urlStr = that._router.createUrlTree([url, {im: true}]);
    let urlStr = that._router.createUrlTree([url]);
    that._router.navigateByUrl(urlStr);
  }

  nav(urls: string[]) {
    let that = this;
    let urlStr = that._router.createUrlTree(urls);
    that._router.navigateByUrl(urlStr);
  }

  quickJump(key: string) {
    let arr: string[] = key.split('-');
    if (arr.length >1 && arr[0].toLowerCase() === 'tc') {
      this.gotoCase(arr[1]);
    }
  }
  gotoCase(id: string) {
    let url = '/pages/org/' + CONSTANT.CURR_ORG_ID + '/prj/' + CONSTANT.CURR_PRJ_ID
      + '/design/case/' + id;
    this._router.navigateByUrl(url);
  }

  caseIdForJump(key: string) {
    let arr: string[] = key.split('-');
    if (arr.length >1 && arr[0].toLowerCase() === 'tc' && !!arr[1]) {
      return arr[1];
    }
    return null;
  }

}

