import {Component, ViewEncapsulation} from '@angular/core';
import {FormGroup, AbstractControl, FormBuilder, Validators} from '@angular/forms';

import { Cookie } from 'ng2-cookies/ng2-cookies';

import { CONSTANT } from '../../../utils/constant';
import { RouteService } from '../../../service/route';
import { UserService } from '../../../service/user';

@Component({
  selector: 'login',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./login.scss')],
  template: require('./login.html'),
})
export class Login {

  public form:FormGroup;
  public email:AbstractControl;
  public password:AbstractControl;
  public rememberMe:AbstractControl;

  public submitted:boolean = false;
  public errors: string;

  constructor(fb:FormBuilder, private userService: UserService, private routeService: RouteService) {
    this.form = fb.group({
      'email': ['', Validators.compose([Validators.required, Validators.minLength(4)])],
      'password': ['', Validators.compose([Validators.required, Validators.minLength(4)])],
      'rememberMe': ['', null]
    });

    this.email = this.form.controls['email'];
    this.password = this.form.controls['password'];
    this.rememberMe = this.form.controls['rememberMe'];
  }

  public onSubmit(values:Object):void {
    let that = this;
    this.submitted = true;

    this.userService.login(values['email'], values['password'], values['rememberMe']).subscribe((json:any) => {
      if (json.code == 1) {
        that.errors = undefined;

        Cookie.set(CONSTANT.COOKIE_KEY, CONSTANT.TOKEN, 30);

        that.routeService.navTo('/pages/dashboard');
      } else {
        that.errors = json.msg;
      }
    });
  }
}
