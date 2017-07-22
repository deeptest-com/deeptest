import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule }  from '@angular/common';

import { ExecutionResultComponent } from './src/execution-result/execution-result.component';

export * from './src/execution-result/execution-result.component';

@NgModule({
    declarations: [ExecutionResultComponent],
    exports: [ExecutionResultComponent],
    providers: [],
    imports: [CommonModule]
})
export class ExecutionReportModule {
    static forRoot(): ModuleWithProviders {
        return {
            ngModule: ExecutionReportModule,
            providers: []
        };
    }
}
