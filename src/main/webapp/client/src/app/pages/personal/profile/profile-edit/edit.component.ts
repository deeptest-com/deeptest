import { Component,ViewEncapsulation, Pipe, OnInit, AfterViewInit, ViewChild } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Router, ActivatedRoute, Params } from '@angular/router';

import {GlobalState} from '../../../../global.state';
import {EmailValidator} from '../../../../validator';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import {ValidatorUtils} from '../../../../validator/validator.utils';

import { RouteService } from '../../../../service/route';

import {AccountService} from './../../../../service/account';

@Component({
  selector: 'profile-edit-property',
  styleUrls: ['./edit.scss'],
  templateUrl: './edit.html'
})
export class ProfileEdit implements OnInit, AfterViewInit {
  model: any = {};
  recentProjects: any[] = [];
  orgs: any[] = [];

  uploaderOptions: any = {url: CONSTANT.UPLOAD_URI};

  constructor(private _routeService: RouteService, private _state:GlobalState, private _route: ActivatedRoute,
              private accountService: AccountService) {

    this._state.subscribe('profile.refresh', (profile) => {
      console.log('profile.refresh', profile);
      this.model = profile;
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
      this.recentProjects =  CONSTANT.RECENT_PROJECTS;
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

}
