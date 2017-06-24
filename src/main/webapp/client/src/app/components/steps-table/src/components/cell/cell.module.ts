import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Ng2CompleterModule } from 'ng2-completer';

import { CellComponent } from './cell.component';
import { DefaultEditComponent } from './cell-edit-mode/default-edit.component';
import { EditCellComponent } from './cell-edit-mode/edit-cell.component';
import { InputEditorComponent } from './cell-editors/input-editor.component';
import { TextareaEditorComponent } from './cell-editors/textarea-editor.component';
import { ViewCellComponent } from './cell-view-mode/view-cell.component';

const CELL_COMPONENTS = [
  CellComponent,
  DefaultEditComponent,
  EditCellComponent,
  InputEditorComponent,
  TextareaEditorComponent,
  ViewCellComponent,
];

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    Ng2CompleterModule,
  ],
  declarations: [
    ...CELL_COMPONENTS,
  ],
  exports: [
    ...CELL_COMPONENTS,
  ],
})
export class CellModule { }
