import {Component, ViewEncapsulation, ViewChild} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
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
  selector: 'user-edit',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./edit.scss')],
  template: require('./edit.html')
})
export class UserEdit implements OnInit, AfterViewInit {

  id: number;
  tab: string = 'info';

  user: any = {disabled: false};
  groups: any[] = [];
  form: FormGroup;
  isSubmitted: boolean;
  @ViewChild('modal') modal: ModalDirective;

  constructor(private _state:GlobalState, private _routeService: RouteService, private _route: ActivatedRoute,
              private fb: FormBuilder, private userService: UserService) {

  }
  ngOnInit() {
    this._route.params.forEach((params: Params) => {
      this.id = +params['id'];
    });

    this.loadData();
    this.buildForm();
  }
  ngAfterViewInit() {}


  selectTab(tab: string) {
    let that = this;
    that.tab = tab;
  }

  buildForm(): void {
    let that = this;
    this.form = this.fb.group(
      {
        'name': ['', [Validators.required]],
        'email': ['', [Validators.required, EmailValidator.validate()]],
        'phone': ['', [Validators.required, PhoneValidator.validate()]],
        'disabled': ['', []],
        'groups': ['', []]
      }, {}
    );

    this.form.valueChanges.subscribe(data => this.onValueChanged(data));
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
      'validate':      '邮箱格式不正确'
    },
    'phone': {
      'required':      '手机不能为空',
      'validate':      '手机格式不正确'
    }
  };

  loadData() {
    let that = this;
    that.userService.get(that.id).subscribe((json:any) => {
      that.user = json.user;
      that.groups = json.groups;

      _.forEach(that.groups, (group: any, index: number) => {
        this.form.addControl('group-' + group.id, new FormControl('', []))
      });
    });
  }

  save() {
    let that = this;

    that.userService.save(that.user, that.groups).subscribe((json:any) => {
      if (json.code == 1) {

        that.formErrors = ['保存成功'];
        that._routeService.navTo("/pages/org-admin/user/list");
      } else {
        that.formErrors = ['保存失败'];
      }
    });
  }

  delete() {
    let that = this;

    that.userService.delete(that.user.id).subscribe((json:any) => {
      if (json.code == 1) {
        that.formErrors = ['删除成功'];
        that._routeService.navTo("/pages/org-admin/user/list");
      } else {
        that.formErrors = ['删除失败'];
      }
    });
  }

  select(key: string) {
    let val = key ==='all'? true: false;
    for (let group of this.groups) {
      group.selecting = val;
    }
  }
  reset() {
    this.loadData();
  }

  showModal(): void {
    this.modal.show();
  }
  onModalShow():void {
    // init jquery components if needed
  }

  hideModal(): void {
    this.modal.hide();
  }

}

