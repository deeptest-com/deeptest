import {Component, ViewEncapsulation, NgModule, Pipe, Compiler, OnInit, AfterViewInit, ViewChild} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';

import {NgbDatepickerI18n, NgbDateParserFormatter, NgbDateStruct, NgbModal, NgbModalRef, ModalDismissReasons} from '@ng-bootstrap/ng-bootstrap';

import {I18n, CustomDatepickerI18n} from '../../../../service/datepicker-I18n';

import {GlobalState} from '../../../../global.state';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import {ValidatorUtils} from '../../../../validator/validator.utils';
import { RouteService } from '../../../../service/route';

import { PlanService } from '../../../../service/plan';
import { RunService } from '../../../../service/run';

import { CaseSelectionComponent } from '../../../../components/case-selection'
import { EnvironmentConfigComponent } from '../../../../components/environment-config'
import { PopDialogComponent } from '../../../../components/pop-dialog'

declare var jQuery;

@Component({
  selector: 'plan-view',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./view.scss'],
  templateUrl: './view.html',
  providers: [I18n, {provide: NgbDatepickerI18n, useClass: CustomDatepickerI18n}]
})
export class PlanView implements OnInit, AfterViewInit {
  planId: number;
  model: any = {};
  form: any;

  testSet: any;
  modalTitle: string;

  constructor(private _state:GlobalState, private _routeService: RouteService, private _route: ActivatedRoute, private fb: FormBuilder,
              private _i18n: I18n, private compiler: Compiler,
              private _planService: PlanService, private _runService: RunService) {

  }
  ngOnInit() {
    let that = this;

    that._route.params.forEach((params: Params) => {
      that.planId = +params['planId'];
    });

    if (that.planId) {
      that.loadData();
    }
  }
  ngAfterViewInit() {}

  loadData() {
    let that = this;
    that._planService.get(that.planId).subscribe((json:any) => {
      that.model = json.data;
    });
  }

}

