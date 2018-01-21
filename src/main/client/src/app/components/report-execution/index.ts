import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule }  from '@angular/common';

import { NgbTabsetModule } from '@ng-bootstrap/ng-bootstrap';

import { AngularEchartsModule } from 'ngx-echarts';

import { ReportExecutionComponent } from './src/report-execution.component';
import { ExecutionResultComponent } from './src/execution-result/execution-result.component';
import { ExecutionProcessComponent } from './src/execution-process/execution-process.component';
import { ExecutionProgressComponent } from './src/execution-progress/execution-progress.component';

export * from './src/report-execution.component';
export * from './src/execution-result/execution-result.component';

@NgModule({
    declarations: [ReportExecutionComponent, ExecutionResultComponent, ExecutionProcessComponent, ExecutionProgressComponent],
    exports: [ReportExecutionComponent, ExecutionResultComponent, ExecutionProcessComponent, ExecutionProgressComponent],
    providers: [],
    imports: [CommonModule, NgbTabsetModule, AngularEchartsModule]
})
export class ReportExecutionModule {
    static forRoot(): ModuleWithProviders {
        return {
            ngModule: ReportExecutionModule,
            providers: []
        };
    }
}
