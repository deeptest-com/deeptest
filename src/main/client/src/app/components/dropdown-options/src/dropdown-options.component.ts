import { Component, OnDestroy, AfterViewInit, OnChanges,ViewChild, Input, Output, EventEmitter, Injector, ElementRef } from '@angular/core';

import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';

import {NgbActiveModal} from '@ng-bootstrap/ng-bootstrap';
import { CONSTANT } from '../../../utils/constant';
import {ValidatorUtils} from '../../../validator';

import 'tinymce';

declare var tinymce: any;

@Component({
  selector: 'dropdown-options',
  templateUrl: './dropdown-options.html',
  styleUrls: ['./dropdown-options.scss']
})
export class DropdownOptionsComponent implements OnDestroy, AfterViewInit, OnChanges {
  @Input() title: string;
  @Output() confirm = new EventEmitter<any>();

  @Input() options: any;
  @Input() height: number;

  form: FormGroup;
  model: any = {};

  constructor(private fb: FormBuilder, private host: ElementRef, public activeModal: NgbActiveModal) {
    this.buildForm();
  }

  save(): any {
    this.activeModal.close({act: 'save', data: this.options});
  }
  dismiss(): any {
    this.activeModal.dismiss({act: 'cancel'});
  }

  ngOnChanges() {
    let editor = tinymce.get("mceEditor");
    if (editor) {editor.setContent(this.options);}
  }

  ngAfterViewInit() {

  }
  ngOnDestroy() {

  }

  onEditorKeyup(event: any) {
    this.options = event;
  }

  edit(item: any) {
    console.log(item);
  }
  delete(item: any) {
    console.log(item);
  }
  up(item: any) {
    console.log(item);
  }
  down(item: any) {
    console.log(item);
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
      'required':      '值不能为空'
    },
    'label': {
      'required':      '名称不能为空'
    }
  };

}
