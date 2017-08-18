import {Component} from '@angular/core';

import {Router} from '@angular/router';
import {Subscription} from 'rxjs/Rx';

import {GlobalState} from '../../../global.state';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';
import { RouteService } from '../../../service/route';
import {OrgService} from "../../../service/org";
import { AccountService } from '../../../service/account';

@Component({
  selector: 'ba-page-top',
  templateUrl: './baPageTop.html',
  styleUrls: ['./baPageTop.scss']
})
export class BaPageTop {
  public profile:any = CONSTANT.PROFILE;
  project: any = CONSTANT.CURRENT_PROJECT;
  projects: any[] = CONSTANT.RECENT_PROJECTS;
  myOrgs: any[] = CONSTANT.MY_ORGS;

  public isScrolled:boolean = false;
  public isMenuCollapsed:boolean = false;
  isOrgsShow:boolean = false;

  constructor(private _router:Router, private _state:GlobalState, private _routeService: RouteService,
              private orgService: OrgService, private accountService: AccountService) {
    let that = this;

    if (!CONSTANT.PROFILE) {
      this._state.subscribe('menu.isCollapsed', (isCollapsed) => {
        this.isMenuCollapsed = isCollapsed;
      });

      this._state.subscribe('my.orgs.change', (myOrgs) => {
        console.log('my.orgs.change', myOrgs);
        if (myOrgs) {
          this.myOrgs = myOrgs;
        }
      });

      this._state.subscribe('recent.projects.change', (projects) => {
        console.log('recent.projects.change', projects);
        if (projects) {
          this.projects = projects;
          this.project = CONSTANT.CURRENT_PROJECT;
        }
      });

      that._state.subscribe('profile.refresh', (profile) => {
        console.log('profile.refresh', profile);
        that.profile = profile;
      });

      this.accountService.loadProfileRemote().subscribe((data: any) => {
        console.log('data', data);
      });
    }

    this._state.subscribe('menu.isCollapsed', (isCollapsed) => {
      this.isMenuCollapsed = isCollapsed;
    });
  }

  public changeOrg(item: any) {
    console.log(item);

    this.orgService.setDefault(item.id, {disabled: false}).subscribe((json:any) => {
      if (json.code == 1) {
        this.accountService.changeRecentProject(json.recentProjects);
      }
    });
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
