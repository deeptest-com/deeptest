import {Component, ViewEncapsulation, Input, Output, EventEmitter} from '@angular/core';
import {Router, Routes, NavigationEnd} from '@angular/router';
import {Subscription} from 'rxjs/Rx';

import { CONSTANT } from '../../../utils/constant';
import {GlobalState} from "../../../global.state";
import { RouteService } from '../../../service/route';

declare var jQuery;

@Component({
  selector: 'slidebar-menu',
  templateUrl: './slidebar-menu.html'
})
export class SlidebarMenu {
  isOrgAdmin: boolean;

  @Input()
  public menuItems: any;
  currLink: string;

  public showHoverElem:boolean;
  public hoverElemHeight:number;
  public hoverElemTop:number;
  public outOfArea:number = -200;

  constructor(private _router:Router, private _state: GlobalState, private _routeService: RouteService) {
    this.currLink = _router.url;

    if (CONSTANT.PROFILE) {
      this.isOrgAdmin = CONSTANT.PROFILE.orgPrivilege.org_admin;
    }

    this._state.subscribe(CONSTANT.STATE_CHANGE_PROFILE, (profile) => {
      console.log(CONSTANT.STATE_CHANGE_PROFILE, profile);

      this.isOrgAdmin = CONSTANT.PROFILE.orgPrivilege.org_admin;
    });
    this._state.subscribe(CONSTANT.STATE_CHANGE_ORGS, (data) => {
      console.log(CONSTANT.STATE_CHANGE_ORGS, data);

      this.isOrgAdmin = CONSTANT.PROFILE.orgPrivilege.org_admin;
    });
  }

  public hoverItem($event):void {
    this.showHoverElem = true;
    this.hoverElemHeight = $event.currentTarget.clientHeight;
    // TODO: get rid of magic 66 constant
    this.hoverElemTop = $event.currentTarget.getBoundingClientRect().top - 66;
  }

  public selectItem($event):void {
    this.currLink = $event.item.link;
    this._routeService.navTo($event.item.link);
  }
}
