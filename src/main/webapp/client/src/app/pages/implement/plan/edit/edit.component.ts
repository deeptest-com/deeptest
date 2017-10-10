import {Component, ViewEncapsulation, NgModule, Pipe, Input, Compiler, OnInit, AfterViewInit, ViewChild} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';

import {NgbDatepickerI18n, NgbDateParserFormatter, NgbDateStruct, NgbModal, NgbModalRef, ModalDismissReasons} from '@ng-bootstrap/ng-bootstrap';

import {I18n, CustomDatepickerI18n} from '../../../../service/datepicker-I18n';

import {GlobalState} from '../../../../global.state';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import {ValidatorUtils, CustomValidator} from '../../../../validator';
import { RouteService } from '../../../../service/route';

import { PlanService } from '../../../../service/plan';
import { RunService } from '../../../../service/run';
import { CaseService } from '../../../../service/case';
import { UserService } from '../../../../service/user';

import { CaseSelectionComponent } from '../../../../components/case-selection'
import { EnvironmentConfigComponent } from '../../../../components/environment-config'
import { PopDialogComponent } from '../../../../components/pop-dialog'

declare var jQuery;

@Component({
  selector: 'plan-edit',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./edit.scss',
    '../../../../../vendor/ztree/css/zTreeStyle/zTreeStyle.css',
    '../../../../components/ztree/src/styles.scss'],
  templateUrl: './edit.html',
  providers: [I18n, {provide: NgbDatepickerI18n, useClass: CustomDatepickerI18n}]
})
export class PlanEdit implements OnInit, AfterViewInit {
  treeSettings: any = {usage: 'selection', isExpanded: true, sonSign: false};

  projectId: number;
  planId: number;
  startDate: any;
  model: any = {};
  run: any = {};
  runIndex: number;
  form: FormGroup;

  @ViewChild('modalEditRun') modalEditRun: PopDialogComponent;
  @ViewChild('modalSelectCase') modalSelectCase: CaseSelectionComponent;

  @ViewChild('modalDelete') modalDelete: PopDialogComponent;
  @ViewChild('modalConfigEnvi') modalConfigEnvi: PopDialogComponent;
  @ViewChild('modalRemoveSet') modalRemoveSet: PopDialogComponent;
  testSet: any;
  modalTitle: string;
  caseSelectionModal: any;
  envSelectionModal: any;

  constructor(private _state:GlobalState, private _routeService: RouteService, private _route: ActivatedRoute, private fb: FormBuilder,
              private _i18n: I18n, private modalService: NgbModal, private compiler: Compiler, private ngbDateParserFormatter: NgbDateParserFormatter,
              private _planService: PlanService, private _runService: RunService, private _caseService: CaseService, private _userService: UserService) {

    this.projectId = CONSTANT.CURRENT_PROJECT.id;
  }
  ngOnInit() {
    this._route.params.forEach((params: Params) => {
      this.planId = +params['planId'];
    });

    if (this.planId) {
      this.loadData();
    }
    this.buildForm();

    var now = new Date();
    this.startDate = { day: now.getDate(), month: now.getMonth() + 1, year: now.getFullYear()};
  }
  ngAfterViewInit() {}

  buildForm(): void {
    this.form = this.fb.group(
      {
        'name': ['', [Validators.required]],
        'descr': ['', []],
        'estimate': ['', [Validators.pattern(/^(?!0+(?:\.0+)?$)(?:[1-9]\d*|0)(?:\.\d{1,2})?$/)]],
        'startTime': ['', []],
        'endTime': ['', []],
        'disabled': ['', []]
      }, {
        validator: CustomValidator.compareDate('dateCompare', 'startTime', 'endTime')
      }
    );

    this.form.valueChanges.debounceTime(CONSTANT.DebounceTime).subscribe(data => this.onValueChanged(data));
    this.onValueChanged();
  }
  onValueChanged(data?: any) {
    this.formErrors = ValidatorUtils.genMsg(this.form, this.validateMsg, ['dateCompare']);
  }

  formErrors = [];
  validateMsg = {
    'name': {
      'required':      '名称不能为空'
    },
    'objective': {
      'required':      '测试目的不能为空'
    },
    'estimate': {
      'pattern':      '耗时必须是最多含2位小数的数字'
    },
    dateCompare: '结束时间必须大于或等于开始时间'
  };

  loadData() {
    let that = this;
    that._planService.get(this.planId).subscribe((json:any) => {
      that.model = json.data;

      this.model.startTime = this.ngbDateParserFormatter.parse(that.model.startTime);
      this.model.endTime = this.ngbDateParserFormatter.parse(that.model.endTime);
    });
  }

  save() {
    this._planService.save(this.model).subscribe((json:any) => {
      if (json.code == 1) {
        this._routeService.navTo("/pages/implement/" + CONSTANT.CURRENT_PROJECT.id + "/plan/list");
      } else {
        this.formErrors = [json.msg];
      }
    });
  }

  reset() {
    this.loadData();
  }

  createRun() {
    this.run = {};
    this.runIndex = this.model.runVos.length;
    this.modalEditRun.showModal();
  }
  editRun(run: any, index: number) {
    this.run = run;
    this.runIndex = index;
    this.modalEditRun.showModal();
  }
  saveRun() {
    this._runService.saveRun(this.planId, this.run).subscribe((json:any) => {
      this.model.runVos[this.runIndex]= json.data;
      this.modalEditRun.closeModal();
    });
  }

  editRunCases(run: any, index: number) {
    this.compiler.clearCacheFor(CaseSelectionComponent);
    this.caseSelectionModal = this.modalService.open(CaseSelectionComponent, {windowClass: 'pop-selection'});
    this.caseSelectionModal.componentInstance.treeSettings = this.treeSettings;

    this._caseService.query(CONSTANT.CURRENT_PROJECT.id).subscribe((json:any) => {
      this.caseSelectionModal.componentInstance.treeModel = json.data;
    });
    this._userService.getUsers(CONSTANT.CURRENT_PROJECT.id).subscribe((json:any) => {
      this.caseSelectionModal.componentInstance.users = json.data;
    });

    this.caseSelectionModal.result.then((result) => {
      let id = run? run.id: undefined;
      this.saveRunCases(id, result.data, index);
    }, (reason) => {
      console.log('reason', reason);
    });
  }
  saveRunCases(runId: any, cases: any[], index: number): void {
    this._runService.saveRunCases(this.planId, runId, cases).subscribe((json:any) => {
      console.log(json);
      this.model.runVos[index]= json.data;
    });
  }

  editEnvi(testSet: any): void {
    this.compiler.clearCacheFor(EnvironmentConfigComponent);
    this.envSelectionModal = this.modalService.open(EnvironmentConfigComponent, {windowClass: 'pop-selection'});
    this.envSelectionModal.result.then((result) => {
      console.log('result', result);
    }, (reason) => {
      console.log('reason', reason);
    });
    this.envSelectionModal.componentInstance.testSet = testSet;
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

  endTimeChanged() {
    console.log('===', this.model);
  }

}

