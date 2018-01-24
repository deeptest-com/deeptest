import {Component, ViewEncapsulation, NgModule, Pipe, Compiler, OnInit, AfterViewInit, ViewChild} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';

import {NgbDatepickerI18n, NgbDateParserFormatter, NgbDateStruct, NgbModal, NgbModalRef, ModalDismissReasons} from '@ng-bootstrap/ng-bootstrap';

import {I18n, CustomDatepickerI18n} from '../../../../service/datepicker-I18n';

import {GlobalState} from '../../../../global.state';

import { CONSTANT } from '../../../../utils/constant';
import { RouteService } from '../../../../service/route';

import { PlanService } from '../../../../service/plan';
import { RunService } from '../../../../service/run';
import { ReportService } from '../../../../service/report';

declare var jQuery;

@Component({
  selector: 'plan-view',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./view.scss'],
  templateUrl: './view.html',
  providers: [I18n, {provide: NgbDatepickerI18n, useClass: CustomDatepickerI18n}]
})
export class PlanView implements OnInit, AfterViewInit {
  orgId: number;
  prjId: number;

  planId: number;
  model: any = {};
  form: any;

  testSet: any;
  modalTitle: string;

  chartData: any = {};

  constructor(private _routeService: RouteService, private _route: ActivatedRoute,
              private _reportService: ReportService,
              private _planService: PlanService, private _runService: RunService) {

  }
  ngOnInit() {
    this.orgId = CONSTANT.CURR_ORG_ID;
    this.prjId = CONSTANT.CURR_PRJ_ID;

    this._route.params.forEach((params: Params) => {
      this.planId = +params['planId'];
    });

    if (this.planId) {
      this.loadData();
      this._reportService.planReport(this.planId).subscribe((json:any) => {
        this.chartData = json.data;
      });
    }
  }
  ngAfterViewInit() {}

  loadData() {
    let that = this;
    that._planService.get(that.planId).subscribe((json:any) => {
      that.model = json.data;
    });
  }

  exeOrView(runId: number) {
    this._routeService.navTo('/pages/org/' + CONSTANT.CURR_ORG_ID + '/prj/' + CONSTANT.CURR_PRJ_ID + '/implement/plan/' + this.planId + '/execution/' + runId);
  }
  close(runId: number, index: number) {
    this._runService.close(runId).subscribe((json:any) => {
      if (json.code == 1) {
        this.model.runVos[index] = json.data;
      }
    });
  }

  returnTo() {
    let url: string = '/pages/org/' + CONSTANT.CURR_ORG_ID + '/prj/' + CONSTANT.CURR_PRJ_ID + '/implement/plan/list';
    this._routeService.navTo(url);
  }

}

