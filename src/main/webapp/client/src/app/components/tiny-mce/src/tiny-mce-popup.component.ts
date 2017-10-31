import { Component, OnDestroy, AfterViewInit, OnChanges,ViewChild, Input, Output, EventEmitter, Injector, ElementRef } from '@angular/core';
import {NgbModal, NgbModalRef, NgbActiveModal, ModalDismissReasons} from '@ng-bootstrap/ng-bootstrap';

import 'tinymce';

declare var tinymce: any;

@Component({
  selector: 'ngx-tiny-mce-popup',
  templateUrl: './tiny-mce-popup.html',
  styleUrls: ['./tiny-mce-popup.scss']
})
export class TinyMCEComponentPopup {
  @Input() title: string;
  @Output() confirm = new EventEmitter<any>();

  @Input() content: any;
  @Input() modelId: any;
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

  onEditorKeyup(event: any) {
    this.content = event;
  }

  removeTinymce() {
    if (tinymce.get("mceEditor")) {
      tinymce.get("mceEditor").remove();
    }
  }

}
