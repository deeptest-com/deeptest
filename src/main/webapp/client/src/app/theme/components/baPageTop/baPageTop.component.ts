import {Component} from '@angular/core';

import {Router} from '@angular/router';
import {Subscription} from 'rxjs/Rx';

import {GlobalState} from '../../../global.state';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';
import { RouteService } from '../../../service/route';
import { AccountService } from '../../../service/account';

@Component({
  selector: 'ba-page-top',
  templateUrl: './baPageTop.html',
  styleUrls: ['./baPageTop.scss']
})
export class BaPageTop {
  public profile:any = CONSTANT.PROFILE;
  projects: any[] = CONSTANT.RECENT_PROJECTS;

  public isScrolled:boolean = false;
  public isMenuCollapsed:boolean = false;

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

    this._state.subscribe('menu.isCollapsed', (isCollapsed) => {
      this.isMenuCollapsed = isCollapsed;
    });
  }

  public toggleMenu() {
    this.isMenuCollapsed = !this.isMenuCollapsed;
    this._state.notifyDataChanged('menu.isCollapsed', this.isMenuCollapsed);
    return false;
  }

  public scrolledChanged(isScrolled) {
    this.isScrolled = isScrolled;
  }

  gotoModule(module: string) {
    console.log('CONSTANT.CURRENT_PROJECT', CONSTANT.CURRENT_PROJECT);

    let url = '';
    if (module == 'design') {
      url = '/pages/design/' + CONSTANT.CURRENT_PROJECT.id +'/case';
    } else if(module == 'implement') {
      url = '/pages/implement/' + CONSTANT.CURRENT_PROJECT.id + '/plan/list';
    } else if (module == 'analysis') {
      url = '/pages/analysis/' + CONSTANT.CURRENT_PROJECT.id + '/report/list';
    }

    this._routeService.navTo(url);

  }
}
