import {Component, ViewEncapsulation, OnInit, AfterViewInit} from "@angular/core";
import { Router, ActivatedRoute, Params } from '@angular/router';

import {GlobalState} from "../../../../global.state";
import {CONSTANT} from "../../../../utils/constant";
import {Utils} from "../../../../utils/utils";
import {RouteService} from "../../../../service/route";
import {SlimLoadingBarService} from "../../../../components/ng2-loading-bar";
import {CaseService} from "../../../../service/case";

declare var jQuery;

@Component({
  selector: 'case-suite',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./suite.scss',
    '../../../../../vendor/ztree/css/zTreeStyle/zTreeStyle.css',
    '../../../../components/ztree/src/styles.scss'],
  templateUrl: './suite.html'
})
export class CaseSuite implements OnInit, AfterViewInit {

  projectId: number;
  public treeModel: any;
  public treeSettings: any = {usage: 'edit', isExpanded: true, sonSign: false};

  constructor(private _routeService:RouteService, private _route: ActivatedRoute, private _state:GlobalState,
              private _caseService:CaseService,
              private slimLoadingBarService:SlimLoadingBarService) {

  }

  ngOnInit() {
    this._route.params.forEach((params: Params) => {
      this.projectId = +params['projectId'];
    });

    this.loadData();
  }

  ngAfterViewInit() {

  }

  delete(id:string):void {

  }

  loadData() {
    this.startLoading();

    this._caseService.query(this.projectId).subscribe((json:any) => {
      this.treeModel = json.data;
      CONSTANT.CUSTOM_FIELD_FOR_PROJECT = json.customFields;

      this.completeLoading();
    });

  }

  rename(event: any) {
    let testCase = event.data;
    this._caseService.rename(this.projectId, testCase).subscribe((json:any) => {
      event.deferred.resolve(json.data);
    });
  }
  remove(event: any) {
    let testCase = event.data;
    this._caseService.delete(testCase.id).subscribe((json:any) => {
      event.deferred.resolve(json.data);
    });
  }
  move(event: any) {
    this._caseService.move(this.projectId, event.data).subscribe((json:any) => {
      event.deferred.resolve(json.data);
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

}

