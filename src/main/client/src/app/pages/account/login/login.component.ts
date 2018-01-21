import {Component, OnInit, ViewEncapsulation} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import {FormGroup, AbstractControl, FormBuilder, Validators} from '@angular/forms';

import { CONSTANT } from '../../../utils/constant';
import {GlobalState} from '../../../global.state';
import {ValidatorUtils} from '../../../validator';

import { RouteService } from '../../../service/route';
import { AccountService } from '../../../service/account';

@Component({
  selector: 'login',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./login.scss'],
  templateUrl: './login.html',
})
export class Login implements OnInit {
  vcode: string;
  public form:FormGroup;
  model: any = { rememberMe: true};

  constructor(private _state: GlobalState, fb:FormBuilder, private _route: ActivatedRoute, private accountService: AccountService, private routeService: RouteService) {
    this.form = fb.group({
      'email': ['', [Validators.required, Validators.email]],
      'password': ['', [Validators.minLength(6)]],
      'rememberMe': []
    });

    this.form.valueChanges.debounceTime(CONSTANT.DebounceTime).subscribe(data => this.onValueChanged(data));
    this.onValueChanged();
  }
  ngOnInit() {
    this._route.params.subscribe(params => {
      this.vcode = params['vcode'];
    });
    if(this.vcode) {
      this.accountService.loginWithVcode(this.vcode).subscribe((errors: any) => {
        this.formErrors = [errors];
        this._state.notifyDataChanged(CONSTANT.EVENT_LOADING_COMPLETE, {});
      });
    } else {
      this._state.notifyDataChanged(CONSTANT.EVENT_LOADING_COMPLETE, {});
    }
  }

  onValueChanged(data?: any) {
    if (!this.form) { return; }

    this.formErrors = ValidatorUtils.genMsg(this.form, this.validateMsg, []);
  }

  public onSubmit(values: Object):void {
    this.accountService.login(this.model).subscribe((errors: any) => {
      this.formErrors = [errors];
    });
  }

  formErrors = [];
  validateMsg = {
    'email': {
      'required':      '邮箱不能为空',
      'email': '邮箱格式错误'
    },
    'password': {
      'required':      '密码不能为空',
      'minlength': '密码不能少于6位'
    }
  };
}

