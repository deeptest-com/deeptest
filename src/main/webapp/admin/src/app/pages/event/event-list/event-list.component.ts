import {Component, ViewEncapsulation} from '@angular/core';

import { NgModule, Pipe, OnInit, AfterViewInit }      from '@angular/core';
import { Router } from '@angular/router';

import { CONSTANT } from '../../../utils/constant';
import { EventService } from '../../../service/event';

@Component({
  selector: 'event-list',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./event-list.scss')],
  template: require('./event-list.html')
})
export class EventList implements OnInit, AfterViewInit {
  maxSize:number = 5;

  totalItems:number = 0;
  currentPage:number = 1;
  itemsPerPage:number = 6;

  model: any = {status: ''};
  statusMap: Array<any> = CONSTANT.EventStatus;
  events: Array<any> = [];

  constructor(private _router: Router,
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

    that._router.navigateByUrl("/pages/event/edit/null/property");
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
