import {Component, ViewEncapsulation, ViewChild} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import { NgModule, Pipe, OnInit, AfterViewInit }      from '@angular/core';

import { NgbModalModule, NgbPaginationModule, NgbDropdownModule,
  NgbTabsetModule, NgbButtonsModule, NgbCollapseModule } from '@ng-bootstrap/ng-bootstrap';
import { BrowserModule } from '@angular/platform-browser';

import {GlobalState} from '../../../../global.state';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import {ValidatorUtils} from '../../../../validator/validator.utils';
import { RouteService } from '../../../../service/route';

import { ProjectRoleService } from '../../../../service/project-role';
import { PopDialogComponent } from '../../../../components/pop-dialog'

declare var jQuery;

@Component({
  selector: 'project-role-edit',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./edit.scss'],
  templateUrl: './edit.html'
})
export class ProjectRoleEdit implements OnInit, AfterViewInit {
  id: number;
  tab: string = 'info';
  projectRole: any = {disabled: false};
  projectPrivileges: any[] = [];
  form: any;
  @ViewChild('modalWrapper') modalWrapper: PopDialogComponent;

  constructor(private _state:GlobalState, private _routeService: RouteService, private _route: ActivatedRoute,
              private fb: FormBuilder, private projectRoleService: ProjectRoleService) {

  }
  ngOnInit() {
    let that = this;

    that._route.params.forEach((params: Params) => {
      that.id = +params['id'];
    });

    that.loadData();
    that.buildForm();
  }
  ngAfterViewInit() {}

  loadData() {
    let that = this;
    that.projectRoleService.get(that.id).subscribe((json:any) => {
      that.projectRole = json.projectRole;
      that.projectPrivileges = json.projectPrivileges;
    });
  }

  save() {
    let that = this;

    that.projectRoleService.save(that.projectRole, that.projectPrivileges).subscribe((json:any) => {
      if (json.code == 1) {

        that.formErrors = ['保存成功'];
        that._routeService.navTo("/pages/org-admin/project-role/list");
      } else {
        that.formErrors = ['保存失败'];
      }
    });
  }

  delete() {
    let that = this;

    that.projectRoleService.delete(that.projectRole.id).subscribe((json:any) => {
      if (json.code == 1) {
        that.formErrors = ['删除成功'];
        that._routeService.navTo("/pages/org-admin/project-role/list");

        this.modalWrapper.closeModal();
      } else {
        that.formErrors = [json.msg];
      }
    });
  }

  select(key: string) {
    let val = key ==='all'? true: false;
    for (let user of this.projectPrivileges) {
      user.selecting = val;
    }
  }
  tabChange(event: any) {
    this.tab = event.nextId;
  }

  buildForm(): void {
    this.form = this.fb.group(
      {
        'name': ['', [Validators.required]],
        'descr': ['', []],
        'disabled': ['', []]
      }, {}
    );

    this.form.valueChanges.debounceTime(CONSTANT.DebounceTime).subscribe(data => this.onValueChanged(data));
    this.onValueChanged();
  }
  onValueChanged(data?: any) {
    let that = this;
    that.formErrors = ValidatorUtils.genMsg(that.form, that.validateMsg, []);
  }

  formErrors = [];
  validateMsg = {
    'name': {
      'required':      '名称不能为空'
    },
    'descr': {
    }
  };

  showModal(): void {
    this.modalWrapper.showModal();
  }

}

