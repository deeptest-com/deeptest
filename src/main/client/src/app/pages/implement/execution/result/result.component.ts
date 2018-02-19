import {Component, ViewEncapsulation, NgModule, Pipe, OnInit, AfterViewInit, OnDestroy} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';

import {GlobalState} from '../../../../global.state';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import {ValidatorUtils} from '../../../../validator/validator.utils';
import { RouteService } from '../../../../service/route';

import { CaseService } from '../../../../service/case';
import { CaseStepService } from '../../../../service/case-step';
import { CaseInRunService } from '../../../../service/case-in-run';
import { ZtreeService } from '../../../../components/ztree';
import { PrivilegeService } from '../../../../service/privilege';

declare var jQuery;

@Component({
  selector: 'execution-result',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./result.scss', '../../../../components/case-comments/comment-edit/src/styles.scss'],
  templateUrl: './result.html'
})
export class ExecutionResult implements OnInit, AfterViewInit, OnDestroy {
  eventCode:string = 'ExecutionResult';

  planId: number;
  runId: number;

  id: number;
  model: any = {};
  settings: any;
  data: any;
  form: any;
  tab: string = 'info';

  fields: any;
  next: boolean = true;

  user: any;
  canEdit: boolean;
  canExe: boolean;

  constructor(private _state:GlobalState, private _routeService: RouteService, private _route: ActivatedRoute, private fb: FormBuilder,
              private _caseService: CaseService, private _caseStepService: CaseStepService, private _caseInRunService: CaseInRunService,
              private _ztreeService: ZtreeService, private privilegeService:PrivilegeService) {

    this.buildForm();
  }
  ngOnInit() {
    this.user = CONSTANT.PROFILE;

    this.canEdit = this.privilegeService.hasPrivilege('cases-update');
    this.canExe = this.privilegeService.hasPrivilege('run-exe');

    this._route.params.forEach((params: Params) => {
      this.planId = +params['planId'];
      this.runId = +params['runId'];
    });

    this._state.subscribe(CONSTANT.EVENT_CASE_EXE, this.eventCode, (data: any) => {
      console.log(CONSTANT.EVENT_CASE_EXE, data);
      let testCase = data.node;
      if (!testCase || testCase.isParent) {
        this.model = {childrenCount: data.childrenCount};
        return;
      }

      this.fields = CONSTANT.CUSTOM_FIELD_FOR_PROJECT;

      if (testCase) {
        this.id = testCase.entityId;

        this.loadData();
      } else {
        this.model = undefined;
      }
    });

    this.settings = {
      canEdit: this.canEdit,
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
    this.form = this.fb.group(
      {
        'result': ['', []],
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
  };

  loadData() {
    let that = this;
    that._caseInRunService.get(that.id).subscribe((json:any) => {
      that.model = json.data;
    });
  }

  setResult(status: string) {
    let next;
    if (this.next) {
      next = this._ztreeService.getNextNode(this.model.id);
    }

    this._caseInRunService.setResult(this.model.entityId, this.model.result, next?next.entityId:null, status).subscribe((json:any) => {
      if (json.code == 1) {
        this.model.status = status;
        this._ztreeService.selectNode(next);

        this._state.notifyDataChanged(CONSTANT.EVENT_CASE_UPDATE, {node: this.model, random: Math.random()});
        this.model = json.data;
      }
    });
  }

  reset() {
    this.loadData();
  }

  saveField (event: any) {
    this._caseService.saveField(this.model.id, event.data).subscribe((json:any) => {
      if (json.code == 1) {
        // this.model = json.data;
        this._state.notifyDataChanged(CONSTANT.EVENT_CASE_UPDATE, {node: this.model, random: Math.random()});
        event.deferred.resolve();
      }
    });
  }

  tabChange(event: any) {
    this.tab = event.nextId;
  }
  changeContentType(contentType: string) {
    this._caseService.changeContentType(contentType, this.model.id).subscribe((json:any) => {
      if (json.code == 1) {
        this.model.contentType = contentType;
      }
    });
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
    this._caseStepService.save(this.model.id, event.newData).subscribe((json:any) => {
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
    let url: string = '/pages/org/' + CONSTANT.CURR_ORG_ID + '/prj/' + CONSTANT.CURR_PRJ_ID + '/implement/plan/' + this.planId + '/view';
    console.log(url);
    this._routeService.navTo(url);
  }

  ngOnDestroy(): void {
    this._state.unsubscribe(CONSTANT.EVENT_CASE_EXE, this.eventCode);
  };

}

