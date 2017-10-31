import { Component, OnDestroy, AfterViewInit, Output, EventEmitter, ElementRef } from '@angular/core';

import 'tinymce';
import 'tinymce/themes/modern';

import 'tinymce/plugins/table';
import 'tinymce/plugins/link';

declare var tinymce: any;

@Component({
  selector: 'ngx-tiny-mce',
  template: '',
})
export class TinyMCEComponent implements OnDestroy, AfterViewInit {

  @Output() editorKeyup = new EventEmitter<any>();

  editor: any;

  constructor(private host: ElementRef) { }

  ngAfterViewInit() {
    tinymce.init({
      target: this.host.nativeElement,
      plugins: ['link', 'table'],
      skin_url: 'assets/vendor/tinymce/skins/lightgray',
      language : "zh_CN",
      language_url : "assets/vendor/tinymce/langs/zh_CN.js",
      setup: editor => {
        this.editor = editor;
        editor.on('keyup', () => {
          this.editorKeyup.emit(editor.getContent());
        });
      },
      height: '320',
    });
  }

  ngOnDestroy() {
    tinymce.remove(this.editor);
  }
}
