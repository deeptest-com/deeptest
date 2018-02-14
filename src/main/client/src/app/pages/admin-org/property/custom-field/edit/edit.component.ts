import {Component, ViewEncapsulation, ViewChild, Compiler} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import { NgModule, Pipe, OnInit, AfterViewInit }      from '@angular/core';

import {NgbModal} from '@ng-bootstrap/ng-bootstrap';

import * as _ from 'lodash';
import {GlobalState} from '../../../../../global.state';

import { CONSTANT } from '../../../../../utils/constant';
import { Utils } from '../../../../../utils/utils';
import {ValidatorUtils, CustomValidator} from '../../../../../validator';
import { RouteService } from '../../../../../service/route';

import { CustomFieldService } from '../../../../../service/custom-field';
import { PopDialogComponent } from '../../../../../components/pop-dialog';

import { DropdownOptionsComponent } from '../../../../../components/dropdown-options';

declare var jQuery;

@Component({
  selector: 'custom-field-edit',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./edit.scss'],
  templateUrl: './edit.html'
})
export class CustomFieldEdit implements OnInit, AfterViewInit {

  id: number;
  tab: string = 'info';

  model: any = {};
  applyToList: string[];
  typeList: string[];
  formatList: string[];

  relations: any[] = [];
  form: FormGroup;
  isSubmitted: boolean;
  @ViewChild('modalWrapper') modalWrapper: PopDialogComponent;
  public dropdownOptionsModal: any;

  constructor(private _state:GlobalState, private _routeService: RouteService, private _route: ActivatedRoute,
              private fb: FormBuilder, private customFieldService: CustomFieldService,
              private compiler: Compiler, private modalService: NgbModal) {

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
    this.form = this.fb.group(
      {
        label: ['', [Validators.required]],
        applyTo: ['', [Validators.required]],
        myColumn: ['', [Validators.required]],
        type: ['', [Validators.required]],
        rows:  ['', [Validators.pattern('^[1-9]$'), CustomValidator.validate('required_if_other_is', 'required_rows', 'rows', 'type', 'text')]],
        format: ['', [CustomValidator.validate('required_if_other_is', 'required_format', 'format', 'type', 'text')]],
        descr: ['', []],
        global: ['', []],
        isRequired: ['', []]
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
    'code': {
      'required':      '编码不能为空'
    },
    'label': {
      'required':      '名称不能为空'
    },
    'applyTo': {
      'required':      '应用对象不能为空'
    },
    'type': {
      'required':      '类型不能为空'
    },
    'rows': {
      'pattern': '字段行数必须为1-9的整数',
      'required_rows':      '字段行数不能为空'
    },
    'format': {
      'required_format':      '字段格式不能为空'
    }
  };

  loadData() {
    let that = this;
    that.customFieldService.get(that.id).subscribe((json:any) => {
      that.model = json.data;

      that.applyToList = json.applyToList;
      that.typeList = json.typeList;
      that.formatList = json.formatList;
      that.relations = json.projects;

      _.forEach(that.relations, (project: any, index: number) => {
        this.form.addControl('project-' + project.id, new FormControl('', []))
      });
    });
  }

  save() {
    let that = this;

    that.customFieldService.save(that.model, that.relations).subscribe((json:any) => {
      if (json.code == 1) {

        that.formErrors = ['保存成功'];
        that._routeService.navTo("/pages/org-admin/property/custom-field/list");
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
        this.modalWrapper.closeModal();
        that._routeService.navTo("/pages/org-admin/property/custom-field/list");
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

  editDropdownOptions() {
    this.compiler.clearCacheFor(DropdownOptionsComponent);

    this.dropdownOptionsModal = this.modalService.open(DropdownOptionsComponent, {windowClass: 'pop-modal'});
    this.dropdownOptionsModal.componentInstance.title = this.model.label;
    this.dropdownOptionsModal.componentInstance.field = this.model;

    this.dropdownOptionsModal.result.then((result) => {
      this.model.options = result.data;
      this.save();
    }, (reason) => {
      console.log('reason', reason);
    });
  }

  showModal(): void {
    this.modalWrapper.showModal();
  }

}

