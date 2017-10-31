import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule }  from '@angular/common';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';

import { ZtreeModule } from '../ztree';
import { SelectModule } from '../ng2-select';

import { SuiteService } from '../../service/suite';
import { CaseService } from '../../service/case';

import { CaseSelectionService } from './src/case-selection.service';
import { CaseSelectionComponent } from './src/case-selection.component';

export * from './src/case-selection.component';
export * from './src/case-selection.service';

@NgModule({
    declarations: [CaseSelectionComponent],
    exports: [CaseSelectionComponent],
    providers: [CaseSelectionService],
    imports: [CommonModule, FormsModule, ReactiveFormsModule, ZtreeModule, SelectModule]
})
export class CaseSelectionModule {
    static forRoot(): ModuleWithProviders {
        return {
            ngModule: CaseSelectionModule,
            providers: [SuiteService, CaseService, CaseSelectionService]
        };
    }
}
