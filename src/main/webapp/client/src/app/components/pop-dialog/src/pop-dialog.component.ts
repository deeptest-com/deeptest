import { Input, Output, EventEmitter, Component, OnInit, AfterViewInit, ViewChild, ElementRef } from '@angular/core';

import { ModalDirective } from 'ng2-bootstrap';

import {GlobalState} from '../../../global.state';

@Component({
  selector: 'pop-dialog',
  styles: [require('./styles.scss')],
  template: require('./pop-dialog.html')
})
export class PopDialogComponent implements OnInit, AfterViewInit{

  @Input() title: string;
  @Output() confirm = new EventEmitter<any>();

  @ViewChild('modal') modal:ModalDirective;

  constructor(private _state:GlobalState, private el: ElementRef) {

  }

  public ngOnInit(): void {
  }

  ngAfterViewInit() {

  }

  public showModal(): void {
    this.modal.show();
  }

  public hideModal(): void {
    this.modal.hide();
  }

  onModalShow():void {
    // init jquery components if needed
  }

  public onConfirm($event):void {
    this.confirm.emit($event);
  }
}
