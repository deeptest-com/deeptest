import {Component, ViewEncapsulation, OnInit, AfterViewInit, ViewChild, ElementRef} from "@angular/core";
import { FormBuilder, FormGroup } from '@angular/forms';
import {GlobalState} from "../../../../global.state";

import {CONSTANT} from "../../../../utils/constant";
import {Utils} from "../../../../utils/utils";
import {RouteService} from "../../../../service/route";
import {GroupService} from "../../../../service/group";

@Component({
  selector: 'group-list',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./list.scss')],
  template: require('./list.html')
})
export class GroupList implements OnInit, AfterViewInit {

  queryForm: FormGroup;
  queryModel:any = {keywords: '', isActive: 'true'};

  models: any;
  maxLevel: number;
  counter = Array;
  statusMap: Array<any> = CONSTANT.EntityActive;

  tableData = [
    {
      groupname: '@mdo',
      status: 'info'
    },
    {
      groupname: '@fat',
      status: 'primary'
    },
    {
      groupname: '@twitter',
      status: 'success'
    }
  ];
  roleData = [
    {
      id: 1,
      name: '主管'
    },
    {
      id: 2,
      name: '工程师'
    }
  ];

  constructor(private _routeService:RouteService, private _state:GlobalState, private fb: FormBuilder, private el: ElementRef,
              private groupService:GroupService) {
  }

  ngOnInit() {
    let that = this;

    that.queryForm = that.fb.group(
      {
        'isActive': [that.queryModel.isActive, []],
        'keywords': [that.queryModel.keywords, []]
      }, {}
    );

    // that.loadData();
  }

  ngAfterViewInit() {
    let that = this;

    this.queryForm.controls['isActive'].valueChanges.debounceTime(500).subscribe(values => this.queryChange(values));
  }

  create():void {
    let that = this;

    that._routeService.navTo("/pages/event/edit/null/property");
  }

  queryChange(values:any):void {
    let that = this;

    that.queryModel.isActive = values;

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
  changeRole(id: number):void {
    let that = this;

    console.log('===', id);
  }

  loadData() {
    let that = this;
    console.log(that.queryModel);

    that.groupService.list(that.queryModel).subscribe((json:any) => {
      that.models = json.data.models;
      that.maxLevel = json.data.maxLevel;
    });
  }

}
