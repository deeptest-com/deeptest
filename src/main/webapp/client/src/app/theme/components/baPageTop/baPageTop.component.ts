import {Component, ViewEncapsulation} from '@angular/core';
import {Router, Routes, NavigationEnd} from '@angular/router';
import {Subscription} from 'rxjs/Rx';

import {GlobalState} from '../../../global.state';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';
import { RouteService } from '../../../service/route';
import { AccountService } from '../../../service/account';

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
  projects: any[] = CONSTANT.RECENT_PROJECT;

  constructor(private _router:Router, private _state:GlobalState, private _routeService: RouteService, private accountService: AccountService) {
    let that = this;

    if (!CONSTANT.PROFILE) {
      this._state.subscribe('recent.projects.change', (projects) => {
        if (projects) {
          this.projects = projects;
          console.log('recent.projects.change', this.projects);
        }
      });

      that._state.subscribe('profile.refresh', (profile) => {
        that.profile = profile;
        console.log('profile.refresh', that.profile);
      });

      this.accountService.loadProfileRemote().subscribe((data: any) => {
        console.log('data', data);
      });
    }

    this._onRouteChange = this._router.events.subscribe((event) => {
      // if (event instanceof NavigationEnd) {
      //   if (event.url && event.url.indexOf('im=true') > -1) {
      //     return;
      //   }
      //   let title = Utils.getRouterUrlParam(event.url, 'title');
      //
      //   this._state.notifyDataChanged('menu.change', title);
      // }
    });

  }

  public scrolledChanged(isScrolled) {
    this.isScrolled = isScrolled;
  }

  logout() {
    this.accountService.logout();
  }

}
