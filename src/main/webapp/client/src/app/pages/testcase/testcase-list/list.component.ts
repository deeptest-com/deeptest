import {Component, ViewEncapsulation} from '@angular/core';

import { NgModule, Pipe, OnInit, AfterViewInit }      from '@angular/core';

import { NodeEvent, TreeModel, RenamableNode } from '../../components/ng2-tree';

import {GlobalState} from '../../../global.state';

import { CONSTANT } from '../../../utils/constant';
import { RouteService } from '../../../service/route';
import { TestcaseService } from '../../../service/testcase';

@Component({
  selector: 'testcase-list',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./list.scss'), require('../../components/ng2-tree/src/styles.css')],
  template: require('./list.html')
})
export class TestcaseList implements OnInit, AfterViewInit {
  totalItems:number = 0;
  currentPage:number = 1;
  itemsPerPage:number = 6;

  model: any = {status: ''};
  statusMap: Array<any> = CONSTANT.EventStatus;
  events: Array<any> = [];

  public tree: TreeModel = {
    value: 'Programming languages by programming paradigm',
    children: [
      {
        value: 'Object-oriented programming',
        children: [
          {value: 'Java'},
          {value: 'C++'},
          {value: 'C#'},
        ]
      },
      {
        value: 'Prototype-based programming',
        children: [
          {value: 'JavaScript'},
          {value: 'CoffeeScript'},
          {value: 'Lua'},
        ]
      }
    ]
  };

  constructor(private _routeService: RouteService, private _state:GlobalState,
              private _testcaseService: TestcaseService) {

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
    // that._testcaseService.list(that.itemsPerPage, that.currentPage, that.model.status).subscribe((json:any) => {
    //   that.totalItems = json.totalItems;
    //   that.events = json.events;

      this._state.notifyDataChanged('title.change', '测试用例');
    // });
  }

    public onNodeRemoved(e: NodeEvent): void {
        this.logEvent(e, 'Removed');
    }

    public onNodeMoved(e: NodeEvent): void {
        this.logEvent(e, 'Moved');
    }

    public onNodeRenamed(e: NodeEvent): void {
        this.logEvent(e, 'Renamed');
    }

    public onNodeCreated(e: NodeEvent): void {
        this.logEvent(e, 'Created');
    }

    public onNodeSelected(e: NodeEvent): void {
        this.logEvent(e, 'Selected');
    }

    public logEvent(e: NodeEvent, message: string): void {
        console.log(e, message);
        console.log(this.tree);
    }
}
