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
import { UserAndGroupService } from '../../../../service/userAndGroup';

declare var jQuery;

@Component({
  selector: 'project-edit',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./edit.scss'],
  templateUrl: './edit.html'
})
export class ProjectEdit implements OnInit, AfterViewInit {
  orgId: number;
  type: string;
  id: number;

  formInfo: FormGroup;
  formAdd: FormGroup;

  tab: string = 'info';

  groups: any[] = [];
  projectRoles: any[] = [];

  entityInRoles: any[] = [];
  entitySearchResult: any[];

  selectedModels: any[] = [];
  modelAdd: any = {roleId: 1};
  model: any = {};
  searchModel: any = {};

  @ViewChild('modalWrapper') modalWrapper: PopDialogComponent;

  constructor(private _state:GlobalState, private _routeService: RouteService, private _route: ActivatedRoute,
              private fb: FormBuilder, private _projectService: ProjectService, private _userAndGroupService: UserAndGroupService) {
    let that = this;
    this.orgId = CONSTANT.CURR_ORG_ID;

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
    this.formInfo = this.fb.group(
      {
        'name': ['', [Validators.required]],
        'descr': ['', []],
        'parentId': ['', parentValidate],
        'disabled': ['', []]
      }, {}
    );

    this.formInfo.valueChanges.debounceTime(500).subscribe(data => this.onValueChanged(data));
    this.onValueChanged();

    this.formAdd = this.fb.group(
      {
        'projectRole': ['', [Validators.required]]
      }, {}
    );

  }
  onValueChanged(data?: any) {
    let that = this;
    that.formErrors = ValidatorUtils.genMsg(that.formInfo, that.validateMsg, []);
  }

  formErrors = [];
  validateMsg = {
    'name': {
      'required':      '姓名不能为空'
    },
    // 'parentId': {
    //   'required':      '项目组不能为空'
    // }
  };

  loadData() {
    let that = this;

    that._projectService.get(that.id).subscribe((json:any) => {
      that.projectRoles = json.projectRoles;
      that.groups = json.groups;
      that.entityInRoles = json.entityInRoles;

      that.model = json.data? json.data: {type: that.type, disabled: false};
    });
  }

  save() {
    let that = this;

    that._projectService.save(that.model).subscribe((json:any) => {
      if (json.code == 1) {
        that.formErrors = ['保存成功'];
        that._routeService.navTo('/pages/org/' + CONSTANT.CURR_ORG_ID + '/prjs');
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
        that._routeService.navTo('/pages/org/' + CONSTANT.CURR_ORG_ID + '/prjs');

        this.modalWrapper.closeModal();
      } else {
        that.formErrors = ['删除失败'];
      }
    });
  }

  showModal(): void {
    this.modalWrapper.showModal();
  }

  tabChange(event: any) {
    this.tab = event.nextId;
  }

  add($event) {
    $event.preventDefault();
    $event.stopPropagation();

    this.modelAdd.projectId = this.id;

    let entityTypeAndIds:string[] = [];
    this.selectedModels.forEach(item => {entityTypeAndIds.push(item.type + ',' + item.id)})

    this._projectService.saveMembers(this.modelAdd, entityTypeAndIds).subscribe((json:any) => {
      if (json.code == 1) {
        this.searchModel = {};
        this.modelAdd = {roleId: 1};
        this.selectedModels = [];
        this.entityInRoles = json.entityInRoles;
      }
    });
  }

  changeSearch(searchModel: any):void {
    let ids = [];
    this.selectedModels.forEach(item => {ids.push(item.id)})

    this._userAndGroupService.search(this.model.orgId, searchModel.keywords, ids).subscribe((json:any) => {
      if (json.data.length == 0) {
        this.entitySearchResult = null;
      } else {
        this.entitySearchResult = json.data;
      }
    });
  }

  changeRole(roleId: number, entityId: number) {
    this._projectService.changeRole(this.model.id, roleId, entityId).subscribe((json:any) => {
      if (json.code == 1) {
        this.entityInRoles = json.entityInRoles;
      }
    });
  }

}

