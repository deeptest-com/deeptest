import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule }  from '@angular/common';

import { NgbTabsetModule } from '@ng-bootstrap/ng-bootstrap';

import { AngularEchartsModule } from 'ngx-echarts';

import { ExecutionReportComponent } from './src/execution-report.component';
import { ExecutionResultComponent } from './src/execution-result/execution-result.component';

export * from './src/execution-report.component';
export * from './src/execution-result/execution-result.component';

@NgModule({
    declarations: [ExecutionReportComponent, ExecutionResultComponent],
    exports: [ExecutionReportComponent, ExecutionResultComponent],
    providers: [],
    imports: [CommonModule, NgbTabsetModule, AngularEchartsModule]
})
export class ExecutionReportModule {
    static forRoot(): ModuleWithProviders {
        return {
            ngModule: ExecutionReportModule,
            providers: []
        };
    }
}
