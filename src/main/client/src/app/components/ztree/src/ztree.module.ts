import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';

import { ZtreeComponent } from './ztree.component';
import { ZtreeService } from './ztree.service';

@NgModule({
  imports: [CommonModule, FormsModule, ReactiveFormsModule],
  declarations: [ZtreeComponent],
  exports: [ZtreeComponent],
  providers: [ZtreeService]
})
export class ZtreeModule {

}
