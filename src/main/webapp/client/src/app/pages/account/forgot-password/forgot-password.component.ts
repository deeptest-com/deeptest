import { Component,ViewEncapsulation, Pipe, OnInit, AfterViewInit, ViewChild } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Router, ActivatedRoute, Params } from '@angular/router';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';
import {ValidatorUtils, EmailValidator} from '../../../validator';

import { RouteService } from '../../../service/route';

import { AccountService } from '../../../service/account';

declare var jQuery;

@Component({
  selector: 'forgot-password',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./forgot-password.scss'],
  templateUrl: './forgot-password.html'
})
export class ForgotPassword implements OnInit, AfterViewInit {
  model: any = {};
  form: any;
  public errors: string;

  constructor(private _routeService: RouteService, private _route: ActivatedRoute,
              private fb: FormBuilder, private accountService: AccountService) {

  }

  ngOnInit() {
    this.buildForm();
  }

  ngAfterViewInit() {

  }

  onSubmit():void {
    this.accountService.forgotPassword(this.model.email).subscribe((json: any) => {
      if (json.code == 1) {
        this.formErrors = ['操作成功，请访问您的邮箱设置新密码！'];
      }  else {
        this.formErrors = [json.msg];
      }
    });
  }

  buildForm(): void {
    this.form = this.fb.group(
      {
        'email': ['', [Validators.required, EmailValidator.validate()]],
      }
    );

    this.form.valueChanges.debounceTime(CONSTANT.DebounceTime).subscribe(data => this.onValueChanged(data));
    this.onValueChanged();
  }
  onValueChanged(data?: any) {
    if (!this.form) { return; }

    this.formErrors = ValidatorUtils.genMsg(this.form, this.validateMsg, []);
  }

  formErrors = [];
  validateMsg = {
    'email': {
      'required': '邮箱不能为空',
      'validate': '邮箱格式错误'
    }
  };

}
