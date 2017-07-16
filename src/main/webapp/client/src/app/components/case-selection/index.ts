import { NgModule, ModuleWithProviders } from "@angular/core";

import { CaseSelectionComponent } from './src/case-selection.component';
import { CaseSelectionService } from './src/case-selection.service';

export * from './src/case-selection.component';
export * from './src/case-selection.service';

@NgModule({
    declarations: [CaseSelectionComponent],
    exports: [CaseSelectionComponent],
    providers: [CaseSelectionService]
})
export class CaseSelectionModule {
    static forRoot(): ModuleWithProviders {
        return {
            ngModule: CaseSelectionModule,
            providers: [CaseSelectionService]
        };
    }
}
