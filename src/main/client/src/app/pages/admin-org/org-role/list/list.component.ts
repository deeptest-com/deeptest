import {Component, ViewEncapsulation, OnInit, AfterViewInit, ViewChild, ElementRef} from "@angular/core";
import { FormBuilder, FormGroup } from '@angular/forms';
import {GlobalState} from "../../../../global.state";

import {CONSTANT} from "../../../../utils/constant";
import {Utils} from "../../../../utils/utils";
import {RouteService} from "../../../../service/route";
import {OrgRoleService} from "../../../../service/org-role";

@Component({
  selector: 'role-list',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./list.scss', '../../../../components/table-tree/src/styles.scss'],
  templateUrl: './list.html'
})
export class OrgRoleList implements OnInit, AfterViewInit {

  queryForm: FormGroup;
  queryModel:any = {keywords: '', disabled: 'false'};
  statusMap: Array<any> = CONSTANT.EntityDisabled;

  models: any;
  collectionSize:number = 0;
  page:number = 1;
  pageSize:number = 6;

  constructor(private _routeService:RouteService, private _state:GlobalState, private fb: FormBuilder, private el: ElementRef,
              private orgRoleService: OrgRoleService) {
  }

  ngOnInit() {
    let that = this;

    that.queryForm = that.fb.group(
      {
        'disabled': ['', []],
        'keywords': ['', []]
      }, {}
    );

    that.loadData();
  }

  ngAfterViewInit() {
    let that = this;

    this.queryForm.valueChanges.debounceTime(CONSTANT.DebounceTime).subscribe(values => this.queryChange(values));
  }

  create():void {
    let that = this;

    that._routeService.navTo("/pages/org-admin/org-role/edit/null");
  }

  queryChange(values:any):void {
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

    that.orgRoleService.list(that.queryModel).subscribe((json:any) => {
      that.models = json.data;
    });
  }

}
