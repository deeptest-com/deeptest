import {Component, ViewEncapsulation, NgModule, Pipe, OnInit, AfterViewInit, OnDestroy, ViewChild} from '@angular/core';
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
  selector: 'case-view',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./view.scss', '../../../../components/comment-edit/src/styles.scss'],
  templateUrl: './view.html'
})
export class CaseView implements OnInit, AfterViewInit, OnDestroy {
  eventCode:string = 'CaseView';

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

  casePropertyMap: any;
  fields: any[] = [];

  constructor(private _state:GlobalState, private fb: FormBuilder, private toastyService:ToastyService,
              private _caseService: CaseService, private _caseStepService: CaseStepService, private _caseCommentsService: CaseCommentsService) {
    this.casePropertyMap = CONSTANT.CASE_PROPERTY_MAP;
  }
  ngOnInit() {
    this.projectId = CONSTANT.CURR_PRJ_ID;

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
      canEdit: false,
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

  loadData() {
    this._caseService.get(this.id).subscribe((json:any) => {
      this.model = json.data;
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

