import {Component, ViewEncapsulation, OnInit, AfterViewInit, ViewChild, ElementRef} from "@angular/core";
import { FormBuilder, FormGroup } from '@angular/forms';

import 'rxjs/add/operator/debounceTime';
import 'rxjs/add/operator/map';

import {GlobalState} from "../../../../global.state";

import {CONSTANT} from "../../../../utils/constant";
import {Utils} from "../../../../utils/utils";
import {RouteService} from "../../../../service/route";
import {ReportService} from "../../../../service/report";

@Component({
  selector: 'report-list',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./list.scss'],
  templateUrl: './list.html'
})
export class ReportList implements OnInit, AfterViewInit {
  queryForm: FormGroup;
  queryModel:any = {keywords: '', status: ''};
  statusMap: any = {'': '所有', 'Pass': '通过', 'Fail': '失败'};

  pageData: any = {};
  pageNumb:number = 0;
  pageSize:number = 6;

  constructor(private _routeService:RouteService, private _state:GlobalState, private fb: FormBuilder,
              private el: ElementRef, private reportService:ReportService) {

    this.queryForm = this.fb.group(
      {
        'status': ['', []],
        'keywords': ['', []]
      }, {}
    );
  }

  ngOnInit() {
    let that = this;

    that.loadData();
  }

  ngAfterViewInit() {
    this.queryForm.valueChanges.debounceTime(CONSTANT.DebounceTime).subscribe(values => this.queryChange(values));
  }

  view(item: any):void {
    let that = this;

    that._routeService.navTo("/pages/report/" + item.id);
  }

  queryChange(values: any):void {
    this.loadData();
  }
  pageChange($event:any){
    this.loadData();
  }

  loadData() {
    let that = this;

    that.reportService.list(that.queryModel, that.pageNumb, that.pageSize).subscribe((json:any) => {
      that.pageData = json.data;
    });
  }

}
