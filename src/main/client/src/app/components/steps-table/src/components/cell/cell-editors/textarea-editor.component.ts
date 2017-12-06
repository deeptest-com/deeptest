import { Component } from '@angular/core';

import { DefaultEditor } from './default-editor';

@Component({
  selector: 'textarea-editor',
  styleUrls: ['./editor.component.scss'],
  template: `
    <textarea class="form-control"
              [(ngModel)]="cell.newValue"
              [name]="cell.getId()"
              [disabled]="!cell.isEditable()"
              [placeholder]="cell.getTitle()">
    </textarea>
    `,
})
export class TextareaEditorComponent extends DefaultEditor {

  constructor() {
    super();
  }
}
