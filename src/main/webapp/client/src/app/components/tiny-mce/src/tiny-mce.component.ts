import { Component, OnDestroy, AfterViewInit, OnChanges, Input, Output, EventEmitter, ElementRef } from '@angular/core';

import 'tinymce';
import 'tinymce/themes/modern';

import 'tinymce/plugins/table';
import 'tinymce/plugins/link';

declare var tinymce: any;

@Component({
  selector: 'ngx-tiny-mce',
  template: '<textarea [(ngModel)]="content" id="mceEditor"></textarea>',
})
export class TinyMCEComponent implements OnDestroy, AfterViewInit, OnChanges {
  @Output() editorKeyup = new EventEmitter<any>();

  @Input() content: any;
  @Input() modelId: any;
  editor: any;

  constructor(private host: ElementRef) { }

  ngAfterViewInit() {
    console.log('=ngAfterViewInit=', this.content);

    tinymce.init({
      document_base_url: '/assets/vendor/tinymce',
      selector: 'textarea#mceEditor',
      plugins: ['link', 'table'],
      skin_url: 'skins/lightgray',
      language : "zh_CN",
      language_url : "assets/vendor/tinymce/langs/zh_CN.js",
      setup: editor => {
        editor.on('keyup', () => {
          this.editor = editor;
          this.editorKeyup.emit(editor.getContent());
        });
      },
      height: '320',
    });
  }

  ngOnChanges() {
    let editor = tinymce.get("mceEditor");
    if (editor) {editor.setContent(this.content);}
  }

  ngOnDestroy() {
    tinymce.remove(this.editor);
  }

}
