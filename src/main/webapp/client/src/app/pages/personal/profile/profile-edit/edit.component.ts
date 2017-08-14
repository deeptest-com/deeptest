import { Component,ViewEncapsulation, Pipe, OnInit, AfterViewInit, ViewChild } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Router, ActivatedRoute, Params } from '@angular/router';

import {EmailValidator} from '../../../../validator';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import {ValidatorUtils} from '../../../../validator/validator.utils';

import { RouteService } from '../../../../service/route';

import {AccountService} from './../../../../service/account';

declare var jQuery;

@Component({
  selector: 'profile-edit-property',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./edit.scss'],
  template: './edit.html'
})
export class ProfileEdit implements OnInit, AfterViewInit {
  modelId: number;
  model: any = {};
  form: any;
  isSubmitted: boolean;
  tabModel: string = 'property';
  needCreate:boolean = false;

  constructor(private _routeService: RouteService, private _route: ActivatedRoute, private fb: FormBuilder,
              private accountService: AccountService) {

    let that = this;
  }

  ngOnInit() {
    let that = this;

    that.buildForm();

    that.loadData();
  }

  ngAfterViewInit() {
    let that = this;
  }

  onSubmit():void {
    let that = this;

    that.accountService.saveProfile(that.model).subscribe((json:any) => {
        if (json.code == 1) {
          Utils.saveProfileLocal(json.data, null);
          that.formErrors = ['保存成功'];
        }
    });
  }

  loadData() {
   let that = this;

   that.accountService.getProfile().subscribe((json:any) => {
      that.model = json.data;
   });
  }

  buildForm(): void {
    let that = this;
    this.form = this.fb.group(
      {
        'name': ['', [Validators.required, Validators.minLength(4)]],
        'phone': ['', [Validators.required, Validators.minLength(11)]],
        'email': ['', [Validators.required, EmailValidator.validate()]]
      }, {}
    );

    this.form.valueChanges.debounceTime(CONSTANT.DebounceTime).subscribe(data => this.onValueChanged(data));
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
