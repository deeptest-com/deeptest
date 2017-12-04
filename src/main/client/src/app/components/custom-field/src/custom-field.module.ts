import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';

import { CustomFieldComponent } from './custom-field.component';
import { CustomFieldService } from './custom-field.service';

@NgModule({
  imports: [CommonModule, FormsModule, ReactiveFormsModule],
  declarations: [CustomFieldComponent],
  exports: [CustomFieldComponent],
  providers: [CustomFieldService]
})
export class CustomFieldModule {

}
