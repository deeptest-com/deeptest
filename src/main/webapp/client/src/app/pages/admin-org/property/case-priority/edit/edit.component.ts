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

import { CasePriorityService } from '../../../../../service/case-priority';
import { PopDialogComponent } from '../../../../../components/pop-dialog'

declare var jQuery;

@Component({
  selector: 'case-priority-edit',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./edit.scss')],
  template: require('./edit.html')
})
export class CasePriorityEdit implements OnInit, AfterViewInit {

  id: number;
  tab: string = 'info';

  field: any = {};
  relations: any[] = [];
  form: FormGroup;
  isSubmitted: boolean;
  @ViewChild('modalWrapper') modalWrapper: PopDialogComponent;

  constructor(private _state:GlobalState, private _routeService: RouteService, private _route: ActivatedRoute,
              private fb: FormBuilder, private casePriorityService: CasePriorityService) {

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
        'email': ['', [Validators.required, EmailValidator.validate()]],
        'phone': ['', [Validators.required, PhoneValidator.validate()]],
        'disabled': ['', []],
        'groups': ['', []]
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
    'email': {
      'required':      '邮箱不能为空',
      'validate':      '邮箱格式不正确'
    },
    'phone': {
      'required':      '手机不能为空',
      'validate':      '手机格式不正确'
    }
  };

  loadData() {
    let that = this;
    that.casePriorityService.get(that.id).subscribe((json:any) => {
      that.field = json.field;
      that.relations = json.relations;

      _.forEach(that.relations, (group: any, index: number) => {
        this.form.addControl('group-' + group.orgGroupId, new FormControl('', []))
      });
    });
  }

  save() {
    let that = this;

    that.casePriorityService.save(that.field).subscribe((json:any) => {
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

    that.casePriorityService.delete(that.field.id).subscribe((json:any) => {
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

