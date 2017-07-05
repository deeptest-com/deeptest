import {Component, ViewEncapsulation, OnInit, AfterViewInit} from "@angular/core";

import {GlobalState} from "../../../../global.state";
import {CONSTANT} from "../../../../utils/constant";
import {Utils} from "../../../../utils/utils";
import {RouteService} from "../../../../service/route";
import {SlimLoadingBarService} from "../../../../components/ng2-loading-bar";
import {TreeService} from "../../../../components/ng2-tree/src/tree.service";
import {PlanService} from "../../../../service/plan";

@Component({
  selector: 'plan-list',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./list.scss'],
  templateUrl: './list.html'
})
export class PlanList implements OnInit, AfterViewInit {
  models:any[];
  query:any = {keywords: '', status: ''};

  constructor(private _routeService:RouteService, private _state:GlobalState,
              private _planService:PlanService) {
  }

  ngOnInit() {
    let that = this;
    that.loadData();
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
    that._planService.query(that.query).subscribe((json:any) => {
      that.models = json.data;
    });
  }

}

