import {Component, ViewEncapsulation, OnInit, AfterViewInit} from "@angular/core";
import { Router, ActivatedRoute, Params } from '@angular/router';

import {
  NodeEvent,
  NodeMovedRemoteEvent,
  NodeRemovedRemoteEvent,
  NodeCreatedEvent,
  NodeRenamedEvent,
  NodeSelectedEvent,
  TreeModel,
  TreeOptions
} from "../../../../components/ng2-tree";

import {GlobalState} from "../../../../global.state";
import {CONSTANT} from "../../../../utils/constant";
import {Utils} from "../../../../utils/utils";
import {RouteService} from "../../../../service/route";
import {SlimLoadingBarService} from "../../../../components/ng2-loading-bar";
import {RunService} from "../../../../service/run";
import {CaseService} from "../../../../service/case";

@Component({
  selector: 'execution-suite',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./suite.scss',
    '../../../../../vendor/ztree/css/zTreeStyle/zTreeStyle.css',
    '../../../../components/ztree/src/styles.scss'],
  templateUrl: './suite.html'
})
export class ExecutionSuite implements OnInit, AfterViewInit {
  runId: number;
  projectId: number;
  public treeModel: any;
  public treeSettings: any = {usage: 'exe', isExpanded: true, sonSign: false};

  constructor(private _routeService:RouteService, private _route: ActivatedRoute, private _state:GlobalState,
              private _runService: RunService, private _caseService: CaseService,
              private slimLoadingBarService:SlimLoadingBarService) {

  }

  ngOnInit() {
    this._route.params.forEach((params: Params) => {
      this.projectId = +params['projectId'];
      this.runId = +params['runId'];
    });

    this.loadData();
  }

  ngAfterViewInit() {

  }

  loadData() {
    this.startLoading();

    this._caseService.queryForSelection(CONSTANT.CURRENT_PROJECT.id, this.runId).subscribe((json:any) => {
      this.treeModel = json.data;
      this.completeLoading();
    });

  }

  public onNodeSelected(e:NodeSelectedEvent):void {
    this._state.notifyDataChanged('exe.suite.change', {id: e.node.node.id, tm: new Date().getTime()});
  }

  startLoading() {
    this.slimLoadingBarService.start(() => {
      console.log('Loading complete');
    });
  }

  completeLoading() {
    let that = this;
    setTimeout(function () {
      that.slimLoadingBarService.complete();
    }, 500);
  }

}

