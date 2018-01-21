import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule }  from '@angular/common';

import { NgbTabsetModule } from '@ng-bootstrap/ng-bootstrap';

import { AngularEchartsModule } from 'ngx-echarts';

import { ReportDesignComponent } from './src/report-design.component';
import { DesignProcessComponent } from './src/design-process/design-process.component';
import { DesignProgressComponent } from './src/design-progress/design-progress.component';

export * from './src/report-design.component';

@NgModule({
    declarations: [ReportDesignComponent, DesignProcessComponent, DesignProgressComponent],
    exports: [ReportDesignComponent, DesignProcessComponent, DesignProgressComponent],
    providers: [],
    imports: [CommonModule, NgbTabsetModule, AngularEchartsModule]
})
export class ReportDesignModule {
    static forRoot(): ModuleWithProviders {
        return {
            ngModule: ReportDesignModule,
            providers: []
        };
    }
}
