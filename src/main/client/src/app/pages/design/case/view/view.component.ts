import {Component, ViewEncapsulation, NgModule, Pipe, OnInit, AfterViewInit, OnDestroy, ViewChild} from '@angular/core';

import {GlobalState} from '../../../../global.state';

import { CONSTANT } from '../../../../utils/constant';

import { CaseService } from '../../../../service/case';
import { CaseStepService } from '../../../../service/case-step';
import { CaseCommentsService } from '../../../../service/case-comments';

declare var jQuery;

@Component({
  selector: 'case-view',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./view.scss', '../../../../components/case-comments/comment-edit/src/styles.scss'],
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

  casePropertyMap: any;
  fields: any[] = [];
  user: any;

  constructor(private _state:GlobalState,
              private _caseService: CaseService, private _caseStepService: CaseStepService, private _caseCommentsService: CaseCommentsService) {
    this.casePropertyMap = CONSTANT.CASE_PROPERTY_MAP;
  }
  ngOnInit() {
    this.projectId = CONSTANT.CURR_PRJ_ID;
    this.user = CONSTANT.PROFILE;

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

  ngOnDestroy(): void {
    this._state.unsubscribe(CONSTANT.EVENT_CASE_EDIT, this.eventCode);
  };
}

