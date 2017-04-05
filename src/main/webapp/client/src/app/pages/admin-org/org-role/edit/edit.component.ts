import {Component, ViewEncapsulation, ViewChild} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import { NgModule, Pipe, OnInit, AfterViewInit }      from '@angular/core';
import { ModalDirective } from 'ng2-bootstrap';
import {GlobalState} from '../../../../global.state';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import {ValidatorUtils} from '../../../../validator/validator.utils';
import { RouteService } from '../../../../service/route';

import { OrgRoleService } from '../../../../service/org-role';
import { PopDialogComponent } from '../../../../components/pop-dialog'

declare var jQuery;

@Component({
  selector: 'org-role-edit',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./edit.scss')],
  template: require('./edit.html')
})
export class OrgRoleEdit implements OnInit, AfterViewInit {
  id: number;
  tab: string = 'info';
  orgRole: any = {disabled: false};
  orgPriviledges: any[] = [];
  form: any;
  @ViewChild('modalWrapper') modalWrapper: PopDialogComponent;

  constructor(private _state:GlobalState, private _routeService: RouteService, private _route: ActivatedRoute,
              private fb: FormBuilder, private orgRoleService: OrgRoleService) {

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
    that.orgRoleService.get(that.id).subscribe((json:any) => {
      that.orgRole = json.orgRole;
      that.orgPriviledges = json.orgPriviledges;

      _.forEach(that.orgPriviledges, (priviledge: any, index: number) => {
        this.form.addControl('priviledge-' + priviledge.id, new FormControl('', []))
      });
    });
  }

  save() {
    let that = this;

    that.orgRoleService.save(that.orgRole, that.orgPriviledges).subscribe((json:any) => {
      if (json.code == 1) {

        that.formErrors = ['保存成功'];
        that._routeService.navTo("/pages/org-admin/org-role/list");
      } else {
        that.formErrors = ['保存失败'];
      }
    });
  }
  
  reset() {
    this.loadData();
  }

  delete() {
    let that = this;

    that.orgRoleService.delete(that.orgRole.id).subscribe((json:any) => {
      if (json.code == 1) {
        that.orgRole = json.data;

        that.formErrors = ['删除成功'];
        that._routeService.navTo("/pages/org-project/list");
      } else {
        that.formErrors = ['删除失败'];
      }
    });
  }

  select(key: string) {
    let val = key ==='all'? true: false;
    for (let user of this.orgPriviledges) {
      user.selecting = val;
    }
  }
  selectTab(tab: string) {
    let that = this;
    that.tab = tab;
  }

  buildForm(): void {
    let that = this;
    this.form = this.fb.group(
      {
        'name': ['', [Validators.required]],
        'descr': ['', []],
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
    'descr': {
    }
  };

  showModal(): void {
    this.modalWrapper.showModal();
  }

}

