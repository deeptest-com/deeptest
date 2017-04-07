import {Component, ViewEncapsulation, ViewChild} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import { NgModule, Pipe, OnInit, AfterViewInit }      from '@angular/core';
import { ModalDirective } from 'ng2-bootstrap';
import {GlobalState} from '../../../../../global.state';

import { CONSTANT } from '../../../../../utils/constant';
import { Utils } from '../../../../../utils/utils';
import {ValidatorUtils, EmailValidator, PhoneValidator} from '../../../../../validator';
import { RouteService } from '../../../../../service/route';

import { CustomFieldService } from '../../../../../service/custom-field';
import { PopDialogComponent } from '../../../../../components/pop-dialog'

declare var jQuery;

@Component({
  selector: 'custom-field-edit',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./edit.scss')],
  template: require('./edit.html')
})
export class CustomFieldEdit implements OnInit, AfterViewInit {

  id: number;
  tab: string = 'info';

  model: any = {};
  relations: any[] = [];
  form: FormGroup;
  isSubmitted: boolean;
  @ViewChild('modalWrapper') modalWrapper: PopDialogComponent;

  constructor(private _state:GlobalState, private _routeService: RouteService, private _route: ActivatedRoute,
              private fb: FormBuilder, private customFieldService: CustomFieldService) {

  }
  ngOnInit() {
    this._route.params.forEach((params: Params) => {
      this.id = +params['id'];
    });

    this.loadData();
    this.buildForm();
  }
  ngAfterViewInit() {}


  selectTab(tab: string) {
    let that = this;
    that.tab = tab;
  }

  buildForm(): void {
    let that = this;
    this.form = this.fb.group(
      {
        'name': ['', [Validators.required]],
        'code': ['', [Validators.required]],
        'applyTo': ['', [Validators.required]],
        type: ['', [Validators.required]],
        format: ['', [Validators.required]],
        descr: ['', []],
        isGlobal: ['', []],
        isRequired: ['', []],
        'disabled': ['', []]
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
      'required':      '姓名不能为空'
    },
    'code': {
      'required':      '编码不能为空',
    },
    'applyTo': {
      'required':      '应用对象不能为空'
    },
    'type': {
      'required':      '类型不能为空'
    },
    'format': {
      'required':      '格式不能为空'
    }
  };

  loadData() {
    let that = this;
    that.customFieldService.get(that.id).subscribe((json:any) => {
      that.model = json.data;
    });
  }

  save() {
    let that = this;

    that.customFieldService.save(that.model, that.relations).subscribe((json:any) => {
      if (json.code == 1) {

        that.formErrors = ['保存成功'];
        that._routeService.navTo("/pages/org-admin/field/list");
      } else {
        that.formErrors = [json.msg];
      }
    });
  }

  delete() {
    let that = this;

    that.customFieldService.delete(that.model.id).subscribe((json:any) => {
      if (json.code == 1) {
        that.formErrors = ['删除成功'];
        that._routeService.navTo("/pages/org-admin/field/list");
      } else {
        that.formErrors = ['删除失败'];
      }
    });
  }

  select(key: string) {
    let val = key ==='all'? true: false;
    for (let group of this.relations) {
      group.selecting = val;
    }
  }

  showModal(): void {
    this.modalWrapper.showModal();
  }

}

