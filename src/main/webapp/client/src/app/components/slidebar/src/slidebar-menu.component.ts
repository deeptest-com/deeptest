import {Component, ViewEncapsulation, Input, Output, EventEmitter} from '@angular/core';
import {Router, Routes, NavigationEnd} from '@angular/router';
import {Subscription} from 'rxjs/Rx';

import { RouteService } from '../../../service/route';

declare var jQuery;

@Component({
  selector: 'slidebar-menu',
  encapsulation: ViewEncapsulation.None,
  styles: [],
  template: require('./slidebar-menu.html')
})
export class SlidebarMenu {
  protected _onRouteChange:Subscription;

  @Input()
  public menuItems: any;
  currLink: string;

  public showHoverElem:boolean;
  public hoverElemHeight:number;
  public hoverElemTop:number;
  public outOfArea:number = -200;

  constructor(private _router:Router, private _routeService: RouteService) {
    this.currLink = _router.url;
  }

  public ngOnInit():void {

  }

  public ngOnDestroy():void {

  }

  public hoverItem($event):void {
    this.showHoverElem = true;
    this.hoverElemHeight = $event.currentTarget.clientHeight;
    // TODO: get rid of magic 66 constant
    this.hoverElemTop = $event.currentTarget.getBoundingClientRect().top - 66;
  }

  public selectItem($event):void {
    this.currLink = $event.item.key;
    this._routeService.navTo($event.item.key);
  }
}
