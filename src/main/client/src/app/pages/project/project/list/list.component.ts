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

    this._state.subscribe(CONSTANT.STATE_CHANGE_ORGS, (data: any) => {
      console.log('=STATE_CHANGE_ORGS=');
      this.loadData();
    });
  }

  ngOnInit() {
  }

  ngAfterViewInit() {

  }

  create(type: string):void {
    this._routeService.navTo('/pages/org/' + CONSTANT.CURR_ORG_ID + '/prjs/null/edit/' + type);
  }

  queryChange(values:any):void {
    console.log('=queryChange=');
    this.loadData();
  }
  loadData() {
    console.log('=loadData=');
    this._projectService.list(this.queryModel).subscribe((json:any) => {
      this.models = json.data;
    });
  }

}
