import { Component,ViewEncapsulation, Pipe, OnInit, AfterViewInit, ViewChild } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Router, ActivatedRoute, Params } from '@angular/router';

import { DropdownModule} from 'ng2-bootstrap/ng2-bootstrap';
import {EmailValidator} from '../../../../validator';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import {ValidatorUtils} from '../../../../validator/validator.utils';

import { RouteService } from '../../../../service/route';

import { UserService } from '../../../../service/user';

declare var jQuery;

@Component({
  selector: 'profile-edit-property',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./edit.scss')],

  template: require('./edit.html')
})
export class ProfileEdit implements OnInit, AfterViewInit {
  modelId: number;
  model: any = {};
  form: any;
  tabModel: string = 'property';
  needCreate:boolean = false;

  constructor(private _routeService: RouteService, private _route: ActivatedRoute, private fb: FormBuilder,
              private _userService: UserService) {

    let that = this;
  }

  ngOnInit() {
    let that = this;

    that.buildForm();
    this._route.params.forEach((params: Params) => {
      that.modelId = +params['id'];
    });

    if (that.modelId) {
        that.loadData();
    }
  }

  ngAfterViewInit() {
    let that = this;

    that.initForm(false);
  }

  onSubmit():void {
    let that = this;

    that.model.status = undefined;

    that._userService.saveProfile(that.model).subscribe((json:any) => {
        if (json.code = 1) {
          that._routeService.navTo("/pages/profile/list");
        }
    });
  }

  goto($profile) {
    let that = this;

    that._routeService.navTo('/pages/profile/edit/' + that.modelId + '/' + $profile.tabModel);
  }
  loadData() {
   let that = this;

   that._userService.getProfile(that.modelId).subscribe((json:any) => {
      that.model = json.profile;

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
        'name': [that.model['name'], [Validators.required, Validators.minLength(4)]],
        'phone': [that.model['phone'], [Validators.required, Validators.minLength(11)]],
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
