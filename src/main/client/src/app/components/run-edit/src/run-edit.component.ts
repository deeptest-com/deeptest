import * as _ from "lodash";

import {Component, Input, OnInit, AfterViewInit} from "@angular/core";
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import {NgbActiveModal} from "@ng-bootstrap/ng-bootstrap";

import {CONSTANT} from "../../../utils/constant";

import {UserService} from "../../../service/user";

import {RunEditService} from "./run-edit.service";

@Component({
  selector: 'run-edit',
  templateUrl: './run-edit.html',
  styleUrls: ['./styles.scss']
})
export class RunEditComponent implements OnInit {

  searchModel: any = {};
  searchResult: any[];
  selectedModels: any[] = [];
  model: any = {};

  form: FormGroup;
  createUsers: any[] = [];
  updateUsers: any[] = [];
  _queryModel: any = {type: {}, priority: {}, createUsers: [], updateUsers: []};
  queryModel: any;

  disabled:boolean = false;

  constructor(public activeModal: NgbActiveModal, private fb: FormBuilder,
              public runEditService: RunEditService, public userService: UserService,) {

    this.queryModel = _.cloneDeep(this._queryModel);
  }

  ngOnInit(): any {
    this.buildForm();
  }

  save(): any {
    this.model.userId = this.selectedModels[0].id;
    this.activeModal.close({act: 'save', data: this.model});
  }

  dismiss(): any {
    this.activeModal.dismiss({act: 'cancel'});
  }

  buildForm(): void {
    this.form = this.fb.group(
      {
        'name': ['', [Validators.required]]
      }, {}
    );

    this.form.valueChanges.debounceTime(CONSTANT.DebounceTime).subscribe(data => this.query(data));
  }

  query(data?: any) {

  }

  changeSearch(searchModel: any):void {
    this.userService.search(CONSTANT.CURR_ORG_ID, searchModel.keywords).subscribe((json:any) => {
      if (json.data.length == 0) {
        this.searchResult = null;
      } else {
        this.searchResult = json.data;
      }
    });
  }

  formErrors = [];
  validateMsg = {
    'name': {
      'required':      '名称不能为空'
    }
  };
}
