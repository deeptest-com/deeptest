import {Component, ViewEncapsulation} from '@angular/core';
import {FormGroup, AbstractControl, FormBuilder, Validators} from '@angular/forms';
import {EmailValidator, EqualPasswordsValidator} from '../../../validator';

import { Cookie } from 'ng2-cookies/ng2-cookies';

import { CONSTANT } from '../../../utils/constant';
import { RouteService } from '../../../service/route';
import { UserService } from '../../../service/user';

@Component({
  selector: 'register',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./register.scss')],
  template: require('./register.html'),
})
export class Register {

  public form:FormGroup;
  public name:AbstractControl;
  public phone:AbstractControl;
  public email:AbstractControl;
  public password:AbstractControl;
  public repeatPassword:AbstractControl;
  public passwords:FormGroup;

  public submitted:boolean = false;
  public errors: string;

  constructor(fb:FormBuilder, private userService: UserService, private routeService: RouteService) {

    this.form = fb.group({
      'name': ['', Validators.compose([Validators.required, Validators.minLength(4)])],
      'phone': ['', Validators.compose([Validators.required, Validators.minLength(11)])],
      'email': ['', Validators.compose([Validators.required, EmailValidator.validate()])],
      'passwords': fb.group({
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

  public onSubmit(values:Object):void {
    let that = this;
    this.submitted = true;
    if (this.form.valid) {
      this.userService.register(values['name'], values['phone'], values['email'], values['passwords']['password']).subscribe((err:any) => {
        that.errors = err;
      });
    }
  }
}
