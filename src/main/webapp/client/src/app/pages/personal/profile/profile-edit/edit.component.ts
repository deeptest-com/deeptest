import { Component,ViewEncapsulation, Pipe, OnInit, AfterViewInit, ViewChild } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Router, ActivatedRoute, Params } from '@angular/router';
import {NgbModal, NgbModalRef, ModalDismissReasons} from '@ng-bootstrap/ng-bootstrap';

import {GlobalState} from '../../../../global.state';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import {ValidatorUtils} from '../../../../validator/validator.utils';

import { RouteService } from '../../../../service/route';

import {AccountService} from './../../../../service/account';

import { PasswordEditPopupComponent } from '../../password';

@Component({
  selector: 'profile-edit-property',
  styleUrls: ['./edit.scss'],
  templateUrl: './edit.html'
})
export class ProfileEdit implements OnInit, AfterViewInit {
  model: any = {};
  recentProjects: any[] = [];
  currProject: any = {};

  orgs: any[] = [];
  orgId: number;

  public passwordPopop: any;

  constructor(private _routeService: RouteService, private _state:GlobalState, private _route: ActivatedRoute,
      private modalService: NgbModal, private accountService: AccountService) {

    this._state.subscribe(CONSTANT.STATE_CHANGE_PROFILE, (profile) => {
      console.log(CONSTANT.STATE_CHANGE_PROFILE, profile);
      this.model = profile;
    });
    this._state.subscribe(CONSTANT.STATE_CHANGE_PROJECTS, (data) => {
      console.log(CONSTANT.STATE_CHANGE_PROJECTS, data);
      this.recentProjects = data.recentProjects;
      this.currProject = {id: CONSTANT.CURR_PRJ_ID, name: CONSTANT.CURR_PRJ_NAME};
    });
    this._state.subscribe(CONSTANT.STATE_CHANGE_ORGS, (data) => {
      console.log(CONSTANT.STATE_CHANGE_ORGS, data);
      if (data.currOrgId) {
        this.orgId = data.currOrgId;
      }
      if (data.orgs) {
        this.orgs = data.orgs;
      }
    });
  }

  ngOnInit() {
    this.loadData();
  }

  ngAfterViewInit() {
  }

  saveField(event: any):void {
    this.accountService.saveInfo(event.data).subscribe((json:any) => {
      if (json.code == 1) {
        this.model = json.data;
        this.accountService.changeProfile(this.model);

        event.deferred.resolve();
      }
    });
  }

  loadData() {
    if(CONSTANT.PROFILE) {
      this.model = CONSTANT.PROFILE;

      this.orgs = CONSTANT.ALL_ORGS;
      this.orgId = CONSTANT.CURR_ORG_ID;

      this.recentProjects =  CONSTANT.RECENT_PROJECTS;
      this.currProject = {id: CONSTANT.CURR_PRJ_ID, name: CONSTANT.CURR_PRJ_NAME};
    }
  }

  uploadedEvent(event: any) {
    console.log('uploadedEvent', event);

    this.accountService.saveInfo({prop: 'avatar', value: event.data}).subscribe((json:any) => {
      if (json.code == 1) {
        this.model = json.data;
        this.accountService.changeProfile(this.model);

        event.deferred.resolve();
      }
    });
  }

  editPassword(event: any) {
    this.passwordPopop = this.modalService.open(PasswordEditPopupComponent, {windowClass: ''});
    this.passwordPopop.result.then((result) => {

    }, (reason) => {
      console.log('reason', reason);
    });
  }

}
