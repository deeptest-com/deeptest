import {Component, ViewEncapsulation} from '@angular/core';

import {GlobalState} from '../../../global.state';

import { Cookie } from 'ng2-cookies/ng2-cookies';

import { CONSTANT } from '../../../utils/constant';
import { RouteService } from '../../../service/route';
import { UserService } from '../../../service/user';

@Component({
  selector: 'ba-page-top',
  styles: [require('./baPageTop.scss')],
  template: require('./baPageTop.html'),
  encapsulation: ViewEncapsulation.None
})
export class BaPageTop {

  public isScrolled:boolean = false;
  public isMenuCollapsed:boolean = false;
  public profile:any = CONSTANT.PROFILE;

  constructor(private _state:GlobalState, private userService: UserService, private routeService: RouteService) {
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

  logout() {

    this.userService.logout().subscribe((json:any) => {
      if (json.code == 1) {
        Cookie.delete(CONSTANT.PROFILE_KEY);

        this.routeService.navTo('/login');
      }
    });


  }
}
