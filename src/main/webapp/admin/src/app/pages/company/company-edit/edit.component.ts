import { Component,ViewEncapsulation, Pipe, OnInit, AfterViewInit, ViewChild } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Router, ActivatedRoute, Params } from '@angular/router';

import { DropdownModule} from 'ng2-bootstrap/ng2-bootstrap';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';
import {ValidatorUtils} from '../../../validator/validator.utils';

import { RouteService } from '../../../service/route';

import { CompanyService } from '../../../service/company';

declare var jQuery;

@Component({
  selector: 'company-edit-property',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./edit.scss')],

  template: require('./edit.html')
})
export class CompanyEdit implements OnInit, AfterViewInit {
  modelId: number;
  model: any = {};
  form: any;
  tabModel: string = 'property';
  needCreate:boolean = false;

  constructor(private _routeService: RouteService, private _route: ActivatedRoute, private fb: FormBuilder,
              private _companyService: CompanyService) {

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

    that._companyService.save(that.model).subscribe((json:any) => {
        if (json.code = 1) {
          that._routeService.navTo("/pages/company/list");
        }
    });
  }

  goto($company) {
    let that = this;

    that._routeService.navTo('/pages/company/edit/' + that.modelId + '/' + $company.tabModel);
  }
  loadData() {
   let that = this;

   that._companyService.get(that.modelId).subscribe((json:any) => {
      that.model = json.company;

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
        'name': [that.model['name'], [Validators.required]]
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
      'required':      '公司名称不能为空'
    }
  };

}
