import {Component, ViewEncapsulation, NgModule, Pipe, OnInit, AfterViewInit, OnDestroy,ViewChild} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';

import {ToastyService, ToastyConfig, ToastOptions, ToastData} from 'ng2-toasty';

import {GlobalState} from '../../../../global.state';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import {ValidatorUtils} from '../../../../validator/validator.utils';
import { RouteService } from '../../../../service/route';

import { CaseService } from '../../../../service/case';
import { CaseStepService } from '../../../../service/case-step';
import { CaseCommentsService } from '../../../../service/case-comments';

import { CommentEditComponent } from '../../../../components/comment-edit';

declare var jQuery;

@Component({
  selector: 'case-edit',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./edit.scss', '../../../../components/comment-edit/src/styles.scss'],
  templateUrl: './edit.html'
})
export class CaseEdit implements OnInit, AfterViewInit, OnDestroy {
  eventCode: string = 'CaseEdit';

  projectId: number;
  id: number;
  model: any = {};
  isModule: true;

  settings: any;
  data: any;
  form: any;
  tab: string = 'content';

  @ViewChild('modalWrapper') modalWrapper: CommentEditComponent;
  comment: any = {};

  fields: any[] = [];
  user: any;

  constructor(private _state:GlobalState, private fb: FormBuilder, private toastyService:ToastyService,
              private _caseService: CaseService, private _caseStepService: CaseStepService, private _caseCommentsService: CaseCommentsService) {

  }
  ngOnInit() {
    this.projectId = CONSTANT.CURR_PRJ_ID;
    this.user = CONSTANT.PROFILE;

    this.buildForm();

    this._state.subscribe(CONSTANT.EVENT_CASE_EDIT, this.eventCode, (data: any) => {
      let testCase = data.node;

      if (!testCase || testCase.isParent) {
        this.model = {};
        return;
      }

      this.fields = CONSTANT.CUSTOM_FIELD_FOR_PROJECT;

      if (testCase) {
        this.id = testCase.id;
        this.loadData();
      } else {
        this.model = null;
      }
    });

    this.settings = {
      canEdit: true,
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
        'name': ['', [Validators.required]],
        'type': ['', [Validators.required]],
        'priority': ['', [Validators.required]],
        'estimate': ['', []],
        'objective': ['', []],
        'pre_condition': ['', []]
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
      'required':      '标题不能为空'
    },
    'type': {
      'required':      '类别不能为空'
    },
    'priority': {
      'required':      '优先级不能为空'
    }
  };

  loadData() {
    this._caseService.get(this.id).subscribe((json:any) => {
      this.model = json.data;
    });
  }

  save() {
    this._caseService.save(this.projectId, this.model).subscribe((json:any) => {
      if (json.code == 1) {
        this.model = json.data;
        this._state.notifyDataChanged(CONSTANT.EVENT_CASE_UPDATE, {node: this.model, random: Math.random()});

        var toastOptions:ToastOptions = {
          title: "保存成功",
          timeout: 2000
        };
        this.toastyService.success(toastOptions);
      }
    });
  }

  review(pass: boolean) {
    if (pass) {
      this.reviewRequest(this.model.id, pass, null);
    } else {
      this.modalWrapper.showModal('comment-edit');
      this.comment = {summary: '评审失败', act: 'reviewFail'};
    }
  }
  reviewRequest(id: number, pass: boolean, comments: string) {
    this._caseService.reviewPass(this.model.id, pass, comments).subscribe((json:any) => {
      if (json.code == 1) {
        this.model = json.data;
        this._state.notifyDataChanged(CONSTANT.EVENT_CASE_UPDATE, {node: this.model, random: Math.random()});
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
    this._caseStepService.up({caseId: this.id, id: event.data.id, ordr: event.data.ordr}).subscribe((json:any) => {
      event.confirm.resolve();
    });
  }

  onDownConfirm(event: any) {
    console.log('onDownConfirm', event);
    this._caseStepService.down({caseId: this.id, id: event.data.id, ordr: event.data.ordr}).subscribe((json:any) => {
      event.confirm.resolve();
    });
  }

  onCreateConfirm(event: any) {
    console.log('onCreateConfirm', event);
    event.confirm.resolve();
  }
  onSaveConfirm(event: any) {
    console.log('onSaveConfirm', event);
    this._caseStepService.save(this.id, event.newData).subscribe((json:any) => {
      event.confirm.resolve();
    });
  }
  onDeleteConfirm(event: any) {
    console.log('onDeleteConfirm', event);
    this._caseStepService.delete(event.data).subscribe((json:any) => {
      event.confirm.resolve();
    });
  }

  onEditorKeyup(event: any) {
    this.model.content = event;
  }

  addComments() {
    this.modalWrapper.showModal('comment-edit');
    this.comment = {summary: '添加备注'};
  }
  editComments(comment: any) {
    this.modalWrapper.showModal('comment-edit');
    this.comment = comment;
    if (this.comment.summary === '添加备注') {
      this.comment.summary = '修改备注';
    }
  }
  removeComments(id: number, indx: number) {
    console.log('remove', id);
    this._caseCommentsService.remove(id).subscribe((json:any) => {
      this.model.comments.splice(indx, 1);
    });
  }

  saveComments() {
    this._caseCommentsService.save(this.id, this.comment).subscribe((json:any) => {
      if (json.code == 1) {
        if (this.comment.act == 'reviewFail') {
          this.reviewRequest(this.model.id, false, null);
        }

        if (this.comment.id != json.data.id) {
          this.model.comments[this.model.comments.length] = json.data;
        }
        this.comment = json.data;
        this.modalWrapper.closeModal();
      }
    });
  }

  ngOnDestroy(): void {
    this._state.unsubscribe(CONSTANT.EVENT_CASE_EDIT, this.eventCode);
  };

}

