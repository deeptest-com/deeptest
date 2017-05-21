import {Component, ViewEncapsulation, ViewChild, QueryList, Query} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { NgModule, Pipe, OnInit, AfterViewInit }      from '@angular/core';

import {GlobalState} from '../../../../global.state';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import {ValidatorUtils} from '../../../../validator/validator.utils';
import { RouteService } from '../../../../service/route';

import { PopDialogComponent } from '../../../../components/pop-dialog'

import { ProjectService } from '../../../../service/project';

declare var jQuery;

@Component({
  selector: 'project-edit',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./edit.scss'],
  templateUrl: './edit.html'
})
export class ProjectEdit implements OnInit, AfterViewInit {
  type: string;
  id: number;
  model: any = {};
  groups: any[] = [];
  form: FormGroup;
  isSubmitted: boolean;

  @ViewChild('modalWrapper') modalWrapper: PopDialogComponent;

  constructor(private _state:GlobalState, private _routeService: RouteService, private _route: ActivatedRoute,
              private fb: FormBuilder, private _projectService: ProjectService) {
    let that = this;

    this._route.params.subscribe(params => {
      that.type = params['type'];
      that.id = +params['id'];

      that.loadData();
    });

    that.buildForm();
  }
  ngOnInit() {

  }
  ngAfterViewInit() {}

  buildForm(): void {
    let that = this;

    let parentValidate = [];
    if (that.type === 'project') {
      parentValidate = [Validators.required];
    }
    this.form = this.fb.group(
      {
        'name': ['', [Validators.required]],
        'descr': ['', []],
        'parentId': ['', parentValidate],
        'disabled': ['', []]
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
    'parentId': {
      'required':      '项目组不能为空'
    }
  };

  loadData() {
    let that = this;

    that._projectService.get(that.id).subscribe((json:any) => {
      that.groups = json.groups;
      that.model = !!json.data? json.data: {type: that.type, disabled: false};
    });
  }

  save() {
    let that = this;

    that._projectService.save(that.model).subscribe((json:any) => {
      if (json.code == 1) {
        that.model = json.data;

        that.formErrors = ['保存成功'];
        that._routeService.navTo("/pages/project/list");
      } else {
        that.formErrors = ['保存失败'];
      }
    });
  }

  delete() {
    let that = this;

    that._projectService.delete(that.model.id).subscribe((json:any) => {
      if (json.code == 1) {
        that.model = json.data;

        that.formErrors = ['删除成功'];
        that._routeService.navTo("/pages/project/list");
      } else {
        that.formErrors = ['删除失败'];
      }
    });
  }

  showModal(): void {
    this.modalWrapper.showModal(null, '');
  }

}

