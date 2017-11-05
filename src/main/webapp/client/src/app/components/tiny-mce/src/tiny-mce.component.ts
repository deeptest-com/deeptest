import { Component, OnDestroy, AfterViewInit, OnChanges, Input, Output, EventEmitter, ElementRef } from '@angular/core';

import 'tinymce';
import 'tinymce/themes/modern';

import 'tinymce/plugins/table';
import 'tinymce/plugins/link';

declare var tinymce: any;

@Component({
  selector: 'tiny-mce',
  templateUrl: './tiny-mce.html',
  styleUrls: ['./tiny-mce.scss']
})
export class TinyMCEComponent implements OnDestroy, AfterViewInit, OnChanges {
  @Output() editorKeyup = new EventEmitter<any>();

  @Input() content: any;
  @Input() modelId: any;
  @Input() height: string;

  constructor(private host: ElementRef) { }

  ngAfterViewInit() {
    let that = this;
    console.log("tinymce ngAfterViewInit");
    if (!$('textarea#mceEditor')) {
      return;
    }

    tinymce.init({
      document_base_url: '/assets/vendor/tinymce',
      selector: 'textarea#mceEditor',
      plugins: ['link', 'table'],
      skin_url: 'skins/lightgray',
      language : "zh_CN",
      language_url : "assets/vendor/tinymce/langs/zh_CN.js",
      setup: editor => {
        editor.on('keyup', () => {
          this.editorKeyup.emit(editor.getContent());
        });
      },
      height: this.height
    }).then(function(editors) {
      that.updateContent();
    });
  }

  ngOnChanges() {
    this.updateContent();
  }

  ngOnDestroy() {
    console.log("tinymce ngOnDestroy");
    if (tinymce.get("mceEditor")) {
      tinymce.get("mceEditor").remove();
    }
  }

  updateContent() {
    let editor = tinymce.get("mceEditor");
    console.log("tinymce ngOnChanges", editor, this.content);
    if (editor) {editor.setContent(this.content);}
  }

}
