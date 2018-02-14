import { Component, OnDestroy, AfterViewInit, OnChanges,ViewChild, Input, Output, EventEmitter, Injector, ElementRef } from '@angular/core';

import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';

import {NgbActiveModal} from '@ng-bootstrap/ng-bootstrap';
import { CONSTANT } from '../../../utils/constant';
import {ValidatorUtils} from '../../../validator';
import {CustomFieldOptionService} from "../../../service/custom-field-option";

import * as _ from "lodash";

@Component({
  selector: 'dropdown-options',
  templateUrl: './dropdown-options.html',
  styleUrls: ['./dropdown-options.scss']
})
export class DropdownOptionsComponent implements OnDestroy, AfterViewInit, OnChanges {
  @Input() title: string;
  @Output() confirm = new EventEmitter<any>();

  @Input() field: any;
  @Input() height: number;

  form: FormGroup;
  model: any = {};

  constructor(private fb: FormBuilder, private host: ElementRef, public activeModal: NgbActiveModal,
              private customFieldOptionService: CustomFieldOptionService) {
    this.buildForm();
  }

  dismiss(): any {
    this.activeModal.dismiss({act: 'cancel'});
  }

  ngOnChanges() {

  }

  ngAfterViewInit() {

  }
  ngOnDestroy() {

  }

  edit(item: any) {
    this.model = _.clone(item);
  }
  cancel(item: any) {
    this.form.reset();
    this.model = {};
  }
  save(): any {
    this.customFieldOptionService.save(this.model, this.field.id).subscribe((json:any) => {
      if (json.code == 1) {
        this.form.reset();
        this.model = {};
        this.field.optionVos = json.data;
      }
    });
  }
  delete(item: any) {
    this.customFieldOptionService.delete(this.model, this.field.id).subscribe((json:any) => {
      if (json.code == 1) {
        this.form.reset();
        this.model = {};
        this.field.optionVos = json.data;
      }
    });
  }
  up(item: any) {
    this.customFieldOptionService.changeOrder(item.id, 'up', this.field.id).subscribe((json:any) => {
      if (json.code == 1) {
        this.model = {};
        this.field.optionVos = json.data;
      }
    });
  }
  down(item: any) {
    this.customFieldOptionService.changeOrder(item.id, 'down', this.field.id).subscribe((json:any) => {
      if (json.code == 1) {
        this.model = {};
        this.field.optionVos = json.data;
      }
    });
  }

  buildForm(): void {
    this.form = this.fb.group(
      {
        label: ['', [Validators.required]],
        value: ['', [Validators.required]]
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
    'value': {
      'required':      '取值不能为空'
    },
    'label': {
      'required':      '名称不能为空'
    }
  };

}
