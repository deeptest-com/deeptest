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
  models: any;

  queryForm: FormGroup;
  queryModel:any = {keywords: '', disabled: 'false'};
  statusMap: Array<any> = CONSTANT.ExeStatus;

  constructor(private _routeService:RouteService, private _route: ActivatedRoute,
              private _state:GlobalState, private fb: FormBuilder, private el: ElementRef,
              private _planService:PlanService) {

  }

  ngOnInit() {
    this.loadData();

    this.queryForm = this.fb.group(
      {
        'disabled': ['', []],
        'keywords': ['', []]
      }, {}
    );

    this.loadData();
  }

  ngAfterViewInit() {
    let that = this;
  }

  create():void {
    let that = this;

    that._routeService.navTo("/pages/implement/" + CONSTANT.CURRENT_PROJECT.id + "/plan/null/edit");
  }

  delete(eventId:string):void {

  }

  loadData() {
    this._planService.query(CONSTANT.CURRENT_PROJECT.id, this.queryModel).subscribe((json:any) => {
      this.models = json.data;
    });
  }

}

