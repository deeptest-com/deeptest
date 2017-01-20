import { Component,ViewEncapsulation, Pipe, OnInit, AfterViewInit, ViewChild } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import {FormGroup, AbstractControl, FormBuilder, Validators} from '@angular/forms';
import {EmailValidator, EqualPasswordsValidator} from '../../../validator';

import { DropdownModule} from 'ng2-bootstrap/ng2-bootstrap';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';
import {Validate} from '../../../service/validate';

import { RouteService } from '../../../service/route';

import { AccountService } from '../../../service/account';

declare var jQuery;

@Component({
  selector: 'account-edit',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./edit.scss')],

  template: require('./edit.html')
})
export class AccountEdit implements OnInit, AfterViewInit {
  accountId: number;
  public form:FormGroup;
  public name:AbstractControl;
  public phone:AbstractControl;
  public email:AbstractControl;
  public password:AbstractControl;
  public repeatPassword:AbstractControl;
  public passwords:FormGroup;

  public submitted:boolean = false;
  public errors: string;

  constructor(fb:FormBuilder, private _routeService: RouteService, private _route: ActivatedRoute, private fb: FormBuilder,
              private _accountService: AccountService) {

    let that = this;
  }

  ngOnInit() {
    let that = this;

    that.buildForm();
    this._route.params.forEach((params: Params) => {
      that.accountId = +params['id'];
    });

    if (that.accountId) {
        that.loadData();
    }
  }

  ngAfterViewInit() {
    let that = this;

    that.initForm(false);
  }

  onSubmit():void {
    let that = this;

    that.model.status = undefined;

    that._accountService.save(that.model).subscribe((json:any) => {
        if (json.code = 1) {
          that._routeService.navTo("/pages/account/list");
        }
    });
  }

  goto($account) {
    let that = this;

    that._routeService.navTo('/pages/account/edit/' + that.accountId + '/' + $account.tabModel);
  }
  loadData() {
   let that = this;

   that._accountService.get(that.accountId).subscribe((json:any) => {
      that.model = json.account;

     that.initForm(true);
   });
  }

  initForm(dataLoaded): void {
    let that = this;

  }

  buildForm(): void {
    let that = this;
    this.form = that.fb.group({
      'name': ['', Validators.compose([Validators.required, Validators.minLength(4)])],
      'phone': ['', Validators.compose([Validators.required, Validators.minLength(11)])],
      'email': ['', Validators.compose([Validators.required, EmailValidator.validate])],
      'passwords': that.fb.group({
        'password': ['', Validators.compose([Validators.required, Validators.minLength(4)])],
        'repeatPassword': ['', Validators.compose([Validators.required, Validators.minLength(4)])]
      }, {validator: EqualPasswordsValidator.validate('password', 'repeatPassword')})
    });

    this.name = this.form.controls['name'];
    this.phone = this.form.controls['phone'];
    this.email = this.form.controls['email'];
    this.passwords = <FormGroup> this.form.controls['passwords'];
    this.password = this.passwords.controls['password'];
    this.repeatPassword = this.passwords.controls['repeatPassword'];
  }
  onValueChanged(data?: any) {
    let that = this;
    if (!that.accountForm) { return; }

    that.formErrors = Validate.genValidateInfo(that.accountForm, that.validateMsg, ['accountTimeCompare', 'registerTimeCompare']);
  }

  formErrors = [];
  validateMsg = {
    'title': {
      'required':      '会议名称不能为空'
    },
    'startDate': {
      'required':      '开始日期不能为空',
      'dateValidator': '开始日期格式不正确'
    },
    'startTime': {
      'required':      '开始时间不能为空',
      'timeValidator': '开始日期格式不正确'
    },
    'endDate': {
      'required':      '结束日期不能为空'
    },
    'endTime': {
      'required':      '结束时间不能为空',
      'timeValidator': '结束时间格式不正确'
    },
    'registerStartDate': {
      'required':      '注册开始日期不能为空',
      'dateValidator': '注册开始日期格式不正确'
    },
    'registerStartTime': {
      'required':      '注册开始时间不能为空',
      'timeValidator': '注册开始时间格式不正确'
    },
    'registerEndDate': {
      'required':      '注册结束日期不能为空',
      'dateValidator': '注册结束日期格式不正确'
    },
    'registerEndTime': {
      'required':      '注册结束时间不能为空',
      'timeValidator':  '注册结束时间格式不正确',
    },
    'signBefore': {
      'required':      '提前注册天数不能为空'
    },
    'address': {
      'required':      '地址不能为空'
    },
    'phone': {
      'required':      '电话不能为空'
    },
    'email': {
      'required':      '邮件不能为空',
      'emailValidator': '邮件格式错误'
    },
    'website': {'required':      '邮件不能为空'},

    'accountTimeCompare': '会议结束时间不能早于开始时间',
    'registerTimeCompare': '报名结束时间不能早于开始时间'
  };

}
