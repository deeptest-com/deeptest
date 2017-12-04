import {Component, ViewEncapsulation, ViewChild} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import { NgModule, Pipe, OnInit, AfterViewInit }      from '@angular/core';

import * as _ from 'lodash';

import {GlobalState} from '../../../../global.state';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import {ValidatorUtils, PhoneValidator} from '../../../../validator';
import { RouteService } from '../../../../service/route';

import { UserService } from '../../../../service/user';
import { PopDialogComponent } from '../../../../components/pop-dialog'

declare var jQuery;

@Component({
  selector: 'user-invite',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./invite.scss'],
  templateUrl: './invite.html'
})
export class UserInvite implements OnInit, AfterViewInit {

  id: number;
  tab: string = 'info';

  user: any = {};
  relations: any[] = [];
  form: FormGroup;
  @ViewChild('modalWrapper') modalWrapper: PopDialogComponent;

  constructor(private _state:GlobalState, private _routeService: RouteService, private _route: ActivatedRoute,
              private fb: FormBuilder, private userService: UserService) {

  }
  ngOnInit() {
    this.loadData();
    this.buildForm();
  }
  ngAfterViewInit() {}

  tabChange(event: any) {
    this.tab = event.nextId;
  }

  buildForm(): void {
    this.form = this.fb.group(
      {
        'name': ['', [Validators.required]],
        'email': ['', [Validators.required, Validators.email]],
        'groups': ['', []]
      }, {}
    );

    this.form.valueChanges.debounceTime(CONSTANT.DebounceTime).subscribe(data => this.onValueChanged(data));
    this.onValueChanged();
  }
  onValueChanged(data?: any) {
    let that = this;
    that.formErrors = ValidatorUtils.genMsg(that.form, that.validateMsg, []);
  }

  formErrors = [];
  validateMsg = {
    'name': {
      'required':      '姓名不能为空'
    },
    'email': {
      'required':      '邮箱不能为空',
      'email':      '邮箱格式错误'
    }
  };

  loadData() {
    let that = this;
    that.userService.get(null).subscribe((json:any) => {
      that.relations = json.relations;

      _.forEach(that.relations, (group: any, index: number) => {
        this.form.addControl('group-' + group.orgGroupId, new FormControl('', []))
      });
    });
  }

  invite() {
    let that = this;

    that.userService.invite(that.user, that.relations).subscribe((json:any) => {
      if (json.code == 1) {

        that.formErrors = ['保存成功'];
        that._routeService.navTo("/pages/org-admin/user/list");

      } else {
        that.formErrors = [json.msg];
      }
    });
  }

  select(key: string) {
    let val = key ==='all'? true: false;
    for (let group of this.relations) {
      group.selecting = val;
    }
  }

}

