import {Component, Input, ChangeDetectionStrategy } from '@angular/core';

import { Cell } from '../../../lib/data-set/cell';

@Component({
  selector: 'table-cell-view-mode',
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `
    <div>
        <div>{{ cell.getValue() }}</div>
    </div>
    `,
})
export class ViewCellComponent {

  @Input() cell: Cell;
}
