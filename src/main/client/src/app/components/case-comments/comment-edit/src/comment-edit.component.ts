import { Input, Output, EventEmitter, Component, ViewChild, OnInit, AfterViewInit, Injector, ElementRef } from '@angular/core';

import {NgbModal, NgbModalRef, ModalDismissReasons} from '@ng-bootstrap/ng-bootstrap';

import {GlobalState} from '../../../../global.state';

@Component({
  selector: 'comment-edit',
  styleUrls: ['./styles.scss'],
  templateUrl: './comment-edit.html'
})
export class CommentEditComponent implements OnInit, AfterViewInit{
  @Input() @Output() model: any = {};
  @ViewChild('content') content: ElementRef;

  @Output() confirm = new EventEmitter<any>();
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

  public onConfirm():void {
    this.confirm.emit();
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
