import { Component, Output, EventEmitter, Input } from '@angular/core';

import { Cell } from '../../../lib/data-set/cell';

export class EditCellDefault {

  @Input() cell: Cell;
}
