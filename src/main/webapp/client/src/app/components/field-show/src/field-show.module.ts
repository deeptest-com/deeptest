import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';

import { FieldShowComponent } from './field-show.component';
import { FieldShowLabelComponent } from './field-show-label.component';
import { FieldShowService } from './field-show.service';

@NgModule({
  imports: [CommonModule, FormsModule, ReactiveFormsModule],
  declarations: [FieldShowComponent, FieldShowLabelComponent],
  exports: [FieldShowComponent, FieldShowLabelComponent],
  providers: [FieldShowService]
})
export class FieldShowModule {

}
