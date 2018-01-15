import {Component, Directive, ElementRef, Inject, Renderer2, OnDestroy, OnInit, AfterViewInit} from "@angular/core";
import { Router, ActivatedRoute, Params } from '@angular/router';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';

import {GlobalState} from "../../../global.state";

import { PrivilegeService } from '../../../service/privilege';

@Component({
  selector: 'case',
  styleUrls: ['./case.scss'],
  templateUrl: './case.html'
})
export class Case implements OnInit, AfterViewInit, OnDestroy {
  eventCode:string = 'Case';

  projectId: number;
  key: number;

  contentHeight = Utils.getContainerHeight(110);
  leftWidth: number;
  canEdit: boolean;

  constructor(private _state: GlobalState, private _route: ActivatedRoute, private privilegeService:PrivilegeService) {

  }

  ngOnInit() {
      this.leftWidth = CONSTANT.PROFILE.leftSize;
      this.canEdit = this.privilegeService.hasPrivilege('cases-update');
  }

  ngAfterViewInit() {

  }

  ngOnDestroy(): void {
  };

}
