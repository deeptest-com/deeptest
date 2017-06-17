import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';

import { CustomFieldComponent } from './custom-field.component';
import { CustomFieldService } from './custom-field.service';

import { InputTextComponent } from './input-text/input-text.component';
import { InputTextService } from './input-text/input-text.service';

@NgModule({
  imports: [CommonModule, FormsModule, ReactiveFormsModule],
  declarations: [CustomFieldComponent, InputTextComponent],
  exports: [CustomFieldComponent, InputTextComponent],
  providers: [CustomFieldService, InputTextService]
})
export class CustomFieldModule {

}
