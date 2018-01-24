import {Component, ViewEncapsulation, ViewChild, QueryList, Query} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { NgModule, Pipe, OnInit, AfterViewInit }      from '@angular/core';

import { CONSTANT } from '../../../../utils/constant';
import { ProjectService } from '../../../../service/project';
import { ReportService } from '../../../../service/report';

declare var jQuery;

@Component({
  selector: 'project-view',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./view.scss'],
  templateUrl: './view.html'
})
export class ProjectView implements OnInit, AfterViewInit {
  orgId: number;
  id: number;

  project: any = {};
  plans: any[] = [];
  histories: any = {};

  chartData: any = {};

  constructor(private _route: ActivatedRoute,
              private _projectService: ProjectService, private _reportService: ReportService) {

  }
  ngOnInit() {
    this.orgId = CONSTANT.CURR_ORG_ID;

    this._route.params.subscribe(params => {
      this.id = +params['id'];
    });
    if (this.id) {
      this.loadData();
      this._reportService.projectReport(this.id).subscribe((json:any) => {
        this.chartData = json.data;
      });
    }
  }
  ngAfterViewInit() {

  }

  loadData() {
      CONSTANT.CURR_PRJ_ID = this.id;
      this._projectService.view(this.id).subscribe((json:any) => {
        this.project = json.project;
        this.plans = json.plans;
        this.histories = json.histories;
      });
  }

}

