import {Component, ViewEncapsulation, OnInit, AfterViewInit, ViewChild, ElementRef} from "@angular/core";
import { FormBuilder, FormGroup } from '@angular/forms';

import 'rxjs/add/operator/debounceTime';
import 'rxjs/add/operator/map';

import {GlobalState} from "../../../../global.state";

import {CONSTANT} from "../../../../utils/constant";
import {Utils} from "../../../../utils/utils";
import {RouteService} from "../../../../service/route";
import {ProjectService} from "../../../../service/project";

@Component({
  selector: 'project-list',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./list.scss', '../../../../components/table-tree/src/styles.scss'],
  templateUrl: './list.html'
})
export class ProjectList implements OnInit, AfterViewInit {

  @ViewChild('#tree')tree :ElementRef;

  queryForm: FormGroup;
  queryModel:any = {keywords: '', disabled: 'false'};

  models: any;
  maxLevel: number;
  statusMap: Array<any> = CONSTANT.EntityDisabled;

  constructor(private _routeService:RouteService, private _state:GlobalState, private fb: FormBuilder, private el: ElementRef,
              private _projectService:ProjectService) {

    this.queryForm = this.fb.group(
      {
        'disabled': ['', []],
        'keywords': ['', []]
      }, {}
    );
    this.queryForm.valueChanges.debounceTime(CONSTANT.DebounceTime).subscribe(values => this.queryChange(values));

    this._state.subscribe('my.orgs.change', (data: any) => {
      console.log(11);
      this.loadData();
    });
  }

  ngOnInit() {
    this.loadData();
  }

  ngAfterViewInit() {

  }

  create(type: string):void {
    let that = this;

    that._routeService.navTo('/pages/project/null/edit/' + type);
  }

  queryChange(values:any):void {
    let that = this;

    that.loadData();
  }
  pageChange(event:any):void {
    this.loadData();
  }

  edit($event: any):void {
    let that = this;
  }
  delete($event: any):void {
    let that = this;

  }

  loadData() {
    let that = this;

    that._projectService.list(that.queryModel).subscribe((json:any) => {
      that.models = json.data;
    });
  }

}
