import {Component, ViewEncapsulation, NgModule, Pipe, OnInit, AfterViewInit, ViewChild} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';

import {NgbDatepickerI18n, NgbDateStruct} from '@ng-bootstrap/ng-bootstrap';
import {I18n, CustomDatepickerI18n} from '../../../../service/datepicker-I18n';

import {GlobalState} from '../../../../global.state';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import {ValidatorUtils} from '../../../../validator/validator.utils';
import { RouteService } from '../../../../service/route';

import { PlanService } from '../../../../service/plan';
import { RunService } from '../../../../service/run';

import { PopDialogComponent } from '../../../../components/pop-dialog'

declare var jQuery;

@Component({
  selector: 'plan-edit',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./edit.scss'],
  templateUrl: './edit.html',
  providers: [I18n, {provide: NgbDatepickerI18n, useClass: CustomDatepickerI18n}]
})
export class PlanEdit implements OnInit, AfterViewInit {
  id: number;
  model: any = {};
  form: any;

  @ViewChild('modalDelete') modalDelete: PopDialogComponent;
  @ViewChild('modalSelectCase') modalSelectCase: PopDialogComponent;
  @ViewChild('modalConfigEnvi') modalConfigEnvi: PopDialogComponent;
  @ViewChild('modalRemoveSet') modalRemoveSet: PopDialogComponent;
  testSet: any;
  modalTitle: string;

  constructor(private _state:GlobalState, private _routeService: RouteService, private _route: ActivatedRoute, private fb: FormBuilder,
              private _i18n: I18n, private _planService: PlanService, private _runService: RunService) {

  }
  ngOnInit() {
    let that = this;

    that._route.params.forEach((params: Params) => {
      that.id = +params['id'];
    });

    if (that.id) {
      that.loadData();
    }
    that.buildForm();
  }
  ngAfterViewInit() {}

  buildForm(): void {
    let that = this;
    this.form = this.fb.group(
      {
        'name': ['', [Validators.required]],
        'descr': ['', []],
        'status': ['', [Validators.required]],
        'estimate': ['', []],
        'startTime': ['', []],
        'endTime': ['', []],
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
    'title': {
      'required':      '简介不能为空'
    },
    'objective': {
      'required':      '描述不能为空'
    }
  };

  loadData() {
    let that = this;
    that._planService.get(that.id).subscribe((json:any) => {
      that.model = json.data;
    });
  }

  save() {
    let that = this;

    that._planService.save(that.model).subscribe((json:any) => {
      if (json.code == 1) {
        that.model = json.data;
      }
    });
  }

  reset() {
    this.loadData();
  }

  delete(): void {
    this.modalTitle = "确认删除";
    this.modalDelete.showModal();
  }
  deleteConfirm() {
    this._planService.delete(this.model.id).subscribe((json:any) => {
      if (json.code == 1) {
        this.formErrors = ['删除成功'];
        this.modalDelete.closeModal();
        this._routeService.navTo("/pages/implement/plan/list");
      } else {
        this.formErrors = ['删除失败'];
      }
    });
  }

  editSet(testSet: any): void {
    this.modalTitle = "请选择用例";
    this.testSet = testSet;
    this.modalSelectCase.showModal();
  }
  editSetConfirm() {
    // this._runService.delete(this.testSet.id).subscribe((json:any) => {
    //   if (json.code == 1) {
    //     this.formErrors = ['删除成功'];
    this.modalSelectCase.closeModal();
    this.loadData();
    //   } else {
    //     this.formErrors = ['删除失败'];
    //   }
    // });
  }

  configEnvi(testSet: any): void {
    this.modalTitle = "请设置环境";
    this.testSet = testSet;
    this.modalConfigEnvi.showModal();
  }
  configEnviConfirm() {
    // this._runService.delete(this.testSet.id).subscribe((json:any) => {
    //   if (json.code == 1) {
    //     this.formErrors = ['删除成功'];
    this.modalConfigEnvi.closeModal();
    this.loadData();
    //   } else {
    //     this.formErrors = ['删除失败'];
    //   }
    // });
  }

  removeSet(testSet: any): void {
    this.modalTitle = "确认删除";
    this.testSet = testSet;
    this.modalRemoveSet.showModal();
  }
  removeSetConfirm() {
    this._runService.delete(this.testSet.id).subscribe((json:any) => {
      if (json.code == 1) {
        this.formErrors = ['删除成功'];
        this.modalRemoveSet.closeModal();
      } else {
        this.formErrors = ['删除失败'];
      }
    });
  }

}

