import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { FormsModule } from "@angular/forms";

import { CommentEditComponent } from './src/comment-edit.component';

export * from './src/comment-edit.component';

@NgModule({
  imports: [CommonModule, RouterModule, FormsModule],
  declarations: [CommentEditComponent],
  exports: [CommentEditComponent],
  providers: []
})
export class CommentEditModule {
  static forRoot(): ModuleWithProviders {
    return {
      ngModule: CommentEditModule,
      providers: []
    };
  }
}
