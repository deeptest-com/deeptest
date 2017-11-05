import { Component,ViewEncapsulation, Pipe, OnInit, AfterViewInit, ViewChild } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Router, ActivatedRoute, Params } from '@angular/router';
import {NgbModal, NgbModalRef, ModalDismissReasons} from '@ng-bootstrap/ng-bootstrap';

import {GlobalState} from '../../../../global.state';
import {EmailValidator} from '../../../../validator';

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
  currOrgId: number;

  public passwordPopop: any;

  constructor(private _routeService: RouteService, private _state:GlobalState, private _route: ActivatedRoute,
      private modalService: NgbModal, private accountService: AccountService) {

    this._state.subscribe('profile.refresh', (profile) => {
      console.log('profile.refresh', profile);
      this.model = profile;
    });
    this._state.subscribe('recent.projects.change', (data) => {
      console.log('recent.projects.change', data);
      this.recentProjects = data.recentProjects;
      this.currProject = data.currProject;
    });
    this._state.subscribe('my.orgs.change', (data) => {
      console.log('recent.projects.change', data);
      this.orgs = data.orgs;
      this.currOrgId = data.currOrgId;
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
      this.currOrgId = CONSTANT.CURR_ORG_ID;

      this.recentProjects =  CONSTANT.RECENT_PROJECTS;
      this.currProject = CONSTANT.CURRENT_PROJECT;

      console.log(this.orgs, this.recentProjects);
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
    console.log('===');
    this.passwordPopop = this.modalService.open(PasswordEditPopupComponent, {windowClass: ''});
    this.passwordPopop.result.then((result) => {

    }, (reason) => {
      console.log('reason', reason);
    });
  }

}
