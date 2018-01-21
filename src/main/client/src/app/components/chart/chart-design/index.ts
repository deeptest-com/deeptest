import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule }  from '@angular/common';

import { NgbTabsetModule } from '@ng-bootstrap/ng-bootstrap';

import { AngularEchartsModule } from 'ngx-echarts';

import { DesignProcessComponent } from './src/design-process/design-process.component';
import { DesignProgressComponent } from './src/design-progress/design-progress.component';

@NgModule({
    declarations: [DesignProcessComponent, DesignProgressComponent],
    exports: [DesignProcessComponent, DesignProgressComponent],
    providers: [],
    imports: [CommonModule, NgbTabsetModule, AngularEchartsModule]
})
export class ChartDesignModule {
    static forRoot(): ModuleWithProviders {
        return {
            ngModule: ChartDesignModule,
            providers: []
        };
    }
}
