import { Component,ViewEncapsulation, Pipe, OnInit, AfterViewInit, ViewChild } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import {FormGroup, AbstractControl, FormBuilder, Validators} from '@angular/forms';
import {EmailValidator} from '../../../../validator';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import {ValidatorUtils} from '../../../../validator/validator.utils';

import { RouteService } from '../../../../service/route';

import { UserService } from '../../../../service/user';
import { AccountService } from '../../../../service/account';

declare var jQuery;

@Component({
  selector: 'account-edit',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./edit.scss')],

  template: require('./edit.html')
})
export class AccountEdit implements OnInit, AfterViewInit {
  public accountId: number;
  public model:any = {};
  public form:FormGroup;

  public errors: string;

  constructor(private _routeService: RouteService, private _route: ActivatedRoute, private fb: FormBuilder,
              private _accountService: AccountService, private _userService: UserService) {

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

    that.initForm();
  }

  onSubmit():void {
    let that = this;

    that._accountService.save(that.model).subscribe((json:any) => {
        if (json.code == 1) {
          that._routeService.navTo("/pages/business/account-list");
        }
    });
  }
  forgotPassword():void {
    let that = this;

    that._userService.forgotPassword(that.model.id).subscribe((json:any) => {
      if (json.code == 1) {
        that.formErrors = ['重置密码成功'];
      }
    });
  }

  loadData() {
   let that = this;

   that._accountService.get(that.accountId).subscribe((json:any) => {
      that.model = json.data;

     that.buildForm();
   });
  }

  back($event: any):void {
    let that = this;
    that._routeService.navTo('/pages/business/account-list');
  }


  initForm() {
    let that = this;
    // init controller
  }
  buildForm(): void {
    let that = this;
    this.form = this.fb.group(
      {
        'name': [that.model['name'], [Validators.required]],
        'phone': [that.model['phone'], [Validators.required]],
        'email': [that.model['email'], [Validators.required, EmailValidator.validate()]]
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
      'required':      '姓名不能为空',
      'minlength': '姓名长度不能少于4位'
    },
    'phone': {
      'required':      '电话不能为空',
      'minlength': '姓名厂部不能少于11位'
    },
    'email': {
      'required':      '邮件不能为空',
      'emailValidator': '邮件格式错误'
    }
  };

}
