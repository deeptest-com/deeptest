import {Component, ViewEncapsulation, ViewChild} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import { NgModule, Pipe, OnInit, AfterViewInit }      from '@angular/core';

import { NgbModalModule, NgbPaginationModule, NgbDropdownModule,
  NgbTabsetModule, NgbButtonsModule, NgbCollapseModule } from '@ng-bootstrap/ng-bootstrap';
import { BrowserModule } from '@angular/platform-browser';

import {GlobalState} from '../../../../../global.state';

import { CONSTANT } from '../../../../../utils/constant';
import { Utils } from '../../../../../utils/utils';
import {ValidatorUtils, PhoneValidator} from '../../../../../validator';
import { RouteService } from '../../../../../service/route';

import { CaseTypeService } from '../../../../../service/case-type';
import { PopDialogComponent } from '../../../../../components/pop-dialog'

declare var jQuery;

@Component({
  selector: 'case-type-edit',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./edit.scss'],
  templateUrl: './edit.html'
})
export class CaseTypeEdit implements OnInit, AfterViewInit {

  id: number;

  model: any = {};
  form: FormGroup;
  isSubmitted: boolean;
  @ViewChild('modalWrapper') modalWrapper: PopDialogComponent;

  constructor(private _state:GlobalState, private _routeService: RouteService, private _route: ActivatedRoute,
              private fb: FormBuilder, private caseTypeService: CaseTypeService) {

  }
  ngOnInit() {
    this._route.params.forEach((params: Params) => {
      this.id = +params['id'];
    });

    this.loadData();
    this.buildForm();
  }
  ngAfterViewInit() {}

  buildForm(): void {
    let that = this;
    this.form = this.fb.group(
      {
        'name': ['', [Validators.required]],
        'descr': ['', []],
        'disabled': ['', []],
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
      'required':      '名称不能为空'
    }
  };

  loadData() {
    let that = this;
    that.caseTypeService.get(that.id).subscribe((json:any) => {
      that.model = json.data;
    });
  }

  save() {
    let that = this;

    that.caseTypeService.save(that.model).subscribe((json:any) => {
      if (json.code == 1) {
        CONSTANT.CASE_PROPERTY_MAP = json.casePropertyMap;

        that.formErrors = ['保存成功'];
        that._routeService.navTo("/pages/org-admin/property/case-type/list");
      } else {
        that.formErrors = [json.msg];
      }
    });
  }

  delete() {
    let that = this;

    that.caseTypeService.delete(that.model.id).subscribe((json:any) => {
      if (json.code == 1) {
        that.formErrors = ['删除成功'];
        that._routeService.navTo("/pages/org-admin/property/case-type/list");
      } else {
        that.formErrors = ['删除失败'];
      }
    });
  }

  showModal(): void {
    this.modalWrapper.showModal();
  }

}

