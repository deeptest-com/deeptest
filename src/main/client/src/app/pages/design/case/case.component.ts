import {Component, Directive, ElementRef, Inject, Renderer2, OnDestroy, OnInit, AfterViewInit} from "@angular/core";
import { Router, ActivatedRoute, Params } from '@angular/router';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';

import {GlobalState} from "../../../global.state";

@Component({
  selector: 'case',
  styleUrls: ['./case.scss'],
  templateUrl: './case.html'
})
export class Case implements OnInit, AfterViewInit {
  projectId: number;
  key: number;

  contentHeight = Utils.getContainerHeight(110);
  leftWidth: number;

  constructor(private _state: GlobalState, private _route: ActivatedRoute) {
    this._state.subscribe(CONSTANT.STATE_CHANGE_PROFILE, (profile) => {
      console.log(CONSTANT.STATE_CHANGE_PROFILE + ' in Case', profile);
      this.leftWidth = CONSTANT.PROFILE.leftSize;
    });

    if (CONSTANT.PROFILE) {
      this.leftWidth = CONSTANT.PROFILE.leftSize;
    }

  }

  ngOnInit() {

  }

  ngAfterViewInit() {

  }

}
