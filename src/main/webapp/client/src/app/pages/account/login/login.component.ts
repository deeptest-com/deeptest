import {Component, ViewEncapsulation} from '@angular/core';
import {FormGroup, AbstractControl, FormBuilder, Validators} from '@angular/forms';

import { CONSTANT } from '../../../utils/constant';

import {ValidatorUtils, EmailValidator} from '../../../validator';

import { RouteService } from '../../../service/route';
import { AccountService } from '../../../service/account';

@Component({
  selector: 'login',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./login.scss'],
  templateUrl: './login.html',
})
export class Login {

  public form:FormGroup;
  model: any = { rememberMe: true};

  constructor(fb:FormBuilder, private accountService: AccountService, private routeService: RouteService) {
    this.form = fb.group({
      'email': ['', [Validators.required, EmailValidator.validate()]],
      'password': ['', [Validators.required, Validators.minLength(6)]],
      'rememberMe': []
    });

    this.form.valueChanges.debounceTime(CONSTANT.DebounceTime).subscribe(data => this.onValueChanged(data));
    this.onValueChanged();
  }

  onValueChanged(data?: any) {
    let that = this;
    if (!that.form) { return; }

    that.formErrors = ValidatorUtils.genMsg(that.form, that.validateMsg, []);
  }

  public onSubmit(values:Object):void {
    let that = this;

    this.accountService.login(this.model).subscribe((errors: any) => {
      this.formErrors = [errors];
    });
  }

  formErrors = [];
  validateMsg = {
    'email': {
      'required':      '邮箱不能为空',
      'validate': '邮箱格式错误'
    },
    'password': {
      'required':      '密码不能为空',
      'minlength': '密码不能少于6位'
    }
  };
}

