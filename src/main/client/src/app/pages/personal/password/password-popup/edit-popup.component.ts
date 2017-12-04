import { Component, OnDestroy, AfterViewInit, OnChanges,ViewChild, Input, Output, EventEmitter, Injector, ElementRef } from '@angular/core';
import {NgbModal, NgbModalRef, NgbActiveModal, ModalDismissReasons} from '@ng-bootstrap/ng-bootstrap';

@Component({
  selector: 'password-edit-popup',
  templateUrl: './edit-popup.html',
  styleUrls: ['./edit-popup.scss']
})
export class PasswordEditPopupComponent implements OnDestroy, AfterViewInit, OnChanges {
  @Input() title: string;
  @Output() confirm = new EventEmitter<any>();

  constructor(private host: ElementRef, public activeModal: NgbActiveModal) {}

  save(): any {
    this.activeModal.close({act: 'save'});
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

}
