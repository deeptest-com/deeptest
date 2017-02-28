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

  query: any = {keywords: '', status: ''};
    cases: Array<any> = [];

  public tree: TreeModel;

  constructor(private _routeService: RouteService, private _state:GlobalState,
              private _testcaseService: TestcaseService) {

  }
  ngOnInit() {
    let that = this;

    that.loadData();
  }

  ngAfterViewInit() {

  }

  create():void {
    let that = this;

    that._routeService.navTo("/pages/event/edit/null/property");
  }
  statusChange(e: any):void {
    let that = this;
    that.query.status = e;
    that.loadData();
  }
  delete(eventId: string):void {
    let that = this;
    console.log('eventId=' + eventId);
  }

  loadData() {
    let that = this;
    that._testcaseService.query(that.query).subscribe((json:any) => {
      that.tree = json.data;

      this._state.notifyDataChanged('title.change', '测试用例');
    });
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
    }
}
