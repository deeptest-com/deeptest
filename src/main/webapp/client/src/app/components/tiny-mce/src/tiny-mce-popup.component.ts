import { Component, OnDestroy, AfterViewInit, OnChanges,ViewChild, Input, Output, EventEmitter, Injector, ElementRef } from '@angular/core';
import {NgbModal, NgbModalRef, NgbActiveModal, ModalDismissReasons} from '@ng-bootstrap/ng-bootstrap';

import 'tinymce';

declare var tinymce: any;

@Component({
  selector: 'tiny-mce-popup',
  templateUrl: './tiny-mce-popup.html',
  styleUrls: ['./tiny-mce-popup.scss']
})
export class TinyMCEComponentPopup implements OnDestroy, AfterViewInit, OnChanges {
  @Input() title: string;
  @Output() confirm = new EventEmitter<any>();

  @Input() content: any;
  @Input() height: number;

  constructor(private host: ElementRef, public activeModal: NgbActiveModal) {}

  save(): any {
    this.removeTinymce();
    this.activeModal.close({act: 'save', data: this.content});
  }
  dismiss(): any {
    this.removeTinymce();
    this.activeModal.dismiss({act: 'cancel'});
  }

  ngOnChanges() {
    let editor = tinymce.get("mceEditor");
    if (editor) {editor.setContent(this.content);}
  }

  ngAfterViewInit() {

  }
  ngOnDestroy() {
    this.removeTinymce();
  }

  onEditorKeyup(event: any) {
    this.content = event;
  }

  removeTinymce() {
    if (tinymce.get("mceEditor")) {
      tinymce.get("mceEditor").remove();
    }
  }

}
