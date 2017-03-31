import {Component, ViewEncapsulation, ViewChild} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { NgModule, Pipe, OnInit, AfterViewInit }      from '@angular/core';

import { GroupService } from '../../../../service/group';

declare var jQuery;

@Component({
  selector: 'user-edit-groups',
  encapsulation: ViewEncapsulation.None,
  styles: [],
  template: require('./edit-groups.html')
})
export class UserEditGroups implements OnInit, AfterViewInit {
  id: number;
  models: any[] = [];
  formErrors: any[] = [];

  constructor(private _route: ActivatedRoute, private groupService: GroupService) {

  }
  ngOnInit() {
    let that = this;

    that._route.params.forEach((params: Params) => {
      that.id = +params['id'];
    });

    if (that.id) {
      that.loadData();
    }
  }
  ngAfterViewInit() {}

  loadData() {
    let that = this;
    that.groupService.listByUser(that.id).subscribe((json:any) => {
      that.models = json.data;
    });
  }

  save() {
    let that = this;

    that.groupService.saveByUser(that.models).subscribe((json:any) => {
      if (json.code == 1) {
        that.formErrors = ['保存成功'];
      } else {
        that.formErrors = ['保存失败'];
      }
    });
  }

  select(key: string) {
    let val = key ==='all'? true: false;
    for (let model of this.models) {
      model.selecting = val;
    }
  }
  reset() {
    this.loadData();
  }
}

