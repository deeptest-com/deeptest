import { NgModule, ModuleWithProviders } from "@angular/core";

import { EnvironmentConfigComponent } from './src/environment-config.component';
import { EnvironmentConfigService } from './src/environment-config.service';

export * from './src/environment-config.component';
export * from './src/environment-config.service';

@NgModule({
    declarations: [EnvironmentConfigComponent],
    exports: [EnvironmentConfigComponent],
    providers: [EnvironmentConfigService]
})
export class EnvironmentConfigModule {
    static forRoot(): ModuleWithProviders {
        return {
            ngModule: EnvironmentConfigModule,
            providers: [EnvironmentConfigService]
        };
    }
}
