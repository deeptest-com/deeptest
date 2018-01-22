import { Component,ViewEncapsulation, Pipe, OnInit, AfterViewInit, ViewChild } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { ActivatedRoute, Params } from '@angular/router';

import { CONSTANT } from '../../../utils/constant';
import {GlobalState} from '../../../global.state';
import {ValidatorUtils} from '../../../validator/validator.utils';
import {PasswordsEqualValidator} from '../../../validator';

import { AccountService } from '../../../service/account';

declare var jQuery;

@Component({
  selector: 'reset-password',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./reset-password.scss'],
  templateUrl: './reset-password.html'
})
export class ResetPassword implements OnInit, AfterViewInit {
  model: any = {};
  form: any;
  vcode: string;
  checkPass: boolean = true;
  msg: string;
  public errors: string;

  constructor(private _state: GlobalState, private _route: ActivatedRoute,
              private fb: FormBuilder, private accountService: AccountService) {

    this._route.params.subscribe(params => {
      this.vcode = params['vcode'];
    });
  }

  ngOnInit() {
    this.buildForm();

    if (this.vcode) {
      this.accountService.checkResetPassword(this.vcode).subscribe((json: any) => {
        if (json.code == 1) {
          this.checkPass = true;
        }  else {
          this.checkPass = false;
          this.msg = json.msg;
        }
        this._state.notifyDataChanged(CONSTANT.EVENT_LOADING_COMPLETE, {});
      });
    }
  }

  ngAfterViewInit() {

  }

  onSubmit():void {
    this.accountService.resetPassword(this.vcode, this.model).subscribe((error: any) => {
      if (error) {
        this.formErrors = [error];
      }
    });
  }

  buildForm(): void {
    this.form = this.fb.group(
      {
        'password': ['', [Validators.required, Validators.pattern(/^[0-9a-zA-Z]{6,10}$/)]],
        'rePassword': ['', [Validators.required, Validators.pattern(/^[0-9a-zA-Z]{6,10}$/)]],
      }, {validator: PasswordsEqualValidator.validate('passwordsEqual', 'password', 'rePassword')}
    );

    this.form.valueChanges.debounceTime(CONSTANT.DebounceTime).subscribe(data => this.onValueChanged(data));
    this.onValueChanged();
  }
  onValueChanged(data?: any) {
    this.formErrors = ValidatorUtils.genMsg(this.form, this.validateMsg, ['passwordsEqual']);
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
