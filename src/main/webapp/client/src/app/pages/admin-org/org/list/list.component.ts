import {Component, ViewEncapsulation, OnInit, AfterViewInit, ViewChild, ElementRef} from "@angular/core";
import { FormBuilder, FormGroup } from '@angular/forms';

import {GlobalState} from "../../../../global.state";

import {CONSTANT} from "../../../../utils/constant";
import {Utils} from "../../../../utils/utils";
import {RouteService} from "../../../../service/route";
import {OrgService} from "../../../../service/org";
import {AccountService} from "../../../../service/account";

@Component({
  selector: 'org-list',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./list.scss'],
  templateUrl: './list.html'
})
export class OrgList implements OnInit, AfterViewInit {

  queryForm: FormGroup;
  queryModel:any = {keywords: '', disabled: 'false'};
  statusMap: Array<any> = CONSTANT.EntityDisabled;

  models: any;

  constructor(private _routeService:RouteService, private _state:GlobalState, private fb: FormBuilder, private el: ElementRef,
              private orgService:OrgService, private accountService: AccountService) {
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
    this._routeService.navTo('/pages/org-admin/org/edit/null');
  }

  queryChange(values:any):void {
    let that = this;
    that.loadData();
  }

  setDefault(item: any):void {
    this.orgService.setDefault(item.id, this.queryModel).subscribe((json:any) => {
      if (json.code == 1) {

        this.models = json.data;
        this.accountService.changeRecentProject(json.recentProjects);
      }
    });
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

    that.orgService.list(that.queryModel).subscribe((json:any) => {
      that.models = json.data;

      if (that.models.length == 0 && !that.queryModel.keywords && !that.queryModel.disabled) {
        this._state.notifyDataChanged('org.ready', false);
        this._routeService.navTo('/pages/org-admin/org/edit/null');
      } else {
        this._state.notifyDataChanged('org.ready', true);
      }
    });
  }
}
