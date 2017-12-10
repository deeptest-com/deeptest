import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule }  from '@angular/common';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';

import { SearchSelectModule } from '../search-select';

import { UserService } from '../../service/user';

import { RunEditService } from './src/run-edit.service';
import { RunEditComponent } from './src/run-edit.component';

export * from './src/run-edit.component';
export * from './src/run-edit.service';

@NgModule({
    declarations: [RunEditComponent],
    exports: [RunEditComponent],
    providers: [RunEditService],
    imports: [CommonModule, FormsModule, ReactiveFormsModule, SearchSelectModule]
})
export class RunEditModule {
    static forRoot(): ModuleWithProviders {
        return {
            ngModule: RunEditModule,
            providers: [UserService, RunEditService]
        };
    }
}
