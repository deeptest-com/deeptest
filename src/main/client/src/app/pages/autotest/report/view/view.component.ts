import {Component, ViewEncapsulation, ViewChild} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { NgModule, Pipe, OnInit, AfterViewInit }      from '@angular/core';

import {NgbModal, ModalDismissReasons} from '@ng-bootstrap/ng-bootstrap';

import {GlobalState} from '../../../../global.state';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import { RouteService } from '../../../../service/route';

import { ReportService } from '../../../../service/report';

@Component({
  selector: 'report-view',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./view.scss'],
  templateUrl: './view.html'
})
export class ReportView implements OnInit, AfterViewInit {

  id:number;
  model:any;
  checkPointContent:any = {};
  logContent:any = {};

  contentCase: any;
  contentType: string;

  constructor(private _state:GlobalState, private _routeService:RouteService, private _route:ActivatedRoute,
              private modalService: NgbModal, private reportService:ReportService) {

  }

  ngOnInit() {
    this._route.params.forEach((params:Params) => {
      this.id = +params['id'];
    });

    this.loadData();
  }

  ngAfterViewInit() {
  }

  loadData() {
    let that = this;
    that.reportService.get(that.id).subscribe((json:any) => {
      that.model = json.data;
      that.genContent();
    });
  }

  pop(content:any, contentCase: any, contentType: string) {
    let that = this;

    that.contentCase = contentCase;
    that.contentType = contentType;

    this.modalService.open(content, { windowClass: 'large-width-modal' });
  }

  genContent() {
    let that = this;

    for (let suite of that.model.result.suites) {
      for (let cas of suite.cases) {
        if (!that.checkPointContent[cas.id]) {
          that.checkPointContent[cas.id] = '';
        }
        for (let checkpoint of cas.checkPoints) {
          that.checkPointContent[cas.id] +=
            '<div>' + checkpoint.name + ' - ' + checkpoint.status + '</div>';
        }

        if (!that.logContent[cas.id]) {
          that.logContent[cas.id] = '';
        }
        for (let line of cas.output) {
          that.logContent[cas.id] += '<div>' + line + '</div>';
        }
      }
    }
  }

}

