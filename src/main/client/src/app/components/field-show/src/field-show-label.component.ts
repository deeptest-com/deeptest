import { Component, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'field-show-label',
  templateUrl: './field-show-label.html',
  styleUrls: ['./label.scss']
})
export class FieldShowLabelComponent {

  @Input() label: string;

}
