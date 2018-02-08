import { Input, Output, Component, ViewChild, OnInit, AfterViewInit, OnDestroy } from '@angular/core';

import {NgbModal} from '@ng-bootstrap/ng-bootstrap';

import {CONSTANT} from '../../../../utils/constant';
import {GlobalState} from '../../../../global.state';

import { CaseService } from '../../../../service/case';
import { CaseCommentsService } from '../../../../service/case-comments';
import { CommentEditComponent } from '../../comment-edit/src/comment-edit.component';

@Component({
  selector: 'comment-list',
  styleUrls: ['./styles.scss'],
  templateUrl: './comment-list.html'
})
export class CommentListComponent implements OnInit, AfterViewInit, OnDestroy{
  @Input() @Output() model: any = {};
  userId: number;

  @ViewChild('modalWrapper') modalWrapper: CommentEditComponent;
  comment: any = {};
  eventCode: string = 'CommentListComponent';

  constructor(private _state:GlobalState, private modalService: NgbModal,
              private _caseService: CaseService, private _caseCommentsService: CaseCommentsService) {
    this._state.subscribe(CONSTANT.EVENT_COMMENTS_EDIT, this.eventCode, (json) => {
      console.log(CONSTANT.EVENT_COMMENTS_EDIT + ' in ' + this.eventCode, json);
      this.addComments(json);
    });

    this._state.subscribe(CONSTANT.EVENT_COMMENTS_SAVE, this.eventCode, (json) => {
      console.log(CONSTANT.EVENT_COMMENTS_SAVE + ' in ' + this.eventCode, json);
      this.comment = json;
      this.saveComments(json.pass);
    });
  }

  public ngOnInit(): void {
    this.userId = CONSTANT.PROFILE.id;
  }

  ngAfterViewInit() {

  }

  addComments(data?: any) {
    this.modalWrapper.showModal('comment-edit');
    if (data) {
      this.comment = data;
    } else {
      this.comment = {summary: '添加备注'};
    }
  }
  editComments(comment: any) {
    this.modalWrapper.showModal('comment-edit');
    this.comment = comment;
    if (this.comment.summary === '添加备注') {
      this.comment.summary = '修改备注';
    }
  }
  removeComments(id: number, indx: number) {
    this._caseCommentsService.remove(id).subscribe((json:any) => {
      this.model.comments.splice(indx, 1);
    });
  }

  saveComments(pass:boolean) {
    this._caseCommentsService.save(this.model.id, this.comment).subscribe((json:any) => {
      if (json.code == 1) {
        if (this.comment.pass != undefined) { // 评审
          this.reviewRequest(this.model.id, this.comment.pass);
        }

        if (this.comment.id != json.data.id) {
          this.model.comments[this.model.comments.length] = json.data;
        }
        this.comment = json.data;
        if (!pass) {
          this.modalWrapper.closeModal();
        }
      }
    });
  }

  reviewRequest(id: number, pass: boolean) {
    this._caseService.reviewPass(id, pass).subscribe((json:any) => {
      if (json.code == 1) {
        this.model.reviewResult = pass;
        this._state.notifyDataChanged(CONSTANT.EVENT_CASE_UPDATE, {node: this.model, random: Math.random()});
      }
    });
  }

  ngOnDestroy(): void {
    this._state.unsubscribe(CONSTANT.EVENT_COMMENTS_EDIT, this.eventCode);
    this._state.unsubscribe(CONSTANT.EVENT_COMMENTS_SAVE, this.eventCode);
  };

}
