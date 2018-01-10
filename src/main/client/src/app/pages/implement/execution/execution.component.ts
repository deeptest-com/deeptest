import {Component, ViewEncapsulation, NgModule, Pipe, OnInit, AfterViewInit} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';

import {GlobalState} from '../../../global.state';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';
import {ValidatorUtils} from '../../../validator/validator.utils';
import { RouteService } from '../../../service/route';

declare var jQuery;

@Component({
  selector: 'execution',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./execution.scss'],
  templateUrl: './execution.html'
})
export class Execution implements OnInit, AfterViewInit {
  contentHeight = Utils.getContainerHeight(110);
  leftWidth: number;

  constructor(private _state: GlobalState, private _route: ActivatedRoute) {
    this._state.subscribe(CONSTANT.STATE_CHANGE_PROFILE, (profile) => {
      console.log(CONSTANT.STATE_CHANGE_PROFILE + ' in Exe', profile);
      this.leftWidth = CONSTANT.PROFILE.leftSize;
    });

    if (CONSTANT.PROFILE) {
      this.leftWidth = CONSTANT.PROFILE.leftSize;
    }

  }

  ngOnInit() {


  }
  ngAfterViewInit() {}


}

