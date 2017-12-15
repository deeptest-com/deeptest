import {Component, ViewEncapsulation, ViewChild} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import { NgModule, Pipe, OnInit, AfterViewInit }      from '@angular/core';

import * as _ from 'lodash';

import {GlobalState} from '../../../../global.state';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import {ValidatorUtils} from '../../../../validator/validator.utils';
import { RouteService } from '../../../../service/route';

import { OrgRoleService } from '../../../../service/org-role';
import { PopDialogComponent } from '../../../../components/pop-dialog'

declare var jQuery;

@Component({
  selector: 'org-role-edit',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./edit.scss'],
  templateUrl: './edit.html'
})
export class OrgRoleEdit implements OnInit, AfterViewInit {
  id: number;
  tab: string = 'info';
  orgRole: any = {disabled: false};
  orgRolePrivileges: any[] = [];
  orgRoleUsers: any[] = [];

  form: any;
  @ViewChild('modalWrapper') modalWrapper: PopDialogComponent;

  constructor(private _state:GlobalState, private _routeService: RouteService, private _route: ActivatedRoute,
              private fb: FormBuilder, private orgRoleService: OrgRoleService) {

  }
  ngOnInit() {
    let that = this;

    that._route.params.forEach((params: Params) => {
      that.id = +params['id'];
    });

    that.loadData();
    that.buildForm();
  }
  ngAfterViewInit() {}

  loadData() {
    let that = this;
    that.orgRoleService.get(that.id).subscribe((json:any) => {
      that.orgRole = json.orgRole;
      that.orgRolePrivileges = json.orgRolePrivileges;
      that.orgRoleUsers = json.orgRoleUsers;

      _.forEach(that.orgRolePrivileges, (privilege: any, index: number) => {
        this.form.addControl('privilege-' + privilege.id, new FormControl('', []))
      });

      _.forEach(that.orgRoleUsers, (user: any, index: number) => {
        this.form.addControl('user-' + user.id, new FormControl('', []))
      });
    });
  }

  save() {
    let that = this;

    that.orgRoleService.save(that.orgRole, that.orgRolePrivileges, that.orgRoleUsers).subscribe((json:any) => {
      if (json.code == 1) {

        that.formErrors = ['保存成功'];
        that._routeService.navTo("/pages/org-admin/org-role/list");
      } else {
        that.formErrors = ['保存失败'];
      }
    });
  }

  delete() {
    let that = this;

    that.orgRoleService.delete(that.orgRole.id).subscribe((json:any) => {
      if (json.code == 1) {
        that.formErrors = ['删除成功'];
        that._routeService.navTo("/pages/org-admin/org-role/list");

        this.modalWrapper.closeModal();
      } else {
        that.formErrors = [json.msg];
      }
    });
  }

  select(key: string) {
    let val = key ==='all'? true: false;
    for (let user of this.orgRolePrivileges) {
      user.selecting = val;
    }
  }
  tabChange(event: any) {
    this.tab = event.nextId;
  }

  buildForm(): void {
    this.form = this.fb.group(
      {
        'name': ['', [Validators.required]],
        'descr': ['', []],
        'disabled': ['', []]
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
      'required':      '名称不能为空'
    },
    'descr': {
    }
  };

  showModal(): void {
    this.modalWrapper.showModal();
  }

}

