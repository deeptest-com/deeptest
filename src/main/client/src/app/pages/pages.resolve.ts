import { Injectable } from '@angular/core';
import { Router, ActivatedRoute, Resolve, ActivatedRouteSnapshot } from '@angular/router';
import {Location} from '@angular/common';
import 'rxjs/add/operator/toPromise';

import {GlobalState} from '../global.state';
import { CONSTANT } from '../utils/constant';
import { Utils } from '../utils/utils';
import { SockService } from '../service/sock';
import { UserService } from '../service/user';

@Injectable()
export class PagesResolve implements Resolve<any> {
  constructor(private location: Location, private _state: GlobalState,
              private _sockService: SockService, private userService: UserService, private router: Router) { }

  resolve(route: ActivatedRouteSnapshot) {
    let context = Utils.getOrgAndPrjId(this.location.path());

    return this.userService.loadProfileRemote(context).toPromise().then(result => {
      console.log('PagesResolve resolve');
      this._sockService.wsConnect();

      this._state.notifyDataChanged(CONSTANT.EVENT_LOADING_COMPLETE, {});

      return result;
    });
  }
}
