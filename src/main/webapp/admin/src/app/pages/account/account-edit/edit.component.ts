import { Component,ViewEncapsulation, Pipe, OnInit, AfterViewInit, ViewChild } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import {FormGroup, AbstractControl, FormBuilder, Validators} from '@angular/forms';
import {EmailValidator} from '../../../validator';

import { DropdownModule} from 'ng2-bootstrap/ng2-bootstrap';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';
import {ValidatorUtils} from '../../../validator/validator.utils';

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
  public model:{};

  public submitted:boolean = false;
  public errors: string;

  constructor(private _routeService: RouteService, private _route: ActivatedRoute, private fb: FormBuilder,
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
