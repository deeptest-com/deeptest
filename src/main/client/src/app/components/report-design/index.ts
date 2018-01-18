import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule }  from '@angular/common';

import { NgbTabsetModule } from '@ng-bootstrap/ng-bootstrap';

import { AngularEchartsModule } from 'ngx-echarts';

import { ReportDesignComponent } from './src/report-design.component';
import { DesignActivityComponent } from './src/design-activity/design-activity.component';
import { DesignProgressComponent } from './src/design-progress/design-progress.component';

export * from './src/report-design.component';

@NgModule({
    declarations: [ReportDesignComponent, DesignActivityComponent, DesignProgressComponent],
    exports: [ReportDesignComponent, DesignActivityComponent, DesignProgressComponent],
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
