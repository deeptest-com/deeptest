import { Input, Component, OnInit, AfterViewInit, EventEmitter, Output, Inject, Compiler, ViewChild } from '@angular/core';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';

import {NgbModal, NgbModalRef, ModalDismissReasons} from '@ng-bootstrap/ng-bootstrap';

import { CONSTANT } from '../../../utils/constant';
import { Utils, Deferred } from '../../../utils/utils';

import { FieldShowService } from './field-show.service';
import { TinyMCEComponentPopup } from '../../tiny-mce';

@Component({
  selector: 'field-show',
  templateUrl: './field-show.html',
  styleUrls: ['./styles.scss'],
  providers: [FieldShowService]
})
export class FieldShowComponent implements OnInit {

  @Input()
  public model: any = {};
  @Input()
  public field: any = {};

  @Input()
  public prop: string;
  @Input()
  public type: string;
  @Input()
  public format: string;
  @Input()
  public rows: number;
  @Input()
  public optionsKey: string;
  @Input()
  public label: string;
  @Input()
  public required: boolean;
  @Input()
  public readonly: boolean = false;

  @Output() onSave = new EventEmitter<any>();
  public form: any;

  public status: string = 'view';
  public temp: string;
  public casePropertyMap: any = {};

  public richTextEditModal: any;

  public constructor(@Inject(FieldShowService) private fieldShowService: FieldShowService, private fb: FormBuilder,
                     private compiler: Compiler, private modalService: NgbModal) {

  }

  public ngOnInit(): void {
    this.casePropertyMap = CONSTANT.CASE_PROPERTY_MAP;

    this.form = this.fb.group({});
    let control: FormControl = new FormControl(this.prop, Validators.required);
    this.form.addControl(this.prop, control);
  }

  edit(event: any, format?: string) {
    event.preventDefault();
    event.stopPropagation();

    if (format == 'richText' || format == 'rich_text') { // show tinymce in popoup
      this.compiler.clearCacheFor(TinyMCEComponentPopup);

      this.richTextEditModal = this.modalService.open(TinyMCEComponentPopup, {windowClass: 'pop-modal'});
      this.richTextEditModal.componentInstance.content = this.model[this.prop];
      this.richTextEditModal.componentInstance.modelId = this.model['id'];
      this.richTextEditModal.componentInstance.editorKeyup = this.onEditorKeyup;

      this.richTextEditModal.result.then((result) => {
        this.model[this.prop] = result.data;
        this.save();
      }, (reason) => {
        console.log('reason', reason);
      });

    } else if (format == 'planText') { // show textarea in popoup

    } else {
      this.status = 'edit';
      this.temp = this.model[this.prop];
    }
  }

  save(event?: any) {
    if (event) {
      event.preventDefault();
      event.stopPropagation();
    }

    let deferred = new Deferred();
    deferred.promise.then((data) => {
      this.status = 'view';
    }).catch((err) => {console.log('err', err);});

    if (this.model[this.prop] != this.temp) {
      this.onSave.emit({deferred: deferred, data: {prop: this.prop, value: this.model[this.prop]}});
    } else {
      this.cancel(event);
    }
  }
  cancel(event: any) {
    if (event) {
      event.preventDefault();
      event.stopPropagation();
    }

    this.status = 'view';
    this.model[this.prop] = this.temp;
  }

  onEditorKeyup() {
    console.log('onEditorKeyup');
  }

}
