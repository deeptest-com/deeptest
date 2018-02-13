import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../../../theme/nga.module';

import { routing }       from './custom-field.routing';

import { NgbModalModule, NgbPaginationModule, NgbDropdownModule,
  NgbTabsetModule, NgbButtonsModule, NgbCollapseModule } from '@ng-bootstrap/ng-bootstrap';

import { PipeModule } from '../../../../pipe/pipe.module';
import { DirectiveModule } from '../../../../directive/directive.module';
import { PopDialogModule } from '../../../../components/pop-dialog';
import { DropdownOptionsModule } from '../../../../components/dropdown-options';
import { DropdownOptionsComponent } from '../../../../components/dropdown-options';

import { RouteService } from '../../../../service/route';
import { RequestService } from '../../../../service/request';
import { DatetimePickerService } from '../../../../service/datetime-picker';
import { CustomFieldService } from '../../../../service/custom-field';

import { CustomField } from './custom-field.component';
import { CustomFieldList } from './list';
import { CustomFieldEdit } from './edit';

@NgModule({
  imports: [
    CommonModule,
    FormsModule, ReactiveFormsModule,
    NgaModule,
    routing,

    NgbModalModule, NgbPaginationModule, NgbDropdownModule,
    NgbTabsetModule, NgbButtonsModule, NgbCollapseModule,

    PipeModule, DirectiveModule, PopDialogModule, DropdownOptionsModule
  ],
  declarations: [
    CustomField,
    CustomFieldList,
    CustomFieldEdit
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    CustomFieldService
  ],
  entryComponents: [DropdownOptionsComponent]
})
export class CustomFieldModule {}

