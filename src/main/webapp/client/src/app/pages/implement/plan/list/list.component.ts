import {Component, ViewEncapsulation, OnInit, AfterViewInit, ViewChild, ElementRef} from "@angular/core";
import { FormBuilder, FormGroup } from '@angular/forms';

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
  queryForm: FormGroup;
  queryModel:any = {keywords: '', disabled: 'false'};
  statusMap: Array<any> = CONSTANT.ExeStatus;

  models: any;

  constructor(private _routeService:RouteService, private _state:GlobalState, private fb: FormBuilder, private el: ElementRef,
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

    that._routeService.navTo("/pages/implement/plan/edit/null");
  }

  delete(eventId:string):void {
    console.log('id=' + eventId);
  }

  loadData() {
    let that = this;
    that._planService.query(that.queryModel).subscribe((json:any) => {
      that.models = json.data;
    });
  }

}

