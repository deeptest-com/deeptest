import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';

import { InputTextComponent } from './input-text/input-text.component';
import { InputTextService } from './input-text/input-text.service';

@NgModule({
  imports: [CommonModule, FormsModule, ReactiveFormsModule],
  declarations: [InputTextComponent],
  exports: [InputTextComponent],
  providers: [InputTextService]
})
export class CustomFieldModule {
}
