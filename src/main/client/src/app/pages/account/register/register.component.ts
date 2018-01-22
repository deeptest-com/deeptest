import {Component, OnInit, ViewEncapsulation} from '@angular/core';
import {FormGroup, FormBuilder, Validators} from '@angular/forms';
import {ValidatorUtils, PhoneValidator, PasswordsEqualValidator} from '../../../validator';

import { CONSTANT } from '../../../utils/constant';
import {GlobalState} from '../../../global.state';
import {SlimLoadingBarService} from "../../../components/ng2-loading-bar";
import { AccountService } from '../../../service/account';

@Component({
  selector: 'register',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./register.scss'],
  templateUrl: './register.html',
})
export class Register implements OnInit {

  public form:FormGroup;

  public model: any = {};

  constructor(private _state: GlobalState, fb:FormBuilder, private accountService: AccountService,
              private slimLoadingBarService: SlimLoadingBarService) {

    this.form = fb.group({
      'name': ['', [Validators.required, Validators.minLength(2)]],
      'email': ['', [Validators.required, Validators.email]],
      'phone': ['', [Validators.required, PhoneValidator.validate()]],
      'password': ['', [Validators.required, Validators.minLength(6)]],
      'repeatPassword': ['', [Validators.required, Validators.minLength(6)]]
      },
      {
        validator: PasswordsEqualValidator.validate('passwordsEqual', 'password', 'repeatPassword')
      });
    this.form.valueChanges.debounceTime(CONSTANT.DebounceTime).subscribe(data => this.onValueChanged(data));
    this.onValueChanged();
  }

  ngOnInit() {
    this._state.notifyDataChanged(CONSTANT.EVENT_LOADING_COMPLETE, {});
  }
  onValueChanged(data?: any) {
    let that = this;
    if (!that.form) { return; }

    that.formErrors = ValidatorUtils.genMsg(that.form, that.validateMsg, ['passwordsEqual']);
  }

  public onSubmit():void {
    this.formErrors = [];
    this.slimLoadingBarService.start(() => { console.log('Loading complete'); });

    this.accountService.register(this.model).subscribe((json: any) => {
      this.formErrors = [json.msg];

      this.slimLoadingBarService.complete();
    });
  }

  formErrors = [];
  validateMsg = {
    'name': {
      'required':      '姓名不能为空',
      'minlength': '姓名不能少于2位'
    },
    'email': {
      'required':      '邮箱不能为空',
      'email': '邮箱格式错误'
    },
    'phone': {
      'required':      '手机不能为空',
      'validate': '手机号码格式错误'
    },
    'password': {
      'required':      '密码不能为空',
      'minlength': '密码不能少于6位'
    },
    'repeatPassword': {
      'required':      '重复密码不能为空',
      'minlength': '重复密码不能少于6位'
    },
    'passwordsEqual': '两次密码不一致'
  };
}
