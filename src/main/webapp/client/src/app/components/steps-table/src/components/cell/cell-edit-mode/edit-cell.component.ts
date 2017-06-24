import { Component, Input, Output, EventEmitter } from '@angular/core';

import { Cell } from '../../../lib/data-set/cell';

@Component({
  selector: 'table-cell-edit-mode',
  template: `
      <div>
        <table-cell-default-editor [cell]="cell">
        </table-cell-default-editor>
      </div>
    `,
})
export class EditCellComponent {

  @Input() cell: Cell;

  getEditorType(): string {
    return this.cell.getColumn().editor && this.cell.getColumn().editor.type;
  }
}
