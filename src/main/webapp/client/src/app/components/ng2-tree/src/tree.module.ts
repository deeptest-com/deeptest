import { NgModule } from '@angular/core';
import { TreeComponent } from './tree.component';
import { TreeInternalComponent } from './tree-internal.component';
import { TreeToolbarComponent } from './tree-toolbar.component';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';
import { NodeDraggableDirective } from './draggable/node-draggable.directive';
import { NodeDraggableService } from './draggable/node-draggable.service';
import { NodeEditableDirective } from './editable/node-editable.directive';
import { NodeMenuComponent } from './menu/node-menu.component';
import { NodeMenuService } from './menu/node-menu.service';
import { TreeService } from './tree.service';

@NgModule({
  imports: [CommonModule, FormsModule, ReactiveFormsModule],
  declarations: [NodeDraggableDirective, TreeComponent, NodeEditableDirective, NodeMenuComponent,
    TreeInternalComponent, TreeToolbarComponent],
  exports: [TreeComponent],
  providers: [NodeDraggableService, NodeMenuService, TreeService]
})
export class TreeModule {
}

