import {Component, ViewEncapsulation, NgModule, Pipe, OnInit, AfterViewInit} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';

import {GlobalState} from '../../../global.state';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';
import {ValidatorUtils} from '../../../validator/validator.utils';
import { RouteService } from '../../../service/route';

import { RunService } from '../../../service/run';

declare var jQuery;

@Component({
  selector: 'execution',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./execution.scss'],
  templateUrl: './execution.html'
})
export class Execution implements OnInit, AfterViewInit {
  contentHeight = Utils.getContainerHeight(110);

  constructor(private _state:GlobalState, private _routeService: RouteService, private _route: ActivatedRoute, private fb: FormBuilder,
              private _runService: RunService) {

  }
  ngOnInit() {


  }
  ngAfterViewInit() {}


}

