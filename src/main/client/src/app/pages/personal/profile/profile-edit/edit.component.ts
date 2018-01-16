import { Component,ViewEncapsulation, Pipe, OnInit, AfterViewInit, OnDestroy, ViewChild } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Router, ActivatedRoute, Params } from '@angular/router';
import {NgbModal, NgbModalRef, ModalDismissReasons} from '@ng-bootstrap/ng-bootstrap';

import {GlobalState} from '../../../../global.state';

import { CONSTANT } from '../../../../utils/constant';

import { RouteService } from '../../../../service/route';

import {UserService} from './../../../../service/user';

import { PasswordEditPopupComponent } from '../../password';

@Component({
  selector: 'profile-edit-property',
  styleUrls: ['./edit.scss'],
  templateUrl: './edit.html'
})
export class ProfileEdit implements OnInit, AfterViewInit, OnDestroy {
  eventCode: string = 'ProfileEdit';

  model: any = {};
  recentProjects: any[] = [];
  currProject: any = {};

  orgs: any[] = [];
  orgId: number;

  public passwordPopop: any;

  constructor(private _routeService: RouteService, private _state:GlobalState, private _route: ActivatedRoute,
      private modalService: NgbModal, private userService: UserService) {

  }

  ngOnInit() {
    this.loadData();
  }

  ngAfterViewInit() {
  }

  saveField(event: any):void {
    this.userService.saveInfo(event.data).subscribe((json:any) => {
      if (json.code == 1) {
        this.model = json.data;
        event.deferred.resolve();
      }
    });
  }

  loadData() {
    if(CONSTANT.PROFILE) {
      this.model = CONSTANT.PROFILE;

      this.orgs = CONSTANT.MY_ORGS;
      this.orgId = CONSTANT.CURR_ORG_ID;

      this.recentProjects =  CONSTANT.RECENT_PROJECTS;
      this.currProject = {id: CONSTANT.CURR_PRJ_ID, name: CONSTANT.CURR_PRJ_NAME};
    }
  }

  uploadedEvent(event: any) {
    console.log('uploadedEvent', event);

    this.userService.saveInfo({prop: 'avatar', value: event.data}).subscribe((json:any) => {
      if (json.code == 1) {
        this.model = json.data;

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

  ngOnDestroy(): void {

  };
}
