import { Component,ViewEncapsulation, Pipe, OnInit, AfterViewInit, ViewChild } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Router, ActivatedRoute, Params } from '@angular/router';

import { DropdownModule} from 'ng2-bootstrap/ng2-bootstrap';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';
import {Validate} from '../../../service/validate';

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
  companyId: number;
  model: any = { signBefore: 3};
  companyForm: any;
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
      that.companyId = +params['id'];
    });

    if (that.companyId) {
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

    that._routeService.navTo('/pages/company/edit/' + that.companyId + '/' + $company.tabModel);
  }
  loadData() {
   let that = this;

   that._companyService.get(that.companyId).subscribe((json:any) => {
      that.model = json.company;

     that.initForm(true);
   });
  }

  initForm(dataLoaded): void {
    let that = this;
  }

  buildForm(): void {
    let that = this;
    that.companyForm = that.fb.group(
        {
          'name': [that.model.email, [Validators.required]]
        }, {
           validator: Validate.compareDatetime([])
        }
    );

    that.companyForm.valueChanges.subscribe(data => that.onValueChanged(data));
    that.onValueChanged();
  }
  onValueChanged(data?: any) {
    let that = this;
    if (!that.companyForm) { return; }

    that.formErrors = Validate.genValidateInfo(that.companyForm, that.validateMsg, []);
  }

  formErrors = [];
  validateMsg = {
    'name': {
      'required':      '公司名称不能为空'
    }
  };

}
