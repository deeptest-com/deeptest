import {Component, ViewEncapsulation} from '@angular/core';

import { NgModule, Pipe, OnInit, AfterViewInit }      from '@angular/core';

import {NodeEvent, NodeMovedEvent, NodeRemovedEvent, NodeDeletedEvent, NodeCreatedEvent, NodeRenamedEvent, NodeSelectedEvent, TreeModel } from '../../components/ng2-tree';

import {GlobalState} from '../../../global.state';

import { CONSTANT } from '../../../utils/constant';
import { RouteService } from '../../../service/route';
import { TreeService } from '../../components/ng2-tree/src/tree.service';

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
              private _treeService: TreeService, private _testcaseService: TestcaseService) {

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

    public onNodeRemoved(e: NodeRemovedEvent): void {
        let that = this;
        this.logEvent(e, 'Removed');
    }

    public onNodeDeleted(e: NodeDeletedEvent): void {
        let that = this;
        this.logEvent(e, 'Deleted');
        that._testcaseService.delete(e.node.node).subscribe((json:any) => {
            this._treeService.fireNodeRemoved(e.node);
        });
    }

    public onNodeMoved(e: NodeMovedEvent): void {
        let that = this;
        this.logEvent(e, 'Moved');
        that._testcaseService.move(e.node.node, e.previousParent.node, e.options).subscribe((json:any) => {

        });
    }

    public onNodeRenamed(e: NodeRenamedEvent): void {
        let that = this;
        this.logEvent(e, 'Renamed');
        that._testcaseService.rename(e.node.node).subscribe((json:any) => {

        });
    }

    public onNodeCreated(e: NodeCreatedEvent): void {
        let that = this;
        that.logEvent(e, 'Created');
        that._testcaseService.create(e.node.node).subscribe((json:any) => {

        });
    }

    public onNodeSelected(e: NodeSelectedEvent): void {
        this.logEvent(e, 'Selected');
    }

    public logEvent(e: NodeEvent, message: string): void {
        console.log(e, message);
    }
}

