import {Component, ViewEncapsulation} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { FileUploader, FileUploaderOptions } from 'ng2-file-upload';
import { NgModule, Pipe, OnInit, AfterViewInit }      from '@angular/core';

import {GlobalState} from '../../../../global.state';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import {ValidatorUtils} from '../../../../validator/validator.utils';
import { RouteService } from '../../../../service/route';

import { ProjectService } from '../../../../service/project';

declare var jQuery;

@Component({
  selector: 'project-edit',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./edit.scss')],
  template: require('./edit.html')
})
export class ProjectEdit implements OnInit, AfterViewInit {
  id: number;
  model: any = {};
  form: any;
  isSubmitted: boolean;

  constructor(private _state:GlobalState, private _routeService: RouteService, private _route: ActivatedRoute,
              private fb: FormBuilder, private _projectService: ProjectService) {

  }
  ngOnInit() {
    let that = this;

    that._route.params.forEach((params: Params) => {
      that.id = +params['id'];
    });

    if (that.id) {
      that.loadData();
    }
    that.buildForm();
  }
  ngAfterViewInit() {}

  buildForm(): void {
    let that = this;
    this.form = this.fb.group(
      {
        'name': [that.model.name, [Validators.required]],
        'title': [that.model.title, [Validators.required]],
        'descr': [that.model.descr, [Validators.required]]
      }, {}
    );

    this.form.valueChanges.subscribe(data => this.onValueChanged(data));
    this.onValueChanged();
  }
  onValueChanged(data?: any) {
    let that = this;
    that.formErrors = ValidatorUtils.genMsg(that.form, that.validateMsg, []);
  }

  formErrors = [];
  validateMsg = {
    'name': {
      'required':      '姓名不能为空'
    },
    'title': {
      'required':      '简介不能为空'
    },
    'descr': {
      'required':      '描述不能为空'
    }
  };

  loadData() {
    let that = this;
    that._projectService.get(that.id).subscribe((json:any) => {
      that.model = json.data;
    });
  }

  save() {
    let that = this;

    that._projectService.save(that.model).subscribe((json:any) => {
      if (json.code == 1) {
        that.model = json.data;
      }
    });
  }

}

