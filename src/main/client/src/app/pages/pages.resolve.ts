import { Injectable } from '@angular/core';
import { Router, ActivatedRoute, Resolve, ActivatedRouteSnapshot } from '@angular/router';

import 'rxjs/add/operator/toPromise';

import { AccountService } from '../service/account';

@Injectable()
export class PagesResolve implements Resolve<any> {
  constructor(private _route: ActivatedRoute, private accountService: AccountService, private router: Router) { }

  resolve(route: ActivatedRouteSnapshot) {
    return this.accountService.loadProfileRemote().toPromise().then(result => {
      return result;
    });
  }
}
