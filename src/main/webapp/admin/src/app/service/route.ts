import {Injectable} from '@angular/core';

import { Router, ActivatedRoute, Params } from '@angular/router';

@Injectable()
export class RouteService {
    constructor(private _router: Router) { }

    navTo(url: string) {
      let that = this;
      let urlStr = that._router.createUrlTree([url, {im: true}]);
      that._router.navigateByUrl(urlStr);
    }
}

