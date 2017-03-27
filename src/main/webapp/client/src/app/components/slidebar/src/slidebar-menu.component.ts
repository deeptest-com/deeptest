import {Component, ViewEncapsulation, Input, Output, EventEmitter} from '@angular/core';
import {Router, Routes, NavigationEnd} from '@angular/router';

import { RouteService } from '../../../service/route';

declare var jQuery;

@Component({
  selector: 'slidebar-menu',
  encapsulation: ViewEncapsulation.None,
  styles: [],
  template: require('./slidebar-menu.html')
})
export class SlidebarMenu {
  @Input()
  public menuItems: any[];

  public showHoverElem:boolean;
  public hoverElemHeight:number;
  public hoverElemTop:number;
  public outOfArea:number = -200;

  constructor(private _routeService: RouteService) {

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
    _.forEach(this.menuItems, (item: any, index: number) => {
      item.selected = false;
    });
    $event.item.selected = true;
    this._routeService.navTo($event.item.link);
  }
}
