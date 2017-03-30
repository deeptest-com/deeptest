import {Component, ViewEncapsulation, ViewChild} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { NgModule, Pipe, OnInit, AfterViewInit }      from '@angular/core';
import { ModalDirective } from 'ng2-bootstrap';
import {GlobalState} from '../../../../global.state';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import {ValidatorUtils, EmailValidator, PhoneValidator} from '../../../../validator';
import { RouteService } from '../../../../service/route';

import { UserService } from '../../../../service/user';

declare var jQuery;

@Component({
  selector: 'user-edit-groups',
  encapsulation: ViewEncapsulation.None,
  styles: [],
  template: require('./edit-groups.html')
})
export class UserEditGroups implements OnInit, AfterViewInit {
  id: number;
  models: any = {};
  modelsClone: any = {};
  form: any;

  constructor(private _state:GlobalState, private _routeService: RouteService, private _route: ActivatedRoute,
              private fb: FormBuilder, private userService: UserService) {

  }
  ngOnInit() {
    let that = this;

    that._route.params.forEach((params: Params) => {
      that.id = +params['id'];
    });

    if (that.id) {
      that.loadData();
    }
    that.buildForm();
  }
  ngAfterViewInit() {}

  buildForm(): void {
    let that = this;
    this.form = this.fb.group(
      {
        'groups': [null, []]
      }, {}
    );
  }

  formErrors = [];

  loadData() {
    let that = this;
    // that.userService.get(that.id).subscribe((json:any) => {
    //   that.model = json.data;
    // });

    that.models = [{name: '测试人员', id: 1}, {name: '开发人员', id: 2}];
  }

  save() {
    let that = this;

    that.userService.saveGroups(that.models).subscribe((json:any) => {
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
      model.checked = val;
    }
  }
  reset() {
    this.loadData();
  }
}

