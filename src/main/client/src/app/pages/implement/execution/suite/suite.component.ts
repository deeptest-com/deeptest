import {Component, ViewEncapsulation, OnInit, AfterViewInit} from "@angular/core";
import { Router, ActivatedRoute, Params } from '@angular/router';

import {GlobalState} from "../../../../global.state";
import {CONSTANT} from "../../../../utils/constant";
import {Utils} from "../../../../utils/utils";
import {RouteService} from "../../../../service/route";
import {SlimLoadingBarService} from "../../../../components/ng2-loading-bar";
import {RunService} from "../../../../service/run";
import {CaseService} from "../../../../service/case";
import { CaseInRunService } from '../../../../service/case-in-run';

@Component({
  selector: 'execution-suite',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./suite.scss',
    '../../../../../assets/vendor/ztree/css/zTreeStyle/zTreeStyle.css',
    '../../../../components/ztree/src/styles.scss'],
  templateUrl: './suite.html'
})
export class ExecutionSuite implements OnInit, AfterViewInit {
  runId: number;
  projectId: number;
  public treeModel: any;
  public treeSettings: any = {usage: 'exe', isExpanded: true, sonSign: false};

  constructor(private _routeService:RouteService, private _route: ActivatedRoute, private _state:GlobalState,
              private _runService: RunService, private _caseInRunService: CaseInRunService,
              private slimLoadingBarService:SlimLoadingBarService) {

  }

  ngOnInit() {
    this._route.params.forEach((params: Params) => {
      this.runId = +params['runId'];
    });

    this.projectId = CONSTANT.CURR_PRJ_ID;
    this.loadData();
  }

  ngAfterViewInit() {

  }

  loadData() {
    this.startLoading();

    this._caseInRunService.query(this.projectId, this.runId).subscribe((json:any) => {
      this.treeModel = json.data;

      CONSTANT.CUSTOM_FIELD_FOR_PROJECT = json.customFields;
      this.completeLoading();
    });

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

  rename(event: any) {
    let testCase = event.data;
    this._caseInRunService.rename(this.projectId, this.runId, testCase).subscribe((json:any) => {
      event.deferred.resolve(json.data);
    });
  }
  delete(event: any) {
    let testCase = event.data;
    this._caseInRunService.delete(testCase.id, testCase.entityId).subscribe((json:any) => {
      event.deferred.resolve(json.data);
    });
  }
  move(event: any) {
    this._caseInRunService.move(this.projectId, this.runId, event.data).subscribe((json:any) => {
      event.deferred.resolve(json.data);
    });
  }

}

