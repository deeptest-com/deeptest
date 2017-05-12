import { Input, Output, EventEmitter, Component, ViewEncapsulation, OnInit, AfterViewInit } from '@angular/core';

import {NgbModal, ModalDismissReasons} from '@ng-bootstrap/ng-bootstrap';

import {GlobalState} from '../../../global.state';

@Component({
  selector: 'pop-dialog',
  styleUrls: ['./styles.scss'],
  templateUrl: './pop-dialog.html'
})
export class PopDialogComponent implements OnInit, AfterViewInit{

  @Input() title: string;
  @Output() confirm = new EventEmitter<any>();
  closeResult: string;

  constructor(private _state:GlobalState, private modalService: NgbModal) {

  }

  public ngOnInit(): void {
  }

  ngAfterViewInit() {

  }

  public showModal(content:any, cls: string): void {
    this.modalService.open(content, { windowClass: cls }).result.then((result) => {
      this.closeResult = `Closed with: ${result}`;
    }, (reason) => {
      this.closeResult = `Dismissed ${this.getDismissReason(reason)}`;
    });
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
