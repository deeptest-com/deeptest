import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule }  from '@angular/common';
import { FormsModule } from "@angular/forms";

import { CommentsService } from './src/comments.service';
import { CommentsComponent } from './src/comments.component';

export * from './src/comments.component';
export * from './src/comments.service';

@NgModule({
    declarations: [CommentsComponent],
    exports: [CommentsComponent],
    providers: [CommentsService],
    imports: [CommonModule, FormsModule]
})
export class CommentsModule {
    static forRoot(): ModuleWithProviders {
        return {
            ngModule: CommentsModule,
            providers: [CommentsService]
        };
    }
}
