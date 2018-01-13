import {Component, ViewEncapsulation, OnInit, AfterViewInit, ViewChild, ElementRef} from "@angular/core";
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router, ActivatedRoute, Params } from '@angular/router';

import {GlobalState} from "../../../../global.state";
import {CONSTANT} from "../../../../utils/constant";
import {Utils} from "../../../../utils/utils";
import {RouteService} from "../../../../service/route";
import {PlanService} from "../../../../service/plan";

@Component({
  selector: 'plan-list',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./list.scss'],
  templateUrl: './list.html'
})
export class PlanList implements OnInit, AfterViewInit {
  orgId: number;
  prjId: number;
  isInit: boolean;

  models: any;

  queryForm: FormGroup;
  queryModel:any = {keywords: '', status: ''};
  statusMap: Array<any> = CONSTANT.ExeStatus;

  constructor(private _routeService:RouteService, private fb: FormBuilder, private el: ElementRef,
              private _planService:PlanService) {

    this.queryForm = this.fb.group(
      {
        'status': ['', []],
        'keywords': ['', []]
      }, {}
    );


  }

  ngOnInit() {
    this.orgId = CONSTANT.CURR_ORG_ID;
    this.prjId = CONSTANT.CURR_PRJ_ID;

    this.loadData();
  }

  ngAfterViewInit() {
    this.queryForm.valueChanges.debounceTime(CONSTANT.DebounceTime).subscribe(values => this.queryChange(values));
  }

  create():void {
    let that = this;

    that._routeService.navTo('/pages/org/' + CONSTANT.CURR_ORG_ID + '/prj/' + CONSTANT.CURR_PRJ_ID
        + '/implement/plan/null/edit');
  }

  delete(projectId: string):void {

  }

  loadData() {
    this._planService.query(CONSTANT.CURR_PRJ_ID, this.queryModel).subscribe((json:any) => {
      this.models = json.data;
    });
  }

  queryChange(values:any):void {
      this.loadData();
  }

}

