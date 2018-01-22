import { Component,ViewEncapsulation, Pipe, OnInit, AfterViewInit, ViewChild, Output, EventEmitter } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Router, ActivatedRoute, Params } from '@angular/router';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import {ValidatorUtils} from '../../../../validator/validator.utils';
import {PasswordsEqualValidator} from '../../../../validator';

import { RouteService } from '../../../../service/route';

import {AccountService} from './../../../../service/account';

declare var jQuery;

@Component({
  selector: 'password-edit',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./edit.scss'],

  templateUrl: './edit.html'
})
export class PasswordEditComponent implements OnInit, AfterViewInit {
  model: any = {};
  form: any;

  @Output() saveEmitter = new EventEmitter<any>();
  @Output() dismissEmitter = new EventEmitter<any>();

  constructor(private _routeService: RouteService, private _route: ActivatedRoute, private fb: FormBuilder,
              private accountService: AccountService) {

    let that = this;
  }

  ngOnInit() {
    this.buildForm();
    this.loadData();
  }

  ngAfterViewInit() {
    let that = this;

  }

  onSubmit():void {
    let that = this;

    that.accountService.changePassword(that.model).subscribe((json:any) => {
        if (json.code == 1) {
          that.loadData();
          that.formErrors = ['修改密码成功'];
          this.saveEmitter.emit({});
        } else {
          that.formErrors = ['修改密码失败'];
        }
    });
  }
  cancel(): any {
    this.dismissEmitter.emit({});
  }

  loadData() {

  }

  buildForm(): void {
    this.form = this.fb.group(
      {
        'oldPassword': ['', [Validators.required]],
        'password': ['', [Validators.required, Validators.pattern(/^[0-9a-zA-Z]{6,10}$/)]],
        'rePassword': ['', [Validators.required, Validators.pattern(/^[0-9a-zA-Z]{6,10}$/)]]
      }, {validator: PasswordsEqualValidator.validate('passwordsEqual', 'password', 'rePassword')}
    );

    this.form.valueChanges.debounceTime(CONSTANT.DebounceTime).subscribe(data => this.onValueChanged(data));
    this.onValueChanged();
  }
  onValueChanged(data?: any) {
    let that = this;
    that.formErrors = ValidatorUtils.genMsg(that.form, that.validateMsg, ['passwordsEqual']);
  }

  formErrors = [];
  validateMsg = {
    'password': {
      'required':      '密码不能为空',
      'pattern': '密码必须为6到10位的字母和数字组合'
    },
    'rePassword': {
      'required':      '重复密码不能为空',
      'pattern': '重复密码必须为6到10位的字母和数字组合'
    },
    'passwordsEqual': '两次密码必须相同'
  };

}
