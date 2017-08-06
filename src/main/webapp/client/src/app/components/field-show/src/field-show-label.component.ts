import { Component, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'field-show-label',
  styleUrls: ['./label.scss'],
  template: `
    <span *ngIf="label" class="field-show-label">{{label}}：</span>
  `,
})
export class FieldShowLabelComponent {

  @Input() label: string;

}
