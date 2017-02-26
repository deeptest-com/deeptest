import {Component, ViewEncapsulation} from '@angular/core';
import {Router, Routes, NavigationEnd} from '@angular/router';
import {Subscription} from 'rxjs/Rx';

import {GlobalState} from '../../../global.state';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';
import { UserService } from '../../../service/user';

@Component({
  selector: 'ba-page-top',
  styles: [require('./baPageTop.scss')],
  template: require('./baPageTop.html'),
  encapsulation: ViewEncapsulation.None
})
export class BaPageTop {
  protected _onRouteChange:Subscription;
  public isScrolled:boolean = false;
  public profile:any = CONSTANT.PROFILE;

  constructor(private _router:Router, private _state:GlobalState, private userService: UserService) {
    let that = this;
    this._onRouteChange = this._router.events.subscribe((event) => {
      if (event instanceof NavigationEnd) {
        if (event.url && event.url.indexOf('im=true') > -1) {
          return;
        }
        let title = Utils.getRouterUrlParam(event.url, 'title');

        this._state.notifyDataChanged('menu.change', title);
      }
    });

    that._state.subscribe('profile.refresh', (profile) => {
      that.profile = profile;
      console.log('profile.refresh', that.profile);
    });
  }

  public scrolledChanged(isScrolled) {
    this.isScrolled = isScrolled;
  }

  logout() {
    this.userService.logout();
  }
}
