import {Component, ViewEncapsulation} from '@angular/core';

import { NgModule, Pipe, OnInit, AfterViewInit }      from '@angular/core';

import { CONSTANT } from '../../../utils/constant';
import { RouteService } from '../../../service/route';
import { EventService } from '../../../service/event';

@Component({
  selector: 'event-list',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./list.scss')],
  template: require('./list.html')
})
export class EventList implements OnInit, AfterViewInit {
  totalItems:number = 0;
  currentPage:number = 1;
  itemsPerPage:number = 6;

  model: any = {status: ''};
  statusMap: Array<any> = CONSTANT.EventStatus;
  events: Array<any> = [];

  constructor(private _routeService: RouteService,
              private _eventService: EventService) {

  }
  ngOnInit() {
    let that = this;
    that.loadData();
  }

  ngAfterViewInit() {

  }

  pageChanged(event:any):void {
    let that = this;
    that.currentPage = event.page;
    that.loadData();
  }
  create():void {
    let that = this;

    that._routeService.navTo("/pages/event/edit/null/property");
  }
  statusChange(e: any):void {
    let that = this;
    that.model.status = e;
    that.loadData();
  }
  delete(eventId: string):void {
    let that = this;
    console.log('eventId=' + eventId);
  }

  loadData() {
    let that = this;
    that._eventService.list(that.itemsPerPage, that.currentPage, that.model.status).subscribe((json:any) => {
      that.totalItems = json.totalItems;
      that.events = json.events;
    });
  }
}
