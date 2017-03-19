import {Component, ViewEncapsulation, OnInit, AfterViewInit, ViewChild} from "@angular/core";
import { FormBuilder, FormGroup } from '@angular/forms';
import {GlobalState} from "../../../../global.state";

import {CONSTANT} from "../../../../utils/constant";
import {Utils} from "../../../../utils/utils";
import {RouteService} from "../../../../service/route";
import {ProjectService} from "../../../../service/project";

@Component({
  selector: 'project-list',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./list.scss'), require('../../../../components/table-tree/src/styles.scss')],
  template: require('./list.html')
})
export class ProjectList implements OnInit, AfterViewInit {

  queryForm: FormGroup;
  queryModel:any = {keywords: '', status: ''};

  models: any;
  maxLevel: number;
  counter = Array;
  statusMap: Array<any> = CONSTANT.EntityActive;

  constructor(private _routeService:RouteService, private _state:GlobalState, private fb: FormBuilder,
              private _projectService:ProjectService) {
  }

  ngOnInit() {
    let that = this;

    that.queryForm = that.fb.group(
      {
        'status': [that.queryModel.status, []],
        'keywords': [that.queryModel.keywords, []]
      }, {}
    );

    that.loadData();
  }

  ngAfterViewInit() {
    let that = this;

    this.queryForm.valueChanges.debounceTime(500).subscribe(values => this.queryChange(values));
  }

  create():void {
    let that = this;

    that._routeService.navTo("/pages/event/edit/null/property");
  }

  queryChange(values:any):void {
    let that = this;

    that.queryModel = values;
    console.log(that.queryModel);

    that.loadData();
  }
  pageChanged(event:any):void {
    let that = this;
    that.loadData();
  }

  edit($event: any):void {
    let that = this;

    console.log($event);
  }
  delete($event: any):void {
    let that = this;

    console.log($event);
  }

  loadData() {
    let that = this;
    that._projectService.list(that.queryModel).subscribe((json:any) => {
      console.log('json', json);
      that.models = json.data;
      that.maxLevel = json.maxLevel;
    });
  }

}

