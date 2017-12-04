import { Input, Output, EventEmitter, Component, ViewChild, OnInit, AfterViewInit, Injector } from '@angular/core';

import {NgbModal, NgbModalRef, ModalDismissReasons} from '@ng-bootstrap/ng-bootstrap';

import {GlobalState} from '../../../global.state';

@Component({
  selector: 'pop-dialog',
  styleUrls: ['./styles.scss'],
  templateUrl: './pop-dialog.html'
})
export class PopDialogComponent implements OnInit, AfterViewInit{

  @Input() title: string;
  @Output() confirm = new EventEmitter<any>();
  @ViewChild('content') content;
  modalRef: NgbModalRef;
  closeResult: string;

  constructor(private _state:GlobalState, private modalService: NgbModal, private injector: Injector) {

  }

  public ngOnInit(): void {
  }

  ngAfterViewInit() {

  }

  public showModal(cls?: string): void {
    let clsMap = cls? { windowClass: cls }: {};
    this.modalRef = this.modalService.open(this.content, clsMap);
  }

  public closeModal(): void {
    this.modalRef.close();
  }

  public onConfirm($event):void {
    this.confirm.emit($event);
  }

  private getDismissReason(reason: any): string {
    if (reason === ModalDismissReasons.ESC) {
      return 'by pressing ESC';
    } else if (reason === ModalDismissReasons.BACKDROP_CLICK) {
      return 'by clicking on a backdrop';
    } else {
      return  `with: ${reason}`;
    }
  }

  onModalShow():void {
    // init jquery components if needed
  }

}
