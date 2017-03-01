import {Component, ViewEncapsulation} from '@angular/core';

import { NgModule, Pipe, OnInit, AfterViewInit }      from '@angular/core';

import {NodeEvent,
  NodeMovedEvent, NodeMovedRemoteEvent,
  NodeRemovedEvent, NodeRemovedRemoteEvent,
  NodeCreatedEvent, NodeCreatedRemoteEvent,
  NodeRenamedEvent, NodeRenamedRemoteEvent,
  NodeSelectedEvent, TreeModel } from '../../components/ng2-tree';

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

    public onNodeRemovedRemote(e: NodeRemovedRemoteEvent): void {
        let that = this;
        this.logEvent(e, 'NodeRemovedRemoteEvent');
        that._testcaseService.delete(e.node.node).subscribe((json:any) => {
            this._treeService.fireNodeRemoved(e.node);
        });
    }

    public onNodeMovedRemote(e: NodeMovedRemoteEvent): void {
        let that = this;
        this.logEvent(e, 'NodeMovedRemoteEvent');
        that._testcaseService.move(e.node.node, e.srcTree.node, e.options).subscribe((json:any) => {
          this._treeService.fireNodeMoved(e.node, e.srcTree, e.options);
        });
    }

    public onNodeRenamedRemote(e: NodeRenamedEvent): void {
        let that = this;
        this.logEvent(e, 'NodeRenamedEvent');
        that._testcaseService.rename(e.node.node).subscribe((json:any) => {

        });
    }

    public onNodeCreatedRemote(e: NodeCreatedEvent): void {
        let that = this;
        that.logEvent(e, 'NodeCreatedEvent');
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

