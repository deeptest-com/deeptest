import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';

import { TinyMCEModule } from '../../../components/tiny-mce';

import { CustomFieldComponent } from './custom-field.component';
import { CustomFieldService } from './custom-field.service';

@NgModule({
  imports: [CommonModule, FormsModule, ReactiveFormsModule, TinyMCEModule],
  declarations: [CustomFieldComponent],
  exports: [CustomFieldComponent],
  providers: [CustomFieldService]
})
export class CustomFieldModule {

}
