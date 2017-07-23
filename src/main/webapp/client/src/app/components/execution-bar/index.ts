import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule }  from '@angular/common';

import { NgbTooltipModule } from '@ng-bootstrap/ng-bootstrap';

import { ExecutionBarComponent } from './src/execution-bar.component';

export * from './src/execution-bar.component';

@NgModule({
    declarations: [ExecutionBarComponent],
    exports: [ExecutionBarComponent],
    providers: [],
    imports: [CommonModule, NgbTooltipModule]
})
export class ExecutionBarModule {
    static forRoot(): ModuleWithProviders {
        return {
            ngModule: ExecutionBarModule,
            providers: []
        };
    }
}
