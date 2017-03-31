import {Component, ViewEncapsulation, OnInit, AfterViewInit, ViewChild, ElementRef} from "@angular/core";
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

  @ViewChild('#tree')tree :ElementRef;

  queryForm: FormGroup;
  queryModel:any = {keywords: '', disabled: 'false'};

  models: any;
  maxLevel: number;
  counter = Array;
  statusMap: Array<any> = CONSTANT.EntityDisabled;

  constructor(private _routeService:RouteService, private _state:GlobalState, private fb: FormBuilder, private el: ElementRef,
              private _projectService:ProjectService) {
  }

  ngOnInit() {
    let that = this;

    that.queryForm = that.fb.group(
      {
        'disabled': [],
        'keywords': []
      }, {}
    );

    that.loadData();
  }

  ngAfterViewInit() {
    let that = this;

    this.queryForm.valueChanges.debounceTime(CONSTANT.DebounceTime).subscribe(values => this.queryChange(values));
  }

  create(type: string):void {
    let that = this;

    that._routeService.navTo('/pages/project/edit/' + type + '/null');
  }

  queryChange(values:any):void {
    let that = this;

    that.loadData();
  }
  pageChanged(event:any):void {
    let that = this;
    that.loadData();
  }

  edit($event: any):void {
    let that = this;
  }
  delete($event: any):void {
    let that = this;

    console.log($event);
  }

  loadData() {
    let that = this;

    that._projectService.list(that.queryModel).subscribe((json:any) => {
      that.models = json.data;
    });
  }

}
