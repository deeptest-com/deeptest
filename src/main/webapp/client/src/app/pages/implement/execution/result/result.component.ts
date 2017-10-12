import {Component, ViewEncapsulation, NgModule, Pipe, OnInit, AfterViewInit} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { NgUploaderModule } from 'ngx-uploader';

import {GlobalState} from '../../../../global.state';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import {ValidatorUtils} from '../../../../validator/validator.utils';
import { RouteService } from '../../../../service/route';

import { CaseService } from '../../../../service/case';
import { CaseStepService } from '../../../../service/case-step';
import { CaseInRunService } from '../../../../service/case-in-run';

declare var jQuery;

@Component({
  selector: 'execution-result',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./result.scss'],
  templateUrl: './result.html'
})
export class ExecutionResult implements OnInit, AfterViewInit {
  planId: number;
  runId: number;

  id: number;
  model: any;
  settings: any;
  data: any;
  form: any;
  tab: string = 'info';

  fields: any;
  next: boolean = false;

  constructor(private _state:GlobalState, private _routeService: RouteService, private _route: ActivatedRoute, private fb: FormBuilder,
              private _caseService: CaseService, private _caseStepService: CaseStepService, private _caseInRunService: CaseInRunService) {

  }
  ngOnInit() {
    let that = this;

    that._route.params.forEach((params: Params) => {
      that.planId = +params['planId'];
      that.runId = +params['runId'];
    });

    this.fields = CONSTANT.CUSTOM_FIELD_FOR_PROJECT;

    that.buildForm();

    this._state.subscribe('case.change', (testCase: any) => {
      if (!testCase || testCase.isParent) {
        this.model = null;
        return;
      }

      this.fields = CONSTANT.CUSTOM_FIELD_FOR_PROJECT;
      if (testCase) {
        that.id = testCase.entityId;

        that.loadData();
      } else {
        that.model = undefined;
      }
    });

    this.settings = {
      columns: {
        ordr: {
          title: '顺序',
        },
        opt: {
          title: '操作',
          editor: {
            type: 'textarea'
          },
        },
        expect: {
          title: '期望结果',
          editor: {
            type: 'textarea',
          },
        }
      },
    };
  }
  ngAfterViewInit() {}

  buildForm(): void {
    let that = this;
    this.form = this.fb.group(
      {
        'title': ['', [Validators.required]],
        'objective': ['', [Validators.required]],
        'pre_condition': ['', []],
        'next':  ['', []]
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
    'title': {
      'required':      '简介不能为空'
    },
    'objective': {
      'required':      '描述不能为空'
    }
  };

  loadData() {
    let that = this;
    that._caseInRunService.get(that.id).subscribe((json:any) => {
      that.model = json.data;
    });
  }

  pass() {
    console.log(this.model);

    // this._caseService.save(that.model).subscribe((json:any) => {
    //   if (json.code == 1) {
    //     this.model = json.data;
    //   }
    // });
  }

  reset() {
    this.loadData();
  }

  saveField (field: any) {
    this._caseService.saveField(this.model.id, field).subscribe((json:any) => {
      if (json.code == 1) {
        this.model = json.data;
      }
    });
  }

  tabChange(event: any) {
    this.tab = event.nextId;
  }

  onUpConfirm(event: any) {
    console.log('onUpConfirm', event);
    this._caseStepService.up(event.data).subscribe((json:any) => {
      event.confirm.resolve();
    });
  }

  onDownConfirm(event: any) {
    console.log('onDownConfirm', event);
    this._caseStepService.down(event.data).subscribe((json:any) => {
      event.confirm.resolve();
    });
  }

  onCreateConfirm(event: any) {
    console.log('onCreateConfirm', event);
    event.confirm.resolve();
  }
  onSaveConfirm(event: any) {
    console.log('onSaveConfirm', event);
    this._caseStepService.save(this.id, event.data).subscribe((json:any) => {
      event.confirm.resolve();
    });
  }
  onDeleteConfirm(event: any) {
    console.log('onDeleteConfirm', event);
    this._caseStepService.delete(event.data).subscribe((json:any) => {
      event.confirm.resolve();
    });
  }

  returnTo() {
    let url: string = '/pages/implement/' + CONSTANT.CURRENT_PROJECT.id + '/plan/' + this.planId + '/view';
    console.log(url);
    this._routeService.navTo(url);
  }

}

