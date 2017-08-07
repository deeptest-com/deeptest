import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule }  from '@angular/common';

import { NgbTabsetModule } from '@ng-bootstrap/ng-bootstrap';

import { AngularEchartsModule } from 'ngx-echarts';

import { ExecutionReportComponent } from './src/execution-report.component';
import { ExecutionResultComponent } from './src/execution-result/execution-result.component';
import { ExecutionActivityComponent } from './src/execution-activity/execution-activity.component';
import { ExecutionProgressComponent } from './src/execution-progress/execution-progress.component';

export * from './src/execution-report.component';
export * from './src/execution-result/execution-result.component';

@NgModule({
    declarations: [ExecutionReportComponent, ExecutionResultComponent, ExecutionActivityComponent, ExecutionProgressComponent],
    exports: [ExecutionReportComponent, ExecutionResultComponent, ExecutionActivityComponent, ExecutionProgressComponent],
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
