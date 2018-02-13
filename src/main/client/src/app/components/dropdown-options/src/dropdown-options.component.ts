import { Component, OnDestroy, AfterViewInit, OnChanges,ViewChild, Input, Output, EventEmitter, Injector, ElementRef } from '@angular/core';
import {NgbModal, NgbModalRef, NgbActiveModal, ModalDismissReasons} from '@ng-bootstrap/ng-bootstrap';

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

  constructor(private host: ElementRef, public activeModal: NgbActiveModal) {}

  save(): any {
    this.removeTinymce();
    this.activeModal.close({act: 'save', data: this.options});
  }
  dismiss(): any {
    this.removeTinymce();
    this.activeModal.dismiss({act: 'cancel'});
  }

  ngOnChanges() {
    let editor = tinymce.get("mceEditor");
    if (editor) {editor.setContent(this.options);}
  }

  ngAfterViewInit() {

  }
  ngOnDestroy() {
    this.removeTinymce();
  }

  onEditorKeyup(event: any) {
    this.options = event;
  }

  removeTinymce() {
    if (tinymce.get("mceEditor")) {
      tinymce.get("mceEditor").remove();
    }
  }

}
